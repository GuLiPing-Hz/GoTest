package net

import (
	"bytes"
	"io"
	"net"
	"runtime"
	"runtime/debug"
	"strings"
	"sync"
	"sync/atomic"
	"time"
	"util"
)

type Conn interface {
	//主动关闭连接
	Close() error
	// 封装发送buffer
	Send(buf [] byte)

	//是否是被对方关闭了连接
	IsClosedByPeer() bool

	//获取最近的错误
	LastErr() error

	//获取最近的调用堆栈，如果有的话
	LastStack() []byte

	// LocalAddr returns the local network address.
	LocalAddr() net.Addr

	// RemoteAddr returns the remote network address.
	RemoteAddr() net.Addr
}

type ConnInner interface {
	SendEx([]byte)
	RecvEx() ([]byte, error)
}

type OnSocket interface {
	/**
	连接上服务器回调
	*/
	OnConnect(Conn)
	/**
	只要我们曾经连接上服务器过，OnClose必定会回调。代表一个当前的socket已经关闭
	*/
	OnClose(Conn)
	/**
	连接超时,写入超时,读取超时回调
	*/
	OnTimeout(Conn)
	/**
	网络错误回调，之后直接close
	*/
	OnNetErr(Conn)
	/**
	接受到信息
	*/
	OnRecvMsg(Conn, []byte) bool
}

type DataDecodeBase interface {
	GetPackageHeadLen() int
	GetPackageLen([]byte) int
}

type ClientBase struct {
	ConnBase  Conn
	ConnInner ConnInner
	StatusErr
	ReadDB      *bytes.Buffer
	WriteDB     *bytes.Buffer
	WMutex      sync.Mutex
	DataDecoder DataDecodeBase
	Ttl         time.Duration
	Rttl        time.Duration
	OnSocket    OnSocket
}

func (this *ClientBase) send() {
	this.WMutex.Lock()
	buffer := make([]byte, this.WriteDB.Len())
	_, _ = this.WriteDB.Read(buffer)
	this.WriteDB.Reset()
	this.WMutex.Unlock()
	if len(buffer) == 0 {
		return
	}

	this.ConnInner.SendEx(buffer)
}

func (this *ClientBase) Send(buf [] byte) {
	if this.GetStatus() != StatusNormal {
		return
	}

	this.WMutex.Lock()
	defer this.WMutex.Unlock()

	n, err := this.WriteDB.Write(buf)
	if err != nil {
		util.E("Send err=%s", err.Error())
		return
	}
	if n == 0 {
		return
	}

	go this.send()
}

func (this *ClientBase) SafeClose(needCB bool) {
	if atomic.LoadInt32(&this.IsConnected) == 1 {
		atomic.StoreInt32(&this.IsConnected, 0)
		this.ChangeStatus(StatusShutdown, nil)
		//util.D("Close socket")
		err := this.ConnBase.Close()
		if err != nil {
			util.E("Close error=%v", err.Error())
		}

		if needCB {
			this.OnSocket.OnClose(this.ConnBase) //如果需要回调，我们就回调一下。
		}
	}
}

func (this *ClientBase) CloseWithErr(err error) {
	this.ChangeStatus(StatusError, err)
	this.OnSocket.OnNetErr(this.ConnBase)
	//把发生错误的socket及时关闭
	this.SafeClose(true)
}

func (this *ClientBase) CloseTimeout() {
	this.ChangeStatus(StatusTimeout, nil)
	this.OnSocket.OnTimeout(this.ConnBase)
	//把发生错误的socket及时关闭
	this.SafeClose(true)
}

func (this *ClientBase) IsClosedByPeer() bool {
	return this.Err == ErrClose
}

//获取最近的错误
func (this *ClientBase) LastErr() error {
	return this.Err
}

//获取最近的调用堆栈，如果有的话
func (this *ClientBase) LastStack() []byte {
	return this.Stack
}

func (this *ClientBase) process() error {
	defer func() {
		this.ReadDB = bytes.NewBuffer(this.ReadDB.Bytes())
	}()

	lenHead := int(this.DataDecoder.GetPackageHeadLen())

	for {
		if this.ReadDB.Len() <= 0 || this.ReadDB.Len() < lenHead { //不足包长
			return nil
		}

		lenPackage := this.DataDecoder.GetPackageLen(this.ReadDB.Bytes())
		if lenPackage == 0 { //异常包
			this.ReadDB.Reset()
			return ErrBuffer
		}

		lenFull := lenHead + lenPackage
		if lenFull > this.ReadDB.Len() { //不足一个包
			return nil
		}

		packageBuf := make([]byte, lenFull)
		_, _ = this.ReadDB.Read(packageBuf)
		//util.D("read copy buf len=%d", lenFull)
		ok := this.OnSocket.OnRecvMsg(this.ConnBase, packageBuf)
		if ! ok {
			this.ReadDB.Reset()
			return io.EOF
		}
	}
}

func (this *ClientBase) Rector() {
	go rector(this)
}

func rector(client *ClientBase) {
	defer func() {
		p := recover()
		if err, ok := p.(error); ok {
			client.ChangeStatusAll(StatusError, err, debug.Stack())
			client.OnSocket.OnNetErr(client.ConnBase)
		}

		client.SafeClose(true)
	}()

	atomic.StoreInt32(&client.IsConnected, 1)
	client.ChangeStatus(StatusNormal, nil)
	client.OnSocket.OnConnect(client.ConnBase)
	for {
		buffer, err := client.ConnInner.RecvEx()
		if err != nil {
			err1, ok := err.(*net.OpError)
			errStr := err.Error()
			if err == io.EOF || (runtime.GOOS == "windows" &&
				strings.Contains(errStr, "An existing connection was forcibly closed by the remote host") ||
				strings.Contains(errStr, "An established connection was aborted by the software in your host machine")) ||
				strings.Contains(errStr, "connection reset by peer") {
				/*
					1.io.EOF
						正常关闭.指客户端读完服务器发送的数据然后close

					2.
					connection reset by peer(linux)
					An existing connection was forcibly closed by the remote host(windows)
						表示客户端 【没有读取/读取部分】就close

					3.An established connection was aborted by the software in your host machine(windows)
						表示服务器发送数据，客户端已经close,这个经过测试只有在windows上才会出现。linux试了很多遍都是返回io.EOF错误
						解决办法就是客户端发送数据的时候需要wait一下，然后再close，这样close的结果就是2了
				*/
				client.CloseWithErr(ErrClose)
			} else if ok && err1 != nil && err1.Timeout() {
				client.CloseTimeout()
			} else {
				if client.ChangeStatus(StatusError, err) { //检查是否已经更改了状态，如果已经更改表示是客户端主动close
					client.OnSocket.OnNetErr(client.ConnBase)
				}
			}
			return
		}

		_, err = client.ReadDB.Write(buffer)
		if err != nil {
			client.CloseWithErr(ErrOOM) //无法把buffer全部塞进去，多半是没有内存了。
			return
		}

		err = client.process()
		if err != nil {
			client.CloseWithErr(err)
			return
		}
	}
}

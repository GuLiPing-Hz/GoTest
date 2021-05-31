package net2

import (
	"bytes"
	"encoding/binary"
	"io"
	"mybase"
	"net"
	"runtime"
	"runtime/debug"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

/**
对外宣称对象
*/
type Conn interface {
	//获取最近的错误
	error

	//安全关闭连接
	SafeClose()

	// 封装发送buffer
	Send(buf []byte) bool

	//是否是被对方关闭了连接
	IsClosedByPeer() bool

	//获取最近的调用堆栈，如果有的话
	Stack() []byte

	LocalAddr() net.Addr

	RemoteAddr() net.Addr
}

type Socket interface {
	Conn

	Close() error

	SendEx(buffer []byte)

	RecvEx() ([]byte, error)
}

/**
Socket 事件分发
*/
type OnSocket interface {
	/**
	连接上服务器回调,或者服务器accept某个客户端连接
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
	@return 返回true表示可以继续热恋，false表示要分手了。
	*/
	OnRecvMsg(Conn, []byte) bool
}

/**
Socket 数据解析
*/
type DataDecodeBase interface {
	GetPackageHeadLen() int
	GetPackageLen([]byte) int
}

type StatusNO int32

const (
	StatusUnknown  StatusNO = iota //0 尚未初始化
	StatusNormal                   //1 正常
	StatusShutdown                 //2 自己关闭的
	StatusTimeout                  //3 超时连接断开
	StatusError                    //4 其他异常断开,服务器主动断开 err为ErrClose
)

type StatusErr struct {
	status      StatusNO //StatusNormal
	err         error    //当为其他异常时，这里会有赋值
	sta         []byte   //调用的错误堆栈信息
	mutex       sync.Mutex
	isConnected int32 //0未连接，1已连接
}

func (imp *StatusErr) Error() string {
	return imp.err.Error()
}

func (imp *StatusErr) Stack() []byte {
	return imp.sta
}

func (imp *StatusErr) GetStatus() StatusNO {
	imp.mutex.Lock()
	defer imp.mutex.Unlock()
	return imp.status
}

func (imp *StatusErr) ChangeStatus(status StatusNO, err error) bool {
	return imp.ChangeStatusAll(status, err, nil)
}

func (imp *StatusErr) ChangeStatusAll(status StatusNO, err error, stack []byte) bool {
	imp.mutex.Lock()
	defer imp.mutex.Unlock()
	//只记录正常状态或者赋值为初始状态。
	if imp.status == StatusNormal || imp.status == StatusUnknown || status == StatusUnknown {
		imp.status = status
		imp.err = err
		imp.sta = stack
		return true
	}
	return false
}

type ClientBase struct {
	StatusErr

	readDB     *bytes.Buffer
	chanSendDB chan []byte
	chanStop   chan bool

	dataDecoder DataDecodeBase

	Ttl  time.Duration //写超时
	RTtl time.Duration //读超时

	onSocket OnSocket
	socket   Socket
}

func (imp *ClientBase) Init(ddb DataDecodeBase, ttl, RTtl time.Duration, onSocket OnSocket, socket Socket) {
	imp.readDB = &bytes.Buffer{}
	imp.dataDecoder = ddb
	if imp.dataDecoder == nil {
		imp.dataDecoder = new(DataDecode)
	}
	imp.Ttl = ttl
	imp.RTtl = RTtl

	imp.onSocket = onSocket
	imp.socket = socket
}

func (imp *ClientBase) SafeClose() {
	if atomic.LoadInt32(&imp.isConnected) == 1 {
		atomic.StoreInt32(&imp.isConnected, 0)
		imp.ChangeStatus(StatusShutdown, nil)
		//mybase.D("Close socket")
		err := imp.socket.Close()
		if err != nil {
			mybase.E("Close error=%v", err.Error())
		}

		close(imp.chanStop)
		close(imp.chanSendDB)
		imp.onSocket.OnClose(imp.socket) //如果需要回调，我们就回调一下。
	}
}

func (imp *ClientBase) send() {
	for {
		select {
		case buf := <-imp.chanSendDB:
			imp.socket.SendEx(buf)
		case <-imp.chanStop:
			return
		}
	}
}

func (imp *ClientBase) Send(buf []byte) bool {
	if imp.GetStatus() != StatusNormal {
		return false
	}
	imp.chanSendDB <- buf
	return true
}

func (imp *ClientBase) IsClosedByPeer() bool {
	return imp.err == mybase.ErrClose
}

func (imp *ClientBase) CloseWithErr(err error) {
	imp.ChangeStatus(StatusError, err)
	imp.onSocket.OnNetErr(imp.socket)
	//把发生错误的socket及时关闭
	imp.SafeClose()
}

func (imp *ClientBase) CloseTimeout() {
	imp.ChangeStatus(StatusTimeout, nil)
	imp.onSocket.OnTimeout(imp.socket)
	//把发生错误的socket及时关闭
	imp.SafeClose()
}

func (imp *ClientBase) process() error {
	defer func() {
		imp.readDB = bytes.NewBuffer(imp.readDB.Bytes()) //舍去已经读取的buffer，保留尚未读取的buffer
	}()

	lenHead := imp.dataDecoder.GetPackageHeadLen()

	for {
		readLen := imp.readDB.Len()
		if readLen <= 0 || readLen < lenHead { //不足包长
			return nil
		}

		lenPackage := imp.dataDecoder.GetPackageLen(imp.readDB.Bytes())
		if lenPackage == 0 { //异常包
			imp.readDB.Reset()
			return mybase.ErrBuffer
		}

		lenFull := lenHead + lenPackage
		if lenFull > imp.readDB.Len() { //不足一个包
			return nil
		}

		packageBuf := make([]byte, lenFull)
		_, _ = imp.readDB.Read(packageBuf)
		//mybase.D("read copy buf len=%d", lenFull)
		ok := imp.onSocket.OnRecvMsg(imp.socket, packageBuf)
		if !ok {
			imp.readDB.Reset()
			return io.EOF
		}
	}
}

func (imp *ClientBase) Rector() {
	atomic.StoreInt32(&imp.isConnected, 1)
	imp.ChangeStatus(StatusNormal, nil)

	imp.chanSendDB = make(chan []byte)
	imp.chanStop = make(chan bool)

	go rector(imp)
}

func rector(client *ClientBase) {
	defer func() {
		p := recover()
		if err, ok := p.(error); ok {
			client.ChangeStatusAll(StatusError, err, debug.Stack())
			client.onSocket.OnNetErr(client.socket)
		}

		client.SafeClose()
	}()

	go client.send()
	client.onSocket.OnConnect(client.socket)
	for {
		buffer, err := client.socket.RecvEx()
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
				client.CloseWithErr(mybase.ErrClose)
			} else if ok && err1 != nil && err1.Timeout() {
				client.CloseTimeout()
			} else {
				if client.ChangeStatus(StatusError, err) { //检查是否已经更改了状态，如果已经更改表示是客户端主动close
					client.onSocket.OnNetErr(client.socket)
				}
			}
			return
		}

		_, err = client.readDB.Write(buffer)
		if err != nil {
			client.CloseWithErr(mybase.ErrOOM) //无法把buffer全部塞进去，多半是没有内存了。
			return
		}

		err = client.process()
		if err != nil {
			client.CloseWithErr(err)
			return
		}
	}
}

//***************** 默认的数据解析
type DataDecode struct {
}

func (*DataDecode) GetPackageHeadLen() int {
	//单一包的头部长度 默认为2字节
	return 2
}
func (*DataDecode) GetPackageLen(bytes []byte) int {
	//大端的形式获取包长
	return int(binary.BigEndian.Uint16(bytes))
}

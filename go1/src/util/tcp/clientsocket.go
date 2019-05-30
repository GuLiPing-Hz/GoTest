package tcp

import (
	"fmt"
	"io"
	"net"
	"runtime"
	"strings"
	"sync"
	"time"
	"util"
)

type OnSocket interface {
	Connect(*ClientSocketBase)
	Timeout(*ClientSocketBase)
	Close(*ClientSocketBase)
	ConnectErr(*ClientSocketBase)
	RecvErr(*ClientSocketBase)
	SendErr(*ClientSocketBase)
	NetErr(*ClientSocketBase)
}

type DataDecodeBase interface {
	Process(*ClientSocketBase)
}

type DataBlock struct {
	buffer []byte
	pos    int32
}

/**
添加到buffer后面
*/
func (this *DataBlock) Append(buf []byte) int32 {
	this.buffer = append(this.buffer, buf...)
	return int32(len(buf))
}

//@todo 这里需要测试一下
/**
copy数据到指定位置后面
*/
func (this *DataBlock) Copy(pos int32, buf []byte) int32 {
	//copy(this.buffer, buf) //Copy returns the number of elements copied, which will be the minimum of len(src) and len(dst)
	//由于内建的copy是可能存在没有拷贝完的效果，我们这里改写下
	this.buffer = this.buffer[0:pos] //这里只保留我们制定位置以前的数据。
	return this.Append(buf)
}

/**
完全覆盖原先的数据
*/
func (this *DataBlock) Move(buf []byte) int32 {
	this.buffer = buf
	return int32(len(this.buffer))
}

func (this *DataBlock) GetBuf() []byte {
	return this.buffer
}

func (this *DataBlock) GetPos() int32 {
	return int32(len(this.buffer))
}

func (this *DataBlock) InitPos() {
	this.buffer = make([]byte, 0)
}

type ClientSocketBase struct {
	StatusConn
	readDB      DataBlock
	rMutex      sync.Mutex
	writeDB     DataBlock
	wMutex      sync.Mutex
	dataDecoder DataDecodeBase
	ttl         int32
	onSocket    OnSocket
}

func (this *ClientSocketBase) GetReadDB() *DataBlock {
	return &this.readDB
}

func (this *ClientSocketBase) GetWriteDB() *DataBlock {
	return &this.writeDB
}

func (this *ClientSocketBase) SetDataDecode(ddb *DataDecodeBase) {
	this.dataDecoder = ddb
}

func (this *ClientSocketBase) Connect(host string, port, ttl int32) error {
	this.ttl = ttl

	//通过域名找IP地址
	ip, err := net.ResolveIPAddr("", host)
	if err != nil {
		return err
	}
	ipStr := fmt.Sprintf("%s:20003", ip.IP.String())

	this.Conn, err = net.DialTimeout("tcp", ipStr, time.Second*time.Duration(this.ttl))
	if err != nil {
		return err
	}

	go rector(this)
	return nil
}

func (this *ClientSocketBase) Send(buf [] byte) bool {
	defer this.wMutex.Unlock()
	this.wMutex.Lock()
	return this.writeDB.Append(buf) > 0
}

func rector(client *ClientSocketBase) {
	defer func() {
		err := client.Close()
		if err != nil {
			util.E("close error=%v", err.Error())
		}
		if client.onSocket != nil {
			client.onSocket.Close(client)
		}
	}()

	if client.onSocket == nil || client.dataDecoder == nil {
		return
	}
	client.onSocket.Connect(client)

	for {
		buffer := make([]byte, 1024)
		n, err := client.Read(buffer)
		if err != nil {
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
				client.Status = StatusClosed
				client.onSocket.Close(client)
			} else {
				client.Status = StatusError
				client.Err = err
				client.onSocket.RecvErr(client)
			}
			return
		}

		len := int(client.readDB.Append(buffer))
		if len != n {
			client.Status = StatusError
			client.Err = ErrOOM
			client.onSocket.NetErr(client)
			return
		}

		client.dataDecoder.Process(client)
	}
}

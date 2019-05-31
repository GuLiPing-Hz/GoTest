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

func CheckTimeout(err error) bool {
	if err != nil {
		if err1, ok := err.(*net.OpError); ok {
			return err1.Timeout()
		}
	}
	return false
}

type OnSocket interface {
	/**
	连接上服务器回调
	*/
	Connect(*ClientSocketBase)
	/**
	连接超时,写入超时,读取超时回调
	*/
	Timeout(*ClientSocketBase)
	/**
	服务器主动关闭回调
	*/
	Close(*ClientSocketBase)
	/**
	网络错误回调
	*/
	NetErr(*ClientSocketBase)

	/**
	以goroutine的形式回调
	*/
	RecvMsg(*ClientSocketBase, []byte)
}

type DataDecodeBase interface {
	GetPackageHeadLen() int
	GetPackageLen([]byte) int
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

func (this *DataBlock) GetPos() int {
	return len(this.buffer)
}

func (this *DataBlock) InitPos() {
	this.buffer = make([]byte, 0)
}

type ClientSocketBase struct {
	StatusConn
	readDB      DataBlock
	writeDB     DataBlock
	wMutex      sync.Mutex
	dataDecoder DataDecodeBase
	ttl         time.Duration
	rttl        time.Duration
	onSocket    OnSocket
}

func (this *ClientSocketBase) Connect(host string, port, ttl time.Duration, onSocket OnSocket, ddb DataDecodeBase) error {
	if onSocket == nil {
		return ErrParam
	}

	this.onSocket = onSocket
	this.ttl = ttl
	this.dataDecoder = ddb
	if ddb == nil {
		this.dataDecoder = new(DataDecode)
	}

	//通过域名找IP地址
	ip, err := net.ResolveIPAddr("", host)
	if err != nil {
		return err
	}
	ipStr := fmt.Sprintf("%s:%d", ip.IP.String(), port)
	this.Conn, err = net.DialTimeout("tcp", ipStr, time.Second*this.ttl)
	if err != nil {
		if err1, ok := err.(*net.OpError); ok && err1.Timeout() {
			return ErrTimeout
		}
		return err
	}

	go rector(this)
	return nil
}

func Agent(conn StatusConn, ttl time.Duration, onSocket OnSocket) *ClientSocketBase {
	if onSocket == nil {
		return nil
	}

	csb := &ClientSocketBase{
		StatusConn: conn,
		ttl:        ttl,
		rttl:       ttl,
		onSocket:   onSocket,
	}
	go rector(csb)
	return csb
}

func (this *ClientSocketBase) close() {
	err := this.Close()
	if err != nil {
		util.E("close error=%v", err.Error())
	}
	this.onSocket.Close(this)
}

func (this *ClientSocketBase) send() {
	this.wMutex.Lock()
	buffer := this.writeDB.GetBuf()
	this.writeDB.InitPos()
	this.wMutex.Unlock()

	err := this.SetWriteDeadline(time.Now().Add(time.Second * this.ttl))
	if err != nil {
		this.Status = StatusError
		this.Err = err
		this.onSocket.NetErr(this)
		//把发生错误的socket及时关闭
		this.close()
	}

	_, err = this.Write(buffer)
	if err != nil {
		if CheckTimeout(err) {
			this.Status = StatusTimeout
			this.onSocket.Timeout(this)
		} else {
			this.Status = StatusError
			this.Err = err
			this.onSocket.NetErr(this)
		}
		//把超时的socket及时关闭
		this.close()
		return
	}
}

func (this *ClientSocketBase) Send(buf [] byte) {
	this.wMutex.Lock()
	defer this.wMutex.Unlock()

	this.writeDB.Append(buf)
	go this.send()
}

func (this *ClientSocketBase) process() error {
	headLen := int(this.dataDecoder.GetPackageHeadLen())
	for {
		if this.readDB.GetPos() < headLen { //不足包长
			return nil
		}

		packageLen := this.dataDecoder.GetPackageLen(this.readDB.GetBuf())
		if packageLen == 0 { //异常包
			this.readDB.InitPos()
			return ErrBuffer
		}

		fullLen := headLen + packageLen
		if fullLen < this.readDB.GetPos() { //不足一个包
			return nil
		}

		packageBuf := make([]byte, fullLen)
		copy(packageBuf, this.readDB.GetBuf()[:fullLen])
		this.readDB.Move(this.readDB.GetBuf()[fullLen:])
		go this.onSocket.RecvMsg(this, packageBuf)
	}
}

func rector(client *ClientSocketBase) {
	defer client.close()
	client.onSocket.Connect(client)
	for {
		if client.rttl != 0 { //如果需要判断读超时。
			err := client.SetReadDeadline(time.Now().Add(time.Second * client.rttl))
			if err != nil {
				client.Status = StatusError
				client.Err = err
				client.onSocket.NetErr(client)
				return
			}
		}

		buffer := make([]byte, 1024)
		n, err := client.Read(buffer)
		if err != nil {
			err1 := err.(*net.OpError)
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
			} else if err1 != nil && err1.Timeout() {
				client.Status = StatusTimeout
				client.onSocket.Timeout(client)
			} else {
				client.Status = StatusError
				client.Err = err
				client.onSocket.NetErr(client)
			}
			return
		}

		appendLen := int(client.readDB.Append(buffer))
		if appendLen != n {
			client.Status = StatusError //无法把buffer全部塞进去，多半是没有内存了。
			client.Err = ErrOOM
			client.onSocket.NetErr(client)
			return
		}

		err = client.process()
		if err != nil {
			client.Status = StatusError
			client.Err = err
			client.onSocket.NetErr(client)
			return
		}
	}
}

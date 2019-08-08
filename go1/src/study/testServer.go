package main

import (
	"fmt"
	"io"
	"net"
	"runtime"
	"strings"
	"time"
)

const (
	TimeFmt = "2006/01/02 15:04:05.000" //毫秒保留3位有效数字
)

func Log(format string, args ...interface{}) {
	fmt.Printf(fmt.Sprintf("%s:%s", time.Now().Format(TimeFmt), format), args...)
}

type StatusNO int32

const (
	StatusNormal  StatusNO = iota //0 正常
	StatusClosed                  //1 对方主动关闭
	StatusTimeout                 //2 超时连接断开
	StatusError                   //3 其他异常断开
)

type StatusConn struct {
	net.Conn
	//StatusNormal
	Status StatusNO
}

func handlerClient(conn *StatusConn) {
	defer func() {
		Log("handlerClient Status=%v\n", conn.Status)
		err := conn.Close()
		if err != nil {
			Log("handlerClient err1=%v\n", err)
		}
	}()

	for {
		buffer := make([]byte, 1024)
		err := conn.SetReadDeadline(time.Now().Add(time.Second * 10))
		if err != nil {
			Log("handlerClient err2=%v\n", err.Error())
			conn.Status = StatusError
			return
		}
		n, err := conn.Read(buffer)
		if err != nil {
			Log("handlerClient err3=%v\n", err.Error())
			errStr := err.Error()
			err1, ok := err.(*net.OpError)
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
				conn.Status = StatusClosed
			} else if ok && err1 != nil && err1.Timeout() {
				conn.Status = StatusTimeout
			} else {
				conn.Status = StatusError
			}
			return
		}

		clientBytes := buffer[:n]
		Log("read client [%d]=%s\n", n, string(clientBytes))
		hello := fmt.Sprintf("from server back【go】:%s", string(clientBytes))

		n, err = conn.Write([]byte(hello))
		Log("send client [%d]=%s\n", n, hello)
		if err != nil {
			Log("handlerClient err4=%v\n", err.Error())
		}

	}
}

func main() {
	//服务器例子
	Log("cur goos=%s\n", runtime.GOOS)

	//获取服务器监听ip地址
	ip, err := net.ResolveTCPAddr("", ":20003")
	if err != nil {
		Log("err1=%v\n", err.Error())
		return
	}

	//创建一个监听的socket
	conn, err := net.ListenTCP("tcp", ip)
	if err != nil {
		Log("err2=%v\n", err.Error())
		return
	}

	Log("server wait...\n")
	for {
		clientConn, err := conn.Accept()
		if err != nil {
			Log("err3=%v\n", err.Error())
			return
		}
		Log("new client conn=%v,ip=%v\n", clientConn, clientConn.RemoteAddr())
		go handlerClient(&StatusConn{clientConn, StatusNormal})
	}
}

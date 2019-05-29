package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

const (
	TimeFmt = "2006/01/02 15:04:05.999" //毫秒保留3位有效数字
)

func Log(format string, args ...interface{}) {
	fmt.Printf(fmt.Sprintf("%s:%s", time.Now().Format(TimeFmt), format), args...)
}

type StatusConn struct {
	net.Conn
	/*
		0 正常
		1 对方主动关闭
		2 异常错误关闭
	*/
	Status int32
}

func handlerClient(conn *StatusConn) {
	defer func() {
		Log("handlerClient Status=%v", conn.Status)
		err := conn.Close()
		if err != nil {
			Log("handlerClient err1=%v\n", err)
		}
	}()

	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			Log("handlerClient err2=%v\n", err.Error())
			if err == io.EOF {
				conn.Status = 1
			} else {
				conn.Status = 2
			}
			return
		}

		Log("read client [%d]=%s\n", n, string(buffer))
		hello := fmt.Sprintf("from server back:%s", string(buffer))

		n, err = conn.Write([]byte(hello))
		Log("send client [%d]=%s\n", n, hello)
	}
}

func main() {
	//服务器例子

	//服务器监听ip地址
	ip, err := net.ResolveTCPAddr("", ":20001")
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

	for {
		Log("server wait...\n")
		clientConn, err := conn.Accept()
		if err != nil {
			Log("err3=%v\n", err.Error())
			return
		}

		go handlerClient(&StatusConn{clientConn, 0})
	}
}

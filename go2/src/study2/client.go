package main

import (
	"fmt"
	"net"
	"time"
)

//一个TCP Client连接服务器的例子
const (
	TimeFmt = "2006/01/02 15:04:05.000" //毫秒保留3位有效数字
)

func Log(format string, args ...interface{}) {
	fmt.Printf(fmt.Sprintf("%s:%s", time.Now().Format(TimeFmt), format), args...)
}

func main() {
	//通过域名找IP地址
	ip, err := net.ResolveIPAddr("", "127.0.0.1")
	//ip, err := net.ResolveIPAddr("", "www.fanyu123.cn")
	if err != nil {
		Log("err1=%v\n", err.Error())
		return
	}
	//ipStr := fmt.Sprintf("%s:20003", ip.IP.String())
	ipStr := fmt.Sprintf("%s:9876", ip.IP.String()) //通过gate服务器
	conn, err := net.Dial("tcp", ipStr)
	if err != nil {
		Log("err2=%v\n", err.Error())
		return
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			Log("err3=%v\n", err.Error())
			return
		}
	}()

	Log("connected ...\n")
	hello := "Hi,here is glp"
	n, err := conn.Write([]byte(hello))
	if err != nil {
		Log("err4=%v\n", err.Error())
		return
	}
	Log("send server [%d]=%s\n", n, hello)

	//接收数据
	//buffer := make([]byte, 1024)
	//n, err = conn.Read(buffer)
	//if err != nil {
	//	Log("err5=%v\n", err.Error())
	//	return
	//}
	//Log("read server [%d]=%s\n", n, string(buffer))
	//
	<-time.After(time.Second)
}

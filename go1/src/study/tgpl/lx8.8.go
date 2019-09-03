package main

import (
	"fmt"
	"net"
	"time"
)

func handleClient(conn net.Conn) {
	chanBuf := make(chan []byte)
	go func() {
		for {
			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Printf("err=%s\n", err.Error())
				return
			}

			//fmt.Printf("read from for %s\n", string(buf[:n]))
			chanBuf <- buf[:n]
		}
	}()

	defer conn.Close()
	for {
		chanClock := time.After(time.Second * 10)
		select {
		case <-chanClock:
			fmt.Printf("time out \n")
			return
		case buf := <-chanBuf:
			fmt.Printf("client:%s \n", string(buf))
			//default: default属于轮询，每次都会执行
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Printf("err=%s\n", err.Error())
		return
	}

	for {
		con, err := listener.Accept()
		if err != nil {
			fmt.Printf("err=%s\n", err.Error())
			return
		}

		go handleClient(con)
	}
}

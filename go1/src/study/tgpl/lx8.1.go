package main

import (
	"net"
	"flag"
	"fmt"
	"bytes"
	"time"
)

func replyTime(conn net.Conn) {
	buf := &bytes.Buffer{}
	for {
		buf.Reset()
		fmt.Fprintf(buf, "time = %s\r\n", time.Now().
			Format("2006-01-02 15:04:05")) //time.RFC3339))
		_, err := conn.Write(buf.Bytes())
		if err != nil {
			fmt.Printf("client err=%s\n", err.Error())
			return
		}
		time.Sleep(time.Second)
	}
}

func main() {
	var port string
	flag.StringVar(&port, "port", "8000", "port")

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Printf("err=%s", err.Error())
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("err=%s", err.Error())
			return
		}

		go replyTime(conn)
	}
}

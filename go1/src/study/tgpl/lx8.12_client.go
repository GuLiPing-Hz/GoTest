package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

var address string

func main() {
	flag.StringVar(&address, "addr", "127.0.0.1:8000", "set the connect address")
	flag.Parse()

	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Printf("err=%s\n", err.Error())
		return
	}

	defer conn.Close()
	chanMsg := make(chan string)
	defer close(chanMsg)
	go func() {
		//chanMsg <- "123"

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			//fmt.Printf("enter %s\n", scanner.Text())
			chanMsg <- scanner.Text()
		}
	}()

	chanStop := make(chan struct{})
	defer close(chanStop)
	go func() {
		for {
			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Printf("err=%s\n", err.Error())
				chanStop <- struct{}{}
				return
			}
			fmt.Printf("from server:\n%s\n", string(buf[:n]))
		}
	}()

	for {
		select {
		case msg, ok := <-chanMsg:
			if !ok {
				return
			}
			_, err := fmt.Fprintln(conn, msg)
			if err != nil {
				fmt.Printf("err=%s\n", err.Error())
				return
			}
		case <-chanStop:
			fmt.Println("stop")
			return
		}
	}
}

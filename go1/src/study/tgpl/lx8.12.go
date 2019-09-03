package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

type Client chan string

var chanEnter = make(chan Client)
var chanExit = make(chan Client)
var chanCenter = make(chan string)

func handleClient2(conn net.Conn) {
	client := make(chan string)
	go func() {
		for msg := range client {
			fmt.Fprintf(conn, "%s", msg)
		}
	}()

	me := "[" + conn.RemoteAddr().String() + "]"
	client <- "You are " + me
	chanEnter <- client
	chanCenter <- me + " enter"

	scanner := bufio.NewScanner(conn)
	chanIdleDetact := make(chan string)
	conn2 := conn
	defer close(chanIdleDetact)
	go func() {
		for {
			idleClock := time.After(time.Second * 10)
			select {
			case msg := <-chanIdleDetact:
				chanCenter <- msg
			case <-idleClock:
				conn.Close()
				conn2 = nil
				return
			}
		}
	}()

	for scanner.Scan() {
		chanIdleDetact <- fmt.Sprintf("%s say: %s", me, scanner.Text())
	}

	chanCenter <- me + " exit"
	chanExit <- client
	close(client)

	if conn2 != nil {
		conn2.Close()
	}
}

func bocardcast() {
	clients := make(map[Client]bool)
	for {
		select {
		case enter := <-chanEnter:
			clients[enter] = true
		case exit := <-chanExit:
			delete(clients, exit)
		case msg := <-chanCenter:
			fmt.Printf("%s\n", msg)
			for k := range clients {
				go func(c Client) {
					c <- msg
				}(k)
			}
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Printf("err=%s\n", err.Error())
		return
	}

	go bocardcast()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("err=%s\n", err.Error())
			return
		}

		go handleClient2(conn)
	}
}

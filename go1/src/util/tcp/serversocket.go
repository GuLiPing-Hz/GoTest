package tcp

import (
	"fmt"
	"net"
	"sync"
	"time"
	"util"
)

type OnSocketServer interface {
	OnSocket
	ServerListen(*ServerSocketBase)
	ServerErr(*ServerSocketBase)
	ServerClose(*ServerSocketBase)
}

type ServerSocketBase struct {
	StatusConnServer
	onSocket OnSocketServer
	ttl      time.Duration

	fd2Client sync.Map //fd -> *ClientSocketBase
}

func (this *ServerSocketBase) Listen(host string, port uint16, onSocket OnSocketServer) error {
	if onSocket == nil {
		return ErrParam
	}
	this.onSocket = onSocket
	//获取服务器监听ip地址
	address := fmt.Sprintf("%s:%d", host, port)
	ip, err := net.ResolveTCPAddr("", address)
	if err != nil {
		return err
	}

	//创建一个监听的socket
	this.listener, err = net.ListenTCP("tcp", ip)
	if err != nil {
		return err
	}

	go reactor(this)
	return nil
}

func (this *ServerSocketBase) close() {
	err := this.listener.Close()
	if err != nil {
		util.E("close error=%v", err.Error())
	}
	this.onSocket.ServerClose(this)
}

func reactor(server *ServerSocketBase) {
	defer server.close()

	server.onSocket.ServerListen(server)
	for {
		clientConn, err := server.listener.Accept()
		if err != nil {
			server.Status = StatusError
			server.Err = err
			server.onSocket.ServerErr(server)
			return
		}

		agent := Agent(StatusConn{clientConn, StatusErr{}}, server.ttl, server.onSocket)
		server.fd2Client.Store(agent, true)
	}
}

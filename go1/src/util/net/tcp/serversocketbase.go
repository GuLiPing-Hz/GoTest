package tcp

import (
	"fmt"
	"net"
	"sync"
	"time"
	"util"
	net2 "util/net"
)

type StatusConnServer struct {
	listener *net.TCPListener
	net2.StatusErr
}

type OnSocketServer interface {
	net2.OnSocket
	OnServerListen(*ServerSocketBase)
	OnServerErr(*ServerSocketBase)
	//这里OnServerClose始终会回调
	OnServerClose(*ServerSocketBase)
}

type ServerSocketBase struct {
	StatusConnServer
	onSocket   OnSocketServer
	ttl        time.Duration
	chanAccept chan net.Conn
	chanStop   chan bool

	fd2Client sync.Map //fd -> *ClientBase
}

func (this *ServerSocketBase) Listen(host string, port uint16, onSocket OnSocketServer) error {
	if onSocket == nil {
		return net2.ErrParam
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
		util.E("Close error=%v", err.Error())
	}
	this.onSocket.OnServerClose(this)
}

func reactor(server *ServerSocketBase) {
	defer server.close()

	server.onSocket.OnServerListen(server)
	go func() {
		clientConn, err := server.listener.Accept()
		if err != nil {
			server.Status = net2.StatusError
			server.Err = err
			server.onSocket.OnServerErr(server)
			return
		}
		server.chanAccept <- clientConn
	}()
	for {
		select {
		case clientConn := <-server.chanAccept:
			agent := Agent(clientConn, server.ttl, server.onSocket)
			server.fd2Client.Store(agent, true)
		case <-server.chanStop:
			server.fd2Client.Range(func(key, value interface{}) bool {
				agent := key.(*ClientBaseSocket)
				agent.SafeClose(true)
				return true
			})
			break
		}
	}
}

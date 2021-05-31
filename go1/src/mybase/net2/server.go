package net2

import (
	"fmt"
	"mybase"
	"net"
	"sync"
	"time"
)

type StatusConnServer struct {
	listener *net.TCPListener
	StatusErr
}

type StackError interface {
	error
	Stack() []byte
}

type OnSocketServer interface {
	OnSocket
	OnServerListen()
	OnServerErr(StackError)
	//这里OnServerClose始终会回调
	OnServerClose()
}

type ServerSocket struct {
	StatusConnServer
	onSocket OnSocketServer
	ttl      time.Duration //监听客户端读取超时时间，如果不需要有超时机制，可以设置为0
	rTtl     time.Duration

	chanAccept chan net.Conn
	chanStop   chan bool

	fd2Client sync.Map //net2.Conn -> true

	clientDataDecoder DataDecodeBase
	listenAddress     string
}

/**
连接上服务器回调
*/
func (imp *ServerSocket) OnConnect(con Conn) {
	imp.fd2Client.Store(con, true)
	imp.onSocket.OnConnect(con)
}

/**
只要我们曾经连接上服务器过，OnClose必定会回调。代表一个当前的socket已经关闭
*/
func (imp *ServerSocket) OnClose(con Conn) {
	imp.fd2Client.Delete(con)
	imp.onSocket.OnClose(con)
}

/**
连接超时,写入超时,读取超时回调，之后会调用OnClose
*/
func (imp *ServerSocket) OnTimeout(con Conn) {
	imp.onSocket.OnTimeout(con)
}

/**
网络错误回调，之后会调用OnClose
*/
func (imp *ServerSocket) OnNetErr(con Conn) {
	imp.onSocket.OnNetErr(con)
}

/**
接受到信息
@return 返回true表示可以继续热恋，false表示要分手了。
*/
func (imp *ServerSocket) OnRecvMsg(con Conn, buf []byte) bool {
	return imp.onSocket.OnRecvMsg(con, buf)
}

func (imp *ServerSocket) Listen() error {
	if imp.onSocket == nil {
		return mybase.ErrParam
	}
	//获取服务器监听ip地址
	ip, err := net.ResolveTCPAddr("", imp.listenAddress)
	if err != nil {
		return err
	}

	//创建一个监听的socket
	imp.listener, err = net.ListenTCP("tcp", ip)
	if err != nil {
		return err
	}

	go reactor(imp)
	return nil
}

func (imp *ServerSocket) Shutdown() {
	//停止服务器运行。
	imp.chanStop <- true
}

func (imp *ServerSocket) close() {
	err := imp.listener.Close()
	if err != nil {
		mybase.E("Close error=%v", err.Error())
	}
	imp.onSocket.OnServerClose()
}

func reactor(server *ServerSocket) {
	defer server.close()

	server.onSocket.OnServerListen()
	go func() {
		for {
			clientConn, err := server.listener.Accept()
			if err != nil {
				server.ChangeStatus(StatusError, err)
				server.onSocket.OnServerErr(server)
				return
			}
			server.chanAccept <- clientConn
		}
	}()

loop:
	for {
		select {
		case clientConn := <-server.chanAccept:
			//agent 可以考虑弄个代理池，或许更高效一点
			agent := Agent(clientConn, server.ttl, server.rTtl, server, server.clientDataDecoder)
			if agent != nil {
				agent.Rector()
			}
		case <-server.chanStop:
			//关闭已经连接的
			server.fd2Client.Range(func(key, value interface{}) bool {
				agent := key.(*ClientBase)
				agent.SafeClose()
				return true
			})
			//关闭监听的socket
			_ = server.listener.Close()
			break loop
		}
	}
}

func NewServerIp(ip string, port uint16, onSocket OnSocketServer, clientDDB DataDecodeBase) *ServerSocket {
	return NewServer(fmt.Sprintf("%s:%d", ip, port), time.Second*30, time.Second*30, onSocket, clientDDB)
}

/**
@ttl 客户端发送超时
@rTtl 客户端读取超时 0表示永远等待读取。
*/
func NewServer(address string, ttl time.Duration, rTtl time.Duration, onSocket OnSocketServer,
	clientDDB DataDecodeBase) *ServerSocket {
	ssb := &ServerSocket{}
	ssb.listenAddress = address
	ssb.onSocket = onSocket
	ssb.chanStop = make(chan bool)
	ssb.chanAccept = make(chan net.Conn)
	ssb.clientDataDecoder = clientDDB
	ssb.ttl = ttl
	ssb.rTtl = rTtl
	return ssb
}

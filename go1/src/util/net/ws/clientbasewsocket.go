package ws

import (
	"github.com/gorilla/websocket"
	"net"
	"time"
	net2 "util/net"
)

type ClientBaseWSocket struct {
	net2.ClientBase

	msgType int
	conn    *websocket.Conn
}

func Agent(conn *websocket.Conn, msgType int, ttl time.Duration, OnSocket net2.OnSocket) *ClientBaseWSocket {
	if OnSocket == nil {
		return nil
	}

	csb := &ClientBaseWSocket{
		ClientBase: net2.ClientBase{
			Ttl:      ttl,
			Rttl:     ttl,
			OnSocket: OnSocket,
		},
		msgType: msgType,
		conn:    conn,
	}
	csb.ConnBase = csb
	csb.ConnInner = csb
	csb.Rector()
	return csb
}

func (this *ClientBaseWSocket) Close() error {
	return this.conn.Close()
}

func (this *ClientBaseWSocket) SendEx(buffer []byte) {
	err := this.conn.SetWriteDeadline(time.Now().Add(this.Ttl))
	if err != nil {
		this.CloseWithErr(err)
	}

	err = this.conn.WriteMessage(this.msgType, buffer)
	//packageLen := binary.BigEndian.Uint16(buffer)
	//util.D("send buffer %d==%d,buf=%x", n, packageLen, buffer)
	if err != nil {
		this.CloseWithErr(err)
		return
	}
}

// LocalAddr returns the local network address.
func (this *ClientBaseWSocket) LocalAddr() net.Addr {
	return this.conn.LocalAddr()
}

// RemoteAddr returns the remote network address.
func (this *ClientBaseWSocket) RemoteAddr() net.Addr {
	return this.conn.RemoteAddr()
}

func (this *ClientBaseWSocket) RecvEx() ([]byte, error) {
	if this.Rttl != 0 { //如果需要判断读超时。
		err := this.conn.SetReadDeadline(time.Now().Add(time.Second * this.Rttl))
		if err != nil {
			return nil, err
		}
	}

	_, buffer, err := this.conn.ReadMessage()
	if err != nil {
		return buffer, nil
	}

	return nil, err
}

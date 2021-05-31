package net2

import (
	"github.com/gorilla/websocket"
	"net"
	"time"
)

type ClientWSocket struct {
	ClientBase
	msgType int //TextMessage or BinaryMessage
	conn    *websocket.Conn
}

func (imp *ClientWSocket) Close() error {
	return imp.conn.Close()
}

// LocalAddr returns the local network address.
func (imp *ClientWSocket) LocalAddr() net.Addr {
	return imp.conn.LocalAddr()
}

// RemoteAddr returns the remote network address.
func (imp *ClientWSocket) RemoteAddr() net.Addr {
	return imp.conn.RemoteAddr()
}

func (imp *ClientWSocket) SendEx(buffer []byte) {
	err := imp.conn.SetWriteDeadline(time.Now().Add(imp.Ttl))
	if err != nil {
		imp.CloseWithErr(err)
	}

	err = imp.conn.WriteMessage(imp.msgType, buffer)
	//packageLen := binary.BigEndian.Uint16(buffer)
	//mybase.D("send buffer %d==%d,buf=%x", n, packageLen, buffer)
	if err != nil {
		imp.CloseWithErr(err)
		return
	}
}

func (imp *ClientWSocket) RecvEx() ([]byte, error) {
	if imp.RTtl != 0 { //如果需要判断读超时。
		err := imp.conn.SetReadDeadline(time.Now().Add(imp.RTtl))
		if err != nil {
			return nil, err
		}
	}

	_, buffer, err := imp.conn.ReadMessage()
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

func WebAgent(conn *websocket.Conn, msgType int, ttl time.Duration, rTtl time.Duration, OnSocket OnSocket, ddb DataDecodeBase) *ClientWSocket {
	if OnSocket == nil {
		return nil
	}

	if ddb == nil {
		ddb = new(DataDecode)
	}
	csb := &ClientWSocket{
		msgType: msgType,
		conn:    conn,
	}
	csb.Init(ddb, ttl, rTtl, OnSocket, csb)
	csb.Rector()
	return csb
}

package net2

import (
	"fmt"
	"mybase"
	"net"
	"time"
)

//*********ClientSocket
func CheckTimeout(err error) bool {
	if err != nil {
		if err1, ok := err.(*net.OpError); ok {
			return err1.Timeout()
		}
	}
	return false
}

type ClientSocket struct {
	ClientBase
	conn net.Conn
}

func (imp *ClientSocket) Close() error {
	return imp.conn.Close()
}

// LocalAddr returns the local network address.
func (imp *ClientSocket) LocalAddr() net.Addr {
	return imp.conn.LocalAddr()
}

// RemoteAddr returns the remote network address.
func (imp *ClientSocket) RemoteAddr() net.Addr {
	return imp.conn.RemoteAddr()
}

func (imp *ClientSocket) SendEx(buffer []byte) {
	if len(buffer) == 0 {
		return
	}

	err := imp.conn.SetWriteDeadline(time.Now().Add(imp.Ttl))
	if err != nil {
		imp.CloseWithErr(err)
	}

	n, err := imp.conn.Write(buffer)
	fmt.Printf("SendEx buf[%d]\n", n)
	if err != nil {
		if CheckTimeout(err) {
			imp.CloseTimeout()
		} else {
			imp.CloseWithErr(err)
		}
		return
	}
}

func (imp *ClientSocket) RecvEx() ([]byte, error) {
	if imp.RTtl != 0 { //如果需要判断读超时。
		err := imp.conn.SetReadDeadline(time.Now().Add(imp.RTtl))
		if err != nil {
			return nil, err
		}
	}

	buffer := make([]byte, 2048)
	n, err := imp.conn.Read(buffer)
	if err != nil {
		return nil, err
	}

	return buffer[:n], nil
}

func (imp *ClientSocket) ConnectHostPort(host string, port uint16, Ttl time.Duration, OnSocket OnSocket, ddb DataDecodeBase) error {
	//通过域名找IP地址
	ip, err := net.ResolveIPAddr("", host)
	if err != nil {
		return err
	}
	var addr = fmt.Sprintf("%s:%d", ip.IP.String(), port)

	return imp.Connect(addr, Ttl, OnSocket, ddb)
}

func (imp *ClientSocket) Connect(addr string, ttl time.Duration, OnSocket OnSocket, ddb DataDecodeBase) error {
	if OnSocket == nil {
		return mybase.ErrParam
	}

	imp.Init(ddb, ttl, 0, OnSocket, imp)

	return imp.ReConnect(addr)
}

func (imp *ClientSocket) ReConnect(addr string) error {
	var err error
	imp.conn, err = net.DialTimeout("tcp", addr, imp.Ttl)
	if err != nil {
		if err1, ok := err.(*net.OpError); ok && err1.Timeout() {
			return mybase.ErrTimeout
		}
		return err
	}

	imp.Rector()
	return nil
}

func Agent(conn net.Conn, ttl time.Duration, rTtl time.Duration, OnSocket OnSocket,
	ddb DataDecodeBase) *ClientSocket {
	if OnSocket == nil {
		return nil
	}

	if ddb == nil {
		ddb = new(DataDecode)
	}

	csb := &ClientSocket{
		conn: conn,
	}
	csb.Init(ddb, ttl, rTtl, OnSocket, csb)
	return csb
}

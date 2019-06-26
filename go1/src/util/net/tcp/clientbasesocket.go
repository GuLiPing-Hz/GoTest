package tcp

import (
	"bytes"
	"fmt"
	"net"
	"time"
	net2 "util/net"
)

func CheckTimeout(err error) bool {
	if err != nil {
		if err1, ok := err.(*net.OpError); ok {
			return err1.Timeout()
		}
	}
	return false
}

type ClientBaseSocket struct {
	net2.ClientBase
	conn net.Conn
}

func (this *ClientBaseSocket) ConnectHostPort(host string, port uint16, Ttl time.Duration, OnSocket net2.OnSocket) error {
	//通过域名找IP地址
	ip, err := net.ResolveIPAddr("", host)
	if err != nil {
		return err
	}
	var addr = fmt.Sprintf("%s:%d", ip.IP.String(), port)

	return this.Connect(addr, Ttl, OnSocket, nil)
}

func (this *ClientBaseSocket) Connect(addr string, Ttl time.Duration, OnSocket net2.OnSocket, ddb net2.DataDecodeBase) error {
	if OnSocket == nil {
		return net2.ErrParam
	}

	this.OnSocket = OnSocket
	this.Ttl = Ttl
	this.DataDecoder = ddb
	this.ReadDB = &bytes.Buffer{}
	this.WriteDB = &bytes.Buffer{}
	if ddb == nil {
		this.DataDecoder = new(net2.DataDecode)
	}

	return this.ReConnect(addr)
}

func (this *ClientBaseSocket) ReConnect(addr string) error {
	if this.OnSocket == nil {
		return net2.ErrParam
	}

	var err error
	this.conn, err = net.DialTimeout("tcp", addr, this.Ttl)
	if err != nil {
		if err1, ok := err.(*net.OpError); ok && err1.Timeout() {
			return net2.ErrTimeout
		}
		return err
	}

	this.Rector()
	return nil
}

func Agent(conn net.Conn, Ttl time.Duration, OnSocket net2.OnSocket) *ClientBaseSocket {
	if OnSocket == nil {
		return nil
	}

	csb := &ClientBaseSocket{
		ClientBase: net2.ClientBase{
			ReadDB:      &bytes.Buffer{},
			WriteDB:     &bytes.Buffer{},
			DataDecoder: new(net2.DataDecode),
			Ttl:         Ttl,
			Rttl:        Ttl,
			OnSocket:    OnSocket,
		},
		conn: conn,
	}
	csb.ConnBase = csb //复制接口
	csb.ConnInner = csb
	csb.Rector()
	return csb
}

func (this *ClientBaseSocket) Close() error {
	return this.conn.Close()
}

func (this *ClientBaseSocket) SendEx(buffer []byte) {
	err := this.conn.SetWriteDeadline(time.Now().Add(this.Ttl))
	if err != nil {
		this.CloseWithErr(err)
	}

	_, err = this.conn.Write(buffer)
	//packageLen := binary.BigEndian.Uint16(buffer)
	//util.D("send buffer %d==%d,buf=%x", n, packageLen, buffer)
	if err != nil {
		if CheckTimeout(err) {
			this.CloseTimeout()
		} else {
			this.CloseWithErr(err)
		}
		return
	}
}

// LocalAddr returns the local network address.
func (this *ClientBaseSocket) LocalAddr() net.Addr {
	return this.conn.LocalAddr()
}

// RemoteAddr returns the remote network address.
func (this *ClientBaseSocket) RemoteAddr() net.Addr {
	return this.conn.RemoteAddr()
}

func (this *ClientBaseSocket) RecvEx() ([]byte, error) {
	if this.Rttl != 0 { //如果需要判断读超时。
		err := this.conn.SetReadDeadline(time.Now().Add(this.Rttl))
		if err != nil {
			return nil, err
		}
	}

	buffer := make([]byte, 1024)
	n, err := this.conn.Read(buffer)
	if err != nil {
		return nil, err
	}

	return buffer[:n], nil
}

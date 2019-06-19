package net

import (
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"sync"
)

type StatusNO int32

const (
	StatusUnknown  StatusNO = iota //0 尚未初始化
	StatusNormal                   //1 正常
	StatusShutdown                 //2 自己关闭的
	StatusTimeout                  //3 超时连接断开
	StatusError                    //4 其他异常断开,服务器主动断开 err为ErrClose
)

var (
	ErrParam   = errors.New("param error")
	ErrBuffer  = errors.New("Server buffer error")
	ErrInner   = errors.New("Server inner error")
	ErrOOM     = errors.New("Server oom")
	ErrTimeout = errors.New("Server time out")
	ErrClose   = errors.New("Closed by the peer")
)

type StatusErr struct {
	Status      StatusNO //StatusNormal
	Err         error    //当为其他异常时，这里会有赋值
	Stack       []byte   //调用的错误堆栈信息
	Mutex       sync.Mutex
	IsConnected int32 //0未连接，1已连接
}

func (this *StatusErr) GetStatus() StatusNO {
	this.Mutex.Lock()
	defer this.Mutex.Unlock()
	return this.Status
}

func (this *StatusErr) ChangeStatus(status StatusNO, err error) bool {
	return this.ChangeStatusAll(status, err, nil)
}

func (this *StatusErr) ChangeStatusAll(status StatusNO, err error, stack []byte) bool {
	this.Mutex.Lock()
	defer this.Mutex.Unlock()
	if this.Status == StatusNormal || this.Status == StatusUnknown || status == StatusUnknown {
		this.Status = status
		this.Err = err
		this.Stack = stack
		return true
	}
	return false
}

type StatusConnw struct {
	Conn *websocket.Conn
	StatusErr
}

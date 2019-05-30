package tcp

import (
	"github.com/pkg/errors"
	"net"
)

type StatusNO int32

const (
	StatusNormal  StatusNO = iota //0 正常
	StatusClosed                  //1 对方主动关闭
	StatusTimeout                 //2 超时连接断开
	StatusError                   //3 其他异常断开
)

var (
	ErrBuffer = errors.New("Server buffer error")
	ErrInner  = errors.New("Server inner error")
	ErrOOM    = errors.New("Server oom")
)

type StatusConn struct {
	net.Conn
	Status StatusNO //StatusNormal
	Err    error    //当为其他异常时，这里会有赋值
}

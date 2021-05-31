package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"go1/src/mybase/net2"
	"net/http"
	"os"
	"time"
)

//const (
//	TimeFmt = "2006/01/02 15:04:05.000" //毫秒保留3位有效数字
//)

var (
	upgrader = websocket.Upgrader{
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func Log2(format string, args ...interface{}) {
	fmt.Printf(fmt.Sprintf("%s:%s", time.Now().Format("2006/01/02 15:04:05.000"), format), args...)
}

type DataDecodeWSText struct {
}

func (this *DataDecodeWSText) GetPackageHeadLen() int {
	return 0
}
func (this *DataDecodeWSText) GetPackageLen(buf []byte) int {
	return len(buf)
}

type wsOnSocket struct {
}

/**
连接上服务器回调
*/
func (this *wsOnSocket) OnConnect(client net2.Conn) {
	Log2("OnConnect web client[%v]\n", client)
}

/**
连接超时,写入超时,读取超时回调
*/
func (this *wsOnSocket) OnTimeout(client net2.Conn) {
	Log2("OnTimeout web client[%v]\n", client)
}

/**
服务器主动关闭回调
*/
func (this *wsOnSocket) OnClose(client net2.Conn) {
	Log2("OnClose web client[%v]\n", client)
}

/**
网络错误回调
*/
func (this *wsOnSocket) OnNetErr(client net2.Conn) {
	if client.IsClosedByPeer() {
		Log2("OnNetErr web client[%v] closed\n", client)
	} else {
		Log2("OnNetErr web client[%v] msg=%s\n", client, client.Error())
		os.Stderr.Write(client.Stack())
	}
}

/**
接受到信息
*/
func (this *wsOnSocket) OnRecvMsg(client net2.Conn, buf []byte) bool {
	Log2("OnRecvMsg web client[%v] msg=%s\n", client, string(buf))
	hello := fmt.Sprintf("from server back[go]:%s", string(buf))
	client.Send([]byte(hello))
	return true
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	// 完成ws协议的握手操作
	// Upgrade:websocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	agent := net2.WebAgent(conn, websocket.TextMessage, time.Second*10, time.Second*10, new(wsOnSocket), new(DataDecodeWSText))
	Log2("new web client agent=%v\n", agent)
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe("0.0.0.0:20004", nil)
}

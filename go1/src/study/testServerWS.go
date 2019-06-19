package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"os"
	"time"
	"util"
	"util/net"
	"util/net/ws"
)

var (
	upgrader = websocket.Upgrader{
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type wsOnSocket struct {
}

/**
连接上服务器回调
*/
func (this *wsOnSocket) OnConnect(client net.Conn) {
	util.I("OnConnect web client[%v]", client)
}

/**
连接超时,写入超时,读取超时回调
*/
func (this *wsOnSocket) OnTimeout(client net.Conn) {
	util.I("OnTimeout web client[%v]", client)
}

/**
服务器主动关闭回调
*/
func (this *wsOnSocket) OnClose(client net.Conn) {
	util.I("OnClose web client[%v]", client)
}

/**
网络错误回调
*/
func (this *wsOnSocket) OnNetErr(client net.Conn) {
	util.E("OnNetErr web client[%v] msg=%s", client, client.LastErr())
	os.Stderr.Write(client.LastStack())
}

/**
接受到信息
*/
func (this *wsOnSocket) OnRecvMsg(client net.Conn, buf []byte) bool {

	util.I("OnRecvMsg web client[%v] msg=%s", client, string(buf))
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

	agent := ws.Agent(conn, websocket.TextMessage, time.Second*10, new(wsOnSocket))
	util.I("new web client agent=%v", agent)
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe("0.0.0.0:20004", nil)
}

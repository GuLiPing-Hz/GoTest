package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"time"
)

//一个Websocket Client连接服务器的例子
// 搭配Nginx 支持wss
/*
#nginx代理websocket wss的一个具体配置例子。
location /wstest {
	#proxy_redirect off; #websocket代理不能用off
	proxy_read_timeout 60s;
	proxy_set_header Host $host;
	proxy_set_header X-Real_IP $remote_addr;
	proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;

	#websocket必须配置配置
	proxy_http_version 1.1;
	proxy_set_header Upgrade $http_upgrade;
	proxy_set_header Connection 'Upgrade';

	proxy_pass http://127.0.0.1:20004/ws; #等于直接访问 ws://127.0.0.1:20004/ws;
}
*/

//html页面实现的一个websocket,部署到nginx即可使用。
/*
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script>
window.addEventListener("load", function(evt) {

    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var ws;

    var print = function(message) {
        var d = document.createElement("div");
        d.innerHTML = message;
        output.appendChild(d);
    };

    document.getElementById("open").onclick = function(evt) {
        if (ws) {
            return false;
        }
		ws = new WebSocket("wss://www.fanyu123.cn/wstest");
        //ws = new WebSocket("{{.}}");
        ws.onopen = function(evt) {
            print("OPEN");
        }
        ws.onclose = function(evt) {
            print("CLOSE");
            ws = null;
        }
        ws.onmessage = function(evt) {
            print("RESPONSE: " + evt.data);
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    };

    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        print("SEND: " + input.value);
        ws.send(input.value);
        return false;
    };

    document.getElementById("close").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };

});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>Click "Open" to create a connection to the server,
"Send" to send a message to the server and "Close" to close the connection.
You can change the message and send multiple times.
<p>
<form>
<button id="open">Open</button>
<button id="close">Close</button>
<p><input id="input" type="text" value="Hello world!">
<button id="send">Send</button>
</form>
</td><td valign="top" width="50%">
<div id="output"></div>
</td></tr></table>
</body>
</html>
*/

func main() {
	u := url.URL{Scheme: "ws", Host: "127.0.0.1:20004", Path: "/ws"}
	url := "wss://www.fanyu123.cn/wstest"
	log.Printf("connecting to %s||%s", url, u.String())

	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	done := make(chan bool)
	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	for {
		select {
		case <-done:
			break
		case t := <-ticker.C:
			fmt.Printf("write s=%s\n", t.String())
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write:", err)
				return
			}
		}
	}
}

package main

import (
	"context"
	"fmt"
	"go1/src/study/RPCFirst"
	"google.golang.org/grpc"
	"net"
	"time"
)

/**
按道理我们应该去这个网站下代码,
https://github.com/grpc/grpc-go
也就是 go get github.com/grpc/grpc-go
但是好像临近国庆节国外网站貌似都不容易访问了。

我们还是从码云的镜像直接下载吧
go get gitee.com/mirrors/grpc-go
然后去 $GoPATH 修改目录名把  gitee.com/mirrors/grpc-go 改成 google.golang.org/grpc
go get github.com/googleapis/go-genproto
	修改目录名把  github.com/googleapis/go-genproto 改成 google.golang.org/genproto

然后对着proto文件再次编译生成gRPC文件
protoc --go_out=plugins=grpc:. PbLobbyData.proto

生成c++代码的gRPC文件
protoc --grpc_out=cpp --plugin=protoc-gen-grpc=D:\sdk\protoc\bin\grpc_cpp_plugin.exe rpcfirst.proto
protoc --cpp_out=./cpp rpcfirst.proto

*/

type RPCServer struct {
}

//RPC普通方法，一次调用一次返回
func (imp *RPCServer) SayHello(ctx context.Context, r *RPCFirst.ReqHello) (*RPCFirst.RespHello, error) {
	fmt.Printf("context : %+v\n", ctx)

	resp := new(RPCFirst.RespHello)
	resp.Hi = fmt.Sprintf("[go] Hi %s,you are %d", r.Name, r.Age)

	return resp, nil
}

//RPC 一次请求，流式返回
func (imp *RPCServer) LotsOfReplies(r *RPCFirst.ReqHello, w RPCFirst.Demo_LotsOfRepliesServer) error {
	index := 0
	//go 上下文环境导致我们不能切换gorountine，只能在这个goroutine里面执行逻辑代码
	for {
		resp := new(RPCFirst.RespHello)
		resp.Hi = fmt.Sprintf("[go] Hi %s,you are %d %d", r.Name, r.Age, index)
		err := w.Send(resp)
		if err != nil {
			fmt.Printf("err=%s\n", err.Error())
		}
		index++

		if index > 10 {
			break
		}
		time.Sleep(time.Second)
	}
	return nil
}

//RPC 流式请求，一次返回
func (imp *RPCServer) LotsOfGreetings(RPCFirst.Demo_LotsOfGreetingsServer) error {
	return fmt.Errorf("not implemented")
}

//RPC 流式请求，流式返回
func (imp *RPCServer) BidiHello(RPCFirst.Demo_BidiHelloServer) error {
	return fmt.Errorf("not implemented")
}

var GSvr = RPCServer{}

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Printf("err=%v\n", err.Error())
		return
	}

	svr := grpc.NewServer()
	RPCFirst.RegisterDemoServer(svr, &GSvr)
	err = svr.Serve(listener)
	if err != nil {
		fmt.Printf("err=%v\n", err.Error())
	}
}

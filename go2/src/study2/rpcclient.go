package main

import (
	"context"
	"fmt"
	"go2/src/study2/RPCFirst"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8000", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("err=%s\n", err.Error())
		return
	}
	defer conn.Close()

	client := RPCFirst.NewDemoClient(conn)

	resp, err := client.SayHello(context.Background(), &RPCFirst.ReqHello{Name: "Jack", Age: 18})
	if err != nil {
		fmt.Printf("err=%s\n", err.Error())
		return
	}

	fmt.Printf("resp=%s\n", resp.Hi)

	client2, err := client.LotsOfReplies(context.Background(), &RPCFirst.ReqHello{Name: "Gu", Age: 20})
	if err != nil {
		fmt.Printf("err=%s\n", err.Error())
		return
	}

	for {
		resp, err := client2.Recv()
		if err != nil {
			fmt.Printf("err=%s\n", err.Error())
			break
		} else {
			fmt.Printf("resp=%s\n", resp.Hi)
		}
	}
}

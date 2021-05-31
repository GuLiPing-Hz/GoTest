package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
	"strconv"
	"sync"
	"time"
)

func testGetEtcd(cli *clientv3.Client, key string) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		fmt.Printf("testGetEtcd err=%s\n", err.Error())
		return
	}
	fmt.Printf("testGetEtcd.Header=%+v\n", resp.Header)
	fmt.Printf("testGetEtcd.Kvs=%+v\n", resp.Kvs)
	fmt.Printf("testGetEtcd.More=%+v\n", resp.More)
	fmt.Printf("testGetEtcd.Count=%+v\n", resp.Count)
}

/**

 */
func testGetEtcdWithPrefix(cli *clientv3.Client, key string) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := cli.Get(ctx, key, clientv3.WithPrefix())
	cancel()
	if err != nil {
		fmt.Printf("testGetEtcdWithPrefix err=%s\n", err.Error())
		return
	}
	fmt.Printf("testGetEtcdWithPrefix.Header=%+v\n", resp.Header)
	fmt.Printf("testGetEtcdWithPrefix.Kvs=%+v\n", resp.Kvs)
	fmt.Printf("testGetEtcdWithPrefix.More=%+v\n", resp.More)
	fmt.Printf("testGetEtcdWithPrefix.Count=%+v\n", resp.Count)
}

func testPutEtcd(cli *clientv3.Client, key, value string) {
	//写
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := cli.Put(ctx, key, value)
	cancel()
	if err != nil {
		fmt.Printf("testPutEtcd err=%s\n", err.Error())
		return
	}
	fmt.Printf("testPutEtcd=%+v\n", resp)
}

func testDeleteEtcd(cli *clientv3.Client, key string) {
	//写
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := cli.Delete(ctx, key)
	cancel()
	if err != nil {
		fmt.Printf("testDeleteEtcd err=%s\n", err.Error())
		return
	}

	fmt.Printf("testDeleteEtcd.Header=%+v\n", resp.Header)
	fmt.Printf("testDeleteEtcd.Deleted=%+v\n", resp.Deleted)
	fmt.Printf("testDeleteEtcd.PrevKvs=%+v\n", resp.PrevKvs)
}

//@注意：使用这个可以设定为一个主服务器。当服务器只能一个为主的时候。备用服务器通过这个可以通过这个知道，自己是备用服务器。
func testTxn(cli *clientv3.Client) {
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Printf("testTxn err=%v\n", err)
	//	}
	//}()

	//写
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	txn := cli.Txn(ctx)
	//如果glp1 > 100 设为99，否则设为101
	//如果没有这个值，会崩溃。。

	fmt.Printf(`"101">"99" = %v,%v\n`, "101" < "98", "99" > "98")
	//clientv3.Value 用于比较字符串。。,而且只能是字符串的比较结果。。  如果是数字的话就会出现上面的错误 101小于98
	resp, err := txn.If(clientv3.Compare(clientv3.Value("glp1"), ">", "98")).Then(
		clientv3.OpPut("glp1", "97")).Else(clientv3.OpPut("glp1", "99")).Commit()
	cancel()
	if err != nil {
		fmt.Printf("deleteEtcd err=%s\n", err.Error())
		return
	}
	fmt.Printf("testTxn.Header=%+v\n", resp.Header)
	fmt.Printf("testTxn.Succeeded=%+v\n", resp.Succeeded)
	fmt.Printf("testTxn.Responses=%+v\n", resp.Responses)
}

func testWatch(cli *clientv3.Client, key string) {
	//写
	watchChan := cli.Watch(context.Background(), key)
	for {
		select {
		case resp, ok := <-watchChan:
			if ok {
				fmt.Println("testWatch ok")
				fmt.Printf("testWatch resp=%+v\n", resp)

				for i := range resp.Events {
					fmt.Printf("testWatch resp.Events[%d]=%+v\n", i, resp.Events[i])
				}
			} else {
				fmt.Printf("testWatch not ok=%+v\n", resp)
			}
			WG.Done()
		}
	}
}

func testLease(cli *clientv3.Client) {
	resp, err := cli.Grant(context.Background(), 5) //获得一个租约 时间为5秒
	if err != nil {
		fmt.Printf("testLease err=%s\n", err.Error())
		return
	}

	fmt.Printf("testLease resp=%+v,id=%v,%v,%v\n", resp, strconv.FormatInt(int64(resp.ID), 16), resp.TTL, resp.Error)

	resp2, err := cli.Put(context.TODO(), "glp3", "1", clientv3.WithLease(resp.ID)) //用这个租约设置字段。
	if err != nil {
		fmt.Printf("testLease err=%s\n", err.Error())
		return
	}
	fmt.Printf("testLease put=%+v\n", resp2)
}

func testMutex(cli *clientv3.Client) {
	fmt.Printf("testMutex Begin\n")
	s, err := concurrency.NewSession(cli, concurrency.WithTTL(10))
	if err != nil {
		fmt.Printf("testMutex err=%s\n", err.Error())
		return
	}
	//err = concurrency.NewMutex(s, "lobby").Lock(context.TODO())//如果锁被占用，等到死。。。
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	err = concurrency.NewMutex(s, "lobby").Lock(ctx) //如果锁被占用，等到死。。。
	cancel()
	if err != nil {
		fmt.Printf("testMutex err=%s\n", err.Error())
		return
	}
	fmt.Printf("testMutex OK\n")
}

var WG sync.WaitGroup

func main() {
	/**
	@注意：编译报错：
	# github.com/coreos/etcd/clientv3/balancer/picker
	C:\GoPath\pkg\mod\github.com\coreos\etcd@v3.3.25+incompatible\clientv3\balancer\picker\err.go:37:44: undefined: balancer.PickOptions
	C:\GoPath\pkg\mod\github.com\coreos\etcd@v3.3.25+incompatible\clientv3\balancer\picker\roundrobin_balanced.go:55:54: undefined: balancer.PickOptions
	# github.com/coreos/etcd/clientv3/balancer/resolver/endpoint
	C:\GoPath\pkg\mod\github.com\coreos\etcd@v3.3.25+incompatible\clientv3\balancer\resolver\endpoint\endpoint.go:114:78: undefined: resolver.BuildOption
	C:\GoPath\pkg\mod\github.com\coreos\etcd@v3.3.25+incompatible\clientv3\balancer\resolver\endpoint\endpoint.go:182:31: undefined: resolver.ResolveNowOption

	改mod google.golang.org/grpc v1.33.1 // indirect 改成=> google.golang.org/grpc v1.26.0
	*/
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379", "127.0.0.1:22379", "127.0.0.1:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("err-61=%s\n", err.Error())
		return
	}
	defer cli.Close()

	//读
	testGetEtcd(cli, "glp")
	//增、改
	testPutEtcd(cli, "glp", "100")
	//查
	testGetEtcdWithPrefix(cli, "glp")
	//删
	testDeleteEtcd(cli, "glp2")
	//条件表达式。
	testTxn(cli)
	//监听
	WG.Add(1)
	go testWatch(cli, "glp1")

	//租约的使用
	testLease(cli)

	//分布式锁，当多个服务器要进行一个互斥的操作的时候。
	testMutex(cli)

	WG.Wait()
}

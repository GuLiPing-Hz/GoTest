package main

import (
	"fmt"
	"time"
	"sync/atomic"
	"sync"
	"strconv"
)

//学习 go chan信道(读取，写入，阻塞) goroutine多线程封装

//关键字 go 为我们开启一个goroutine

func do(start int) {
	var end = start + 5
	for i := start; i < end; i++ {
		fmt.Printf("%d ", i)
	}
}

//普通的执行测试
func test1() {
	do(0)
	do(5)

	/**
	输出
	0 1 2 3 4 5 6 7 8 9
	 */
}

//使用goroutine
func test2() {
	go do(0)
	do(5)
	/**
	非调试模式运行输出
	5 6 7 8 9 //偶尔可能都输出一点

	由于主线程结束的太快，我们的goroutine还没来得及跑，只do了一次
	 */
}

//创建无缓冲信道消息，主线程会等待消息的返回
func test3() {
	//定义一个信道，chan关键字表示信道，里面存储string，也可以存其他类型
	messages := make(chan string) //似乎全局信道不能这么写
	go func(message string) {
		time.Sleep(time.Second * 3)
		messages <- message // 存消息 使用 '<-', 指向信道表示存储消息
	}("hello chan message")

	fmt.Println("等待信道消息...")
	fmt.Println("收到来自信道的消息", <-messages) // 取消息 '<-', 反向信道表示读取消息
}

var ch chan int = make(chan int) //无缓冲信道
func do1() {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
	}
	ch <- 0 //存消息 通知主线程我们goroutine完成了
}

//测试无缓冲信道
func test4() {
	go do1()
	<-ch //这里可以阻塞主线程,如果我们不等待，do1将不会得到执行，也就不会有输出日志
}

var ch1 = make(chan int, 2) //2缓冲信道
var ch1Closed int32 = 0
var mutex = sync.Mutex{}
var wg = sync.WaitGroup{}

func produce() { //生产goroutine
	defer wg.Done()

	for i := 0; i < 100; i++ {
		mutex.Lock() //增加临界区的保护

		//增加原子锁也是没有用的。。我们必须对写入信道的时候实现临界区保护才行
		close := atomic.LoadInt32(&ch1Closed)

		if close == 1 { //
			//if ch1Closed == 1 { //单纯的这样写还是会panic
			fmt.Println("信道关闭，结束写入,再写入就会panic")
			mutex.Unlock()
			break
		}

		fmt.Println("send i=", i)
		ch1 <- i

		mutex.Unlock()
	}
}

func consume() { //消费goroutine
	//使用for循环读取太麻烦了，，，
	//for i := 0; i < 10; i++ {
	//	fmt.Println("read i = ", <-ch1)
	//	time.Sleep(time.Second)
	//}
	defer wg.Done()

	//go允许我们使用range一次读取信道中的多个数据
	for v := range ch1 {
		fmt.Println("read i = ", v)
		time.Sleep(time.Second)

		if ch1Closed == 1 && len(ch) <= 0 {
			// 如果当前信道已经关闭，等到信道中的数据量为0，跳出循环
			break
		}
	}
}

//测试缓冲信道
func test5() {
	//不用Sleep了，使用wait方法等待子线程结束
	wg.Add(2)
	go produce()
	go consume()
	time.Sleep(time.Second * 3)

	mutex.Lock()
	fmt.Println("关闭信道")
	// 显式地关闭信道
	//ch1Closed = 1//单纯的这样写还是会引发panic
	atomic.StoreInt32(&ch1Closed, 1)
	close(ch1) //关闭信道后,变成只读，如果还向里面写数据就会panic,
	mutex.Unlock()

	//time.Sleep(time.Second * 100)//老是用sleep也不是办法
	//go 提供c++类似 join等待子线程结束的方法

	fmt.Println("主线程wait ...")
	wg.Wait()
	fmt.Println("主线程wait 结束")
}

// 测试之前的流程控制select
// Go的select语句让程序线程在多个channel的操作上等待，
// select语句在goroutine 和channel结合的操作中发挥着关键的作用
func testSelect() {
	defer fmt.Println("testSelect end")
	//一般情况下，我们如果需要同时对多个信道进行读取操作，那么我们需要使用select

	var chanSelect1 = make(chan string)
	var chanSelect2 = make(chan int)
	var chanSelect3 = make(chan int32)
	var chanCnt int32 = 0

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(2 * time.Second)
			chanSelect1 <- "str_" + strconv.Itoa(i)
		}

		cnt := atomic.AddInt32(&chanCnt, 1)
		chanSelect3 <- cnt
	}()

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			chanSelect2 <- i * 2
		}

		cnt := atomic.AddInt32(&chanCnt, 1)
		chanSelect3 <- cnt
	}()

	fmt.Println("进入for循环 select")
	for { //select搭配for循环，可以起到一直运行的效果
		select { //如果只是select语句，那么其中一个case可以执行就执行下去了
		case v1 := <-chanSelect1:
			{
				fmt.Println("v1=", v1)
			}
		case v2 := <-chanSelect2:
			{
				fmt.Println("v2=", v2)
			}
		case v3 := <-chanSelect3:
			{
				fmt.Println("v3=", v3)
				if v3 == 2 {
					//break//只能跳出select
					goto end
				}
			}
		}
	}

end:
	fmt.Println("离开for循环 select")

}

func main() {
	//test1()
	//fmt.Println("\n" + strings.Repeat("*", 100))
	////test2()
	//fmt.Println("\n" + strings.Repeat("*", 100))
	//test3()
	//fmt.Println("\n" + strings.Repeat("*", 100))
	//test4()
	//fmt.Println("\n" + strings.Repeat("*", 100))
	//test5()
	//fmt.Println("\n" + strings.Repeat("*", 100))
	testSelect()
}

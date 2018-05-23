package main

import (
	"fmt"
	"time"
)

//学习 go goroutine

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

//创建无缓冲信道消息，使用sleep等待完成
func test3() {
	//定义一个信道，chan关键字表示信道，里面存储string，也可以存其他类型
	messages := make(chan string) //似乎全局通道不能这么写
	go func(message string) {
		messages <- message // 存消息 使用 '<-', 指向信道表示存储消息
	}("Ping!")

	fmt.Println(<-messages) // 取消息 '<-', 反向信道表示读取消息
	time.Sleep(time.Second) // 停顿一秒,等待goroutine执行完成
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

func produce() { //生产goroutine
	for i := 0; i < 10; i++ {
		fmt.Println("send i=", i)
		ch1 <- i
	}
}

func consume() { //消费goroutine
	//使用for循环读取太麻烦了，，，
	//for i := 0; i < 10; i++ {
	//	fmt.Println("read i = ", <-ch1)
	//	time.Sleep(time.Second)
	//}

	//go允许我们使用range读取信道
	for v := range ch1 {
		fmt.Println("read i = ", v)
		time.Sleep(time.Second)

		//if len(ch) <= 0 { // 如果现有数据量为0，跳出循环
		//	break //这个有点不友好
		//}
	}

	// 显式地关闭信道
	//close(ch)
}

//测试缓冲信道
func test5() {
	go produce()
	go consume()
	time.Sleep(time.Minute)
}

func main() {
	//test1()
	//test2()
	//test3()
	//test4()
	test5()
}

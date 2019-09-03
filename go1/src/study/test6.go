package main

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

//学习 go chan信道(读取，写入，阻塞) goroutine多线程封装

//关键字 go 为我们开启一个goroutine

/*
Go中channel可以是只读、只写、同时可读写的。
//定义只读的channel
read_only := make (<-chan int)
//定义只写的channel
write_only := make (chan<- int) //不带缓冲的信道。
//可同时读写
read_write := make (chan int, 10) //带10个缓冲的信道

注意使用信道的时候注意多个goroutine向同一个goroutine发送数据，如果没有足够的缓冲，
会导致goroutine泄漏的问题

对于使用有无缓冲，缓冲大小的问题：
无缓存channel更强地保证了每个发送操作与相应的同步接收操作；但 是对于带缓存channel，这些操作是解耦的
*/

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
		//close(messages)
		messages <- message // 存消息 使用 '<-', 指向信道表示存储消息
	}("hello chan message")

	fmt.Println("等待信道消息...")

	msg, ok := <-messages                        //第二个表示是否成功读取到消息,如果信道关闭，那么返回值就是false了
	fmt.Println("收到来自信道的消息=", msg, "；是否成功=", ok) // 取消息 '<-', 反向信道表示读取消息
}

func test3_1() {
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}

func test3_2() {
	//只写入chan
	chanOnlyW := make(chan bool)
	chanOnlyW <- true //如果chan没有缓冲，那么这个写操作需要等待有读操作的执行才会执行
	fmt.Printf("finish only w")
}

func test3_3() {
	//只读取chan
	chanOnlyR := make(chan bool)
	<-chanOnlyR //同理，如果chan没有缓冲，那么这个读操作需要等待有写操作的执行才会执行
	fmt.Printf("finish only r")
}

var ch chan int = make(chan int) //无缓冲信道
func do1() {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
	}
	ch <- 0 //存消息 通知主线程我们goroutine完成了
}

//理解只读，只写信道，编译期检查
//对于只写的信道，不能执行close
//这里并没有反向转换的语法：
//      就是不能一个将类似chan<-int类型的单向型的 channel转换为 chan int 类型的双向型的channel
func counter(out chan<- int) { //只写信道
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}
func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}
func printer(in <-chan int) { //只读信道
	for v := range in {
		fmt.Println(v)
	}
}

//测试无缓冲信道
func test4() {
	go do1()
	<-ch //这里可以阻塞主线程,如果我们不等待，do1将不会得到执行，也就不会有输出日志
}

//2缓冲信道,就是我们可以写入两个数据，写入第三个的时候会阻塞当前goroutine
var ch1 = make(chan int, 2)
var ch1Closed int32 = 0
var mutex = sync.Mutex{}
var wg = sync.WaitGroup{}

func produce() { //生产goroutine
	defer wg.Done()

	//len取到当前信道含有数据个数，cap获取当前信道的缓冲大小
	fmt.Printf("chan len=%d,cap=%d", len(ch1), cap(ch1))
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

func consume() {
	//消费goroutine
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

	//for range 信道 其实就是下面写法的简洁替代。
	//for {
	//	v, ok := <-ch1
	//	if !ok {
	//		break
	//	}
	//}
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

	go func() { //匿名函数
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

func testSelect1() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select { case x := <-ch:
			fmt.Println(x) // "0" "2" "4" "6" "8"
		case ch <- i:
		}
	}
}

func main() {
	//study()
	//fmt.Println("\n" + strings.Repeat("*", 100))
	////test2()
	//fmt.Println("\n" + strings.Repeat("*", 100))
	//test3()
	//test3_1()
	//test3_2()
	//test3_3()
	//fmt.Println("\n" + strings.Repeat("*", 100))
	//test4()
	//fmt.Println("\n" + strings.Repeat("*", 100))
	//test5()
	//fmt.Println("\n" + strings.Repeat("*", 100))
	//testSelect()
	testSelect1()

	//构造一个等待动画
	//go func() {
	//	for {
	//		for _, v := range `-\|/` {
	//			fmt.Printf("\r%c", v)
	//			time.Sleep(time.Millisecond * 100)
	//		}
	//	}
	//}()
	//time.Sleep(time.Second * 10)
}

package main

import (
	"fmt"
	"go1/src/pkglearn"
	"sync"
)

func testChanSync() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		fmt.Println("存100")
		pkglearn.Deposit(100)
		fmt.Println("存100,剩余", pkglearn.Balance())
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fmt.Println("取100")
		ok := pkglearn.Withdraw(100)
		fmt.Println("取100", ok, ",剩余", pkglearn.Balance())
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fmt.Println("取300")
		ok := pkglearn.Withdraw(300)
		fmt.Println("取300", ok, "剩余", pkglearn.Balance())
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fmt.Println("存120")
		pkglearn.Deposit(120)
		fmt.Println("存120,剩余", pkglearn.Balance())
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fmt.Println("存180")
		pkglearn.Deposit(180)
		fmt.Println("存180,剩余", pkglearn.Balance())
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("finish")
}

func testChanSync2() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		fmt.Println("存100")
		pkglearn.Deposit2(100)
		fmt.Println("存100,剩余", pkglearn.Balance2())
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fmt.Println("取100")
		ok := pkglearn.Withdraw2(100)
		fmt.Println("取100", ok, ",剩余", pkglearn.Balance2())
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fmt.Println("取300")
		ok := pkglearn.Withdraw2(300)
		fmt.Println("取300", ok, "剩余", pkglearn.Balance2())
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fmt.Println("存120")
		pkglearn.Deposit2(120)
		fmt.Println("存120,剩余", pkglearn.Balance2())
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fmt.Println("存180")
		pkglearn.Deposit2(180)
		fmt.Println("存180,剩余", pkglearn.Balance2())
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("finish")
}

var mutex sync.Mutex

func testMutexTwice() {
	mutex.Lock()
	defer mutex.Unlock()

	fmt.Println("第二次调用会卡住")
	//mutex.Lock() //引发panic
}

func testChanSync3() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		fmt.Println("存100")
		pkglearn.Deposit3(100)
		fmt.Println("存100,剩余", pkglearn.Balance3())
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fmt.Println("取100")
		ok := pkglearn.Withdraw3(100)
		fmt.Println("取100", ok, ",剩余", pkglearn.Balance3())
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fmt.Println("取300")
		ok := pkglearn.Withdraw3(300)
		fmt.Println("取300", ok, "剩余", pkglearn.Balance3())
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fmt.Println("存120")
		pkglearn.Deposit3(120)
		fmt.Println("存120,剩余", pkglearn.Balance3())
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fmt.Println("存180")
		pkglearn.Deposit3(180)
		fmt.Println("存180,剩余", pkglearn.Balance3())
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("finish")
}

var WG2 sync.WaitGroup

func testGoRace() {
	var x int32

	WG2.Add(3)
	go func() {
		defer WG2.Done()
		x += 1
		fmt.Println("1 x=", x)
	}()

	go func() {
		defer WG2.Done()
		x += 10
		fmt.Println("2 x=", x)
	}()

	go func() {
		defer WG2.Done()
		x += 100
		fmt.Println("3 x=", x)
	}()

	WG2.Wait()
	fmt.Println("finish trace")
}

func main() {
	//testChanSync()
	//testChanSync2()
	//testMutexTwice()
	//testChanSync3()

	/**
	对于查询余额加锁的解释：

	第一Balance不会在其它操作 比如Withdraw“中间”执行。
	第二(更重要)的是"同步"不仅仅是一堆goroutine执行顺序的问题； 同样也会涉及到内存的问题


	只要在go build，go run或者go test命令后面加上-race的flag，就会使编译器创建一个你的应
	用的“修改”版或者一个附带了能够记录所有运行期对共享变量访问工具的test，
	并且会记录下 每一个读或者写共享变量的goroutine的身份信息
	*/
	testGoRace()

	//go build -race xxx.go
	//运行编译出来的二进制文件就能看到数据竞争分析报告

	//var chanC chan byte
	//<-chanC //fatal error: all goroutines are asleep - deadlock!
}

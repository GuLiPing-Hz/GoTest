package main

import "fmt"

/**
defer 可以把表达式推入一个栈中，先进后出的方式执行
 */
func testDefer() {
	fmt.Println("testDefer begin")
	for i := 0; i < 5; i++ {
		defer fmt.Println(i) // 倒叙执行
	}

	{
		//实际执行效果告诉我们defer并没有闭包的概念，defer都是在函数返回之后执行的
		fmt.Println("testDefer scope begin")
		for i := 10; i < 15; i++ {
			defer fmt.Println(i) // 倒叙执行
		}
		fmt.Println("testDefer scope end")
	}

	fmt.Println("testDefer end")
}

func main() {
	fmt.Println("hello 1")

	for i := 100; i < 105; i++ {
		defer fmt.Println(i) // 倒叙执行
	}

	//函数中的defer会优先执行完毕，跟着函数走的defer
	testDefer()

	fmt.Println("hello 1 end")
}

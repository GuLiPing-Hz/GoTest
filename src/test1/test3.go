package main

import "fmt"

//学习 函数
//Go语言默认使用
// 值传递，修改参数不会影响到实参
// 引用传递，修改参数会影响到实参

func myPrintln(str string) {
	fmt.Println(str)
}

func add2(a int, b int) int {
	return a + b
}

//如果两个参数类型一致，可以像这样写成一个
func add(a, b int) (int) {
	return a + b
}

//函数返回多个值

func swap(a, b int) (int, int, int) {
	return b, a, a + b
}

//上面都是值传递

func sub(a, b *int) (int) {
	*a += 100
	return *a - *b
}

//前面test1和test2中都有main函数，一个go程序实例只允许存在一个main函数
func main() {
	myPrintln("myPrintln")
	fmt.Println("1+2=", add(1, 2))
	a, b := 3, 4
	fmt.Println(a, b)
	a, b, _ = swap(a, b) //go约定'_'是个只写变量，不可读
	fmt.Println(a, b)

	var c int = sub(&a, &b)
	fmt.Println(a, b, c)

	var f = add        //函数类型的变量
	fmt.Println(f(10, 11)) //调用

	//函数闭包
	fmt.Println("闭包例子\n", func(a, b int) int {
		return a*b + c //闭包里可以直接使用外面的参数
	}(4, 5))

	//定义结构体
	type Circle struct {
		radius float64
	}
	getArea := func(c Circle) float64 {//计算圆面积
		fmt.Println(c.radius)
		return 3.14 * c.radius * c.radius
	}
	var c1 Circle
	c1.radius = 10
	fmt.Println("area = ", getArea(c1))

}

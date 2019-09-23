package course

import "fmt"

//学习 函数
/*Go语言默认使用
实参通过值的方式传递，因此函数的形参是实参的拷贝。对形参进行修改不会影响实参。但是
如果实参包括引用类型，如指针，slice(切片)、map、function、channel等类型，实参可能会被函数内修改

@注意：数组传递是传递拷贝复制的形式，如不想拷贝直接传递指针类型


以下函数申明等价：
func f(i, j, k int, s, t string) {  }
func f(i int, j int, k int, s string, t string) {  }

以下4种方法声明拥有2个int型参数和1个int型返回值的函数：
func add(x int, y int) int {return x + y}
func sub(x, y int) (z int) { z = x - y; return} //这里给返回值起别名
func first(x int, _ int) int { return x }
func zero(int, int) int { return 0 }
*/
//func CGo(int, int) int //没有函数体，表示申明其他语言实现 c++

func myPrintln(str string) {
	fmt.Println(str)
}

func add2(a int, b int) (c int) {
	c = a + b
	return
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

//指定函数返回值的别名
func multi(a, b int) (product int) {
	product = a * b //这个用法可以在使用defer关键字的地方好用，defer后面再学习。
	return
}

/**
go 语言不支持函数的默认参数
func addEx(a=0, b=1 int) int {
	return a + b
}
*/

//可变长参数传递
func average(val ...int) int {
	result := 0
	for i, v := range val {
		fmt.Println(i, v)
		result += v
	}
	return result / len(val)
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

	var pro = multi(4, 6)
	fmt.Printf("multi 4*6 = %d\n", pro)

	fmt.Println("平均值1为", average(10, 5, 3, 4, 5, 6)) // 位置参数调用
	arra := []int{10, 5, 4}                          //数组或者切片传递给可变长参数函数调用
	fmt.Println("平均值2为", average(arra...))           // 位置参数调用

	var f = add            //函数类型的变量
	fmt.Println(f(10, 11)) //调用

	//函数闭包
	fmt.Println("闭包例子\n", func(a, b int) int {
		return a*b + c //闭包里可以直接使用外面的参数
	}(4, 5))

	//定义结构体
	type Circle struct {
		radius float64
	}
	getArea := func(c Circle) float64 { //计算圆面积
		fmt.Println(c.radius)
		return 3.14 * c.radius * c.radius
	}
	var c1 Circle
	c1.radius = 10
	fmt.Println("area = ", getArea(c1))
}

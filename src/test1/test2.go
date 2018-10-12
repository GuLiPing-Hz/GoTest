package main

import (
	"unsafe"
	"strings"
	"fmt"
)

/**
	学习 go语言类型，多重赋值，if/else/switch/select流程控制，for循环
	全局变量，局部变量

	go支持的类型：
		布尔类型 true/false
		数字类型
			整数：
				int8 uint8 int16 uint16 int32 uint32 int64 uint64
			浮点数：
				float32 float64
			复数：
				complex64  复数32位实数，
				complex128 复数64位实数
			其他数字：
				byte => uint8
				rune => int32
				uint int 根据机器32位或者64位
		字符串类型 utf-8
		派生类型
			(a) 指针类型（Pointer）
			(b) 数组类型
			(c) 结构化类型(struct)
			(d) Channel 类型
			(e) 函数类型
			(f) 切片类型
			(g) 接口类型（interface）
			(h) Map 类型
 */

//多维变量声明
var a, b, c = 1, 1.5, "123" //全局变量

func joke(arg int) { //arg是形参
	c := 2     //局部变量
	println(c) //未使用的局部变量编译会报错
}

//joke()//这里无法调用

func main() {
	joke(1)

	//查看大小长度
	println(a, b, c, len(c))
	// 字符串类型在 go 里是个结构, 包含指向底层数组的指针和长度,
	// 这两部分每部分都是 8 个字节，所以字符串类型大小为 16 个字节。
	println(unsafe.Sizeof(a), unsafe.Sizeof(b), unsafe.Sizeof(c))

	//多重赋值机制跟lua一样,变量可以下面这样交换值
	x := 1
	y := 2
	println("x=", x, ";y=", y)
	x, y = y, x
	println("x=", x, ";y=", y)

	//go 流程控制
	if true {
		println("True")
	}

	println("if else if")
	a = 1 //使用全局变量
	if a == 0 {
		println("a==0")
	} else if a == 1 {
		println("a==1")
	} else {
		println("!(a==0 || a==1)")
	}

	println("switch")
	switch a {
	case 0: //go中的case不需要写break，从上到下
		println("a==0")
	case 1:
		println("a==1")
	default:
		println("!(a==0 || a==1)")
	}

	var xInterface interface{} //nil
	xInterface = 10            //int 类型
	xInterface = 12.5          //float64
	xInterface = func(int) {}  //func(int) 类型
	xInterface = true

	println("sizeoof xInterface = ", unsafe.Sizeof(xInterface))
	switch xInterface.(type) {
	case nil:
		println("nil 类型")
	case int8:
		println("int16 类型")
	case int16:
		println("int16 类型")
	case int32:
		println("int32 类型")
	case int:
		println("int 类型")
	case float32:
		println("float32 类型")
	case float64:
		println("float64 类型")
	case func(int):
		println("func(int) 类型")
	case bool, string:
		println("bool, string 类型")
	default:
		println("未知类型")
	}

	//三目不行
	//b := a == 1 ? println("a==1"):println("a!=1")

	//go select语句
	//go 独特的select语句
	/**
	select {
    case communication clause  :
       statement(s);
    case communication clause  :
       statement(s);
    // 你可以定义任意数量的 case
	default : //可选
		statement(s);
	}

以下描述了 select 语句的语法：
	每个case都必须是一个通信
	所有channel表达式都会被求值
	所有被发送的表达式都会被求值
	如果任意某个通信可以进行，它就执行；其他被忽略。
	如果有多个case都可以运行，Select会随机公平地选出一个执行。其他不会执行。
	否则：
	如果有default子句，则执行该语句。
	如果没有default字句，select将阻塞，直到某个通信可以运行；Go不会重新对channel或值进行求值。

	//对于chan的教程，后面再说
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	  case i1 = <-c1:
		 fmt.Printf("received ", i1, " from c1\n")
	  case c2 <- i2:
		 fmt.Printf("sent ", i2, " to c2\n")
	  case i3, ok := (<-c3):  // same as: i3, ok := <-c3
		 if ok {
			fmt.Printf("received ", i3, " from c3\n")
		 } else {
			fmt.Printf("c3 is closed\n")
		 }
	  default:
		 fmt.Printf("no communication\n")
	}
	*/
	println(strings.Repeat("*", 50))

	//go 循环语句
	for j := 1; j <= 5; j++ { //c风格循环
		println("for 2 i=", j)
	}
	println(strings.Repeat("*", 50))
	i := 1
	for i <= 3 { // 类似while循环
		println("for 1 i=", i)
		i++
	}
	println(strings.Repeat("*", 50))
	for { // while(true),必须用break或者return跳出循环
		println("for 3 ")
		break //同样具有continue goto语句，用法同c
	}
	println(strings.Repeat("*", 50))
	//定义数组
	var numbers = [6]int{1, 2, 3, 4, 5, 6}
	for i, x := range numbers {
		//println("numbers["+strconv.Itoa(i)+"]=",x)
		//println("numbers[%d]=%d",i,x)//这个是不行的
		//fmt.Printf("numbers[%d]=%d;", i, x)
		println(fmt.Sprintf("numbers[%d]=%d;", i, x))
	}

	//支持循环嵌套for语句
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			println("for for", i, j)
		}
	}
}

package main

import (
	"fmt"
	"strings"
	"unsafe"
)

/**
学习 go语言类型，多重赋值，if/else/switch/select流程控制，for循环
全局变量，局部变量

@注意：捕获迭代变量的问题
需要注意for 循环中申明的循环体变量和局部变量赋值匿名函数的处理

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
	c := 2         //局部变量
	fmt.Println(c) //未使用的局部变量编译会报错
}

//joke()//这里无法调用

func main() {
	joke(1)

	//查看大小长度
	fmt.Println(a, b, c, len(c))
	// 字符串类型在 go 里是个结构, 包含指向底层数组的指针和长度,
	// 这两部分每部分都是 8 个字节，所以字符串类型大小为 16 个字节。
	fmt.Println(unsafe.Sizeof(a), unsafe.Sizeof(b), unsafe.Sizeof(c))

	//多重赋值机制跟lua一样,变量可以下面这样交换值
	x := 1
	y := 2
	fmt.Println("x=", x, ";y=", y)
	x, y = y, x
	fmt.Println("x=", x, ";y=", y)

	//go 流程控制
	if true {
		fmt.Println("True")
	}

	fmt.Println("if else if")
	a = 1 //使用全局变量
	if a == 0 {
		fmt.Println("a==0")
	} else
	if a == 1 {
		fmt.Println("a==1")
	} else
	{
		fmt.Println("!(a==0 || a==1)")
	}

	fmt.Println("switch")
	switch a {
	case 0: //go中的case不需要写break，从上到下
		fmt.Println("a==0")
	case 1:
		fmt.Println("a==1")
	default:
		fmt.Println("!(a==0 || a==1)")
	}

	var xInterface interface{} //nil
	xInterface = 10            //int 类型
	xInterface = 12.5          //float64
	xInterface = func(int) {}  //func(int) 类型
	xInterface = true

	xInterfaceBool, ok := xInterface.(bool) //第二个返回值，判断转换是否成功
	fmt.Printf("xInterfaceBool=%v,ok=%t\n", xInterfaceBool, ok)

	fmt.Println("sizeoof xInterface = ", unsafe.Sizeof(xInterface))
	switch xInterface.(type) { //不需要再写break,可以添加fallthrough来阻止默认的break
	case nil:
		fmt.Println("nil 类型")
	case int8:
		fmt.Println("int16 类型")
	case int16:
		fmt.Println("int16 类型")
	case int32:
		fmt.Println("int32 类型")
	case int:
		fmt.Println("int 类型")
	case float32:
		fmt.Println("float32 类型")
	case float64:
		fmt.Println("float64 类型")
	case func(int):
		fmt.Println("func(int) 类型")
	case bool, string:
		fmt.Println("bool, string 类型")
	default:
		fmt.Println("未知类型")
	}

	xInt := 11
	switch { //不带表达式，相当于 switch true,然后根据case里面比较，相当于 if else if
	case xInt > 10:
		fmt.Println("xInt > 10")
		fallthrough //表示当前标签不需要break
	case xInt > 5:
		fmt.Println("xInt > 5")
	default:
		fmt.Println("xInt = 5")
	case xInt < 5:
		fmt.Println("xInt < 5")
	}
	//三目不行
	//b := a == 1 ? fmt.Println("a==1"):fmt.Println("a!=1")

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
	fmt.Println(strings.Repeat("*", 50))

	//go 循环语句
	for j := 1; j <= 5; j++ { //c风格循环
		fmt.Println("for 2 i=", j)
	}
	fmt.Println(strings.Repeat("*", 50))
	i := 1
	for i <= 3 { // 类似while循环
		fmt.Println("for 1 i=", i)
		i++
	}
	fmt.Println(strings.Repeat("*", 50))
	for { // while(true),必须用break或者return跳出循环
		fmt.Println("for 3 ")
		break //同样具有continue goto语句，用法同c
	}
	fmt.Println(strings.Repeat("*", 50))
	//定义数组
	var numbers = [6]int{1, 2, 3, 4, 5, 6}
	for i, x := range numbers {
		//fmt.Println("numbers["+strconv.Itoa(i)+"]=",x)
		//fmt.Println("numbers[%d]=%d",i,x)//这个是不行的
		//fmt.Printf("numbers[%d]=%d;", i, x)
		fmt.Println(fmt.Sprintf("numbers[%d]=%d;", i, x))
	}

	//支持循环嵌套for语句
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println("for for", i, j)
		}
	}
}

package main

import (
	"fmt"
	"strconv"
	"runtime"
	"path"
	"strings"
	"reflect"
)

/*

初学Go

Go的安装：https://studygolang.com/dl

window用户可以使用msi安装，并且必须在环境变量中指定GOPATH,我把GoPath放在C盘，跟GoRoot位置一致

注释语法类似c //单行，多行就像我现在写的提示文字

 */

//import "fmt"最常用的一种形式
//import "./test"导入同一目录下test包中的内容
//import f "fmt"导入fmt，并给他启别名ｆ
//import . "fmt"，将fmt启用别名"."，这样就可以直接使用其内容，而不用再添加ｆｍｔ，如fmt.Println可以直接写成Println
//import  _ "fmt" 表示不使用该包，而是只是使用该包的init函数，并不显示的使用该包的其他内容。注意：这种形式的import，当import时就执行了fmt包中的init函数，而不能够使用该包的其他函数。

func main() {
	//最简单的打印-注释(单行)
	fmt.Println("Hello 世界") //两个语句写同一行时才需要分号

	//查看变量地址
	//查看变量类型
	var vGlp int = 1
	fmt.Printf("vGlp address = %p, type = %s\n", &vGlp, reflect.TypeOf(vGlp))

	//声明常量
	const ConstA = 1
	const ( //常量当枚举使用
		Unknown = 0
		Female  = 1
		Male    = 2
	)
	// iota 特殊常量，可以认为是一个可以被编译器修改的常量
	// 在没一个const关键字出现是，被重置为0，
	// 然后再下一个const出现之前，每出现一次iota，其所代表的数字会自动增加1。
	const (
		cA = iota
		cB = iota
		cC = iota
	)
	print("cA,cB,cC=")
	println(cA, cB, cC)
	//上面可以简写为
	const (
		cA1 = iota
		cB1
		cC1
	)
	print("cA1,cB1,cC1=")
	println(cA1, cB1, cC1)

	const (
		a2 = iota //0
		b2        //1
		c2        //2
		d2 = "ha" //独立值，iota += 1
		e2        //"ha"   iota += 1
		f2 = 100  //iota +=1
		g2        //100  iota +=1
		h2 = iota //7,恢复计数
		i2        //8
	)
	fmt.Println(a2, b2, c2, d2, e2, f2, g2, h2, i2)

	//声明变量
	var a = 10
	b := 20 //省略var

	//声明一般类型
	type newInt int

	//1.算术运算符:
	println("a+b=", a+b) //加
	println("a-b=", a-b) //减
	println("a*b=", a*b) //乘
	println("b/a=", b/a) //除
	println("b%a=", b%a) //模
	println("-a=", -a)

	//2.比较操作符:
	println("a==b", a == b)
	println("a!=b", a != b)
	println("a>b", a > b)
	println("a<b", a < b)
	println("a>=b", a >= b)
	println("a<=b", a <= b)

	println("a =", a)
	a ++
	//a11 := a ++ //语法错误
	println("a ++=", a)
	a --
	//a12 := a -- //语法错误
	println("a --=", a)

	//3.赋值运算符：
	var c = 1
	c += a
	println("c+=a =", c)
	c -= a
	println("c-=a =", c)

	c *= 2
	println(c)
	c /= 2
	println(c)
	c %= 1
	println(c)

	// 4.位运算符：
	a = 3  //二进制的表示  0000 0011
	b = 10 //二进制的表示  0000 1010

	println("a&b =", a&b)   //按位与
	println("a|b =", a|b)   //按位或
	println("a^b =", a^b)   //按位异或
	println("~b =", ^b)     //按位取反
	println("a<<2 =", a<<2) //按位左移
	println("b>>2 =", b>>2) //按位右移

	// 5.逻辑运算符:
	a1 := true // 这里必须大写
	b1 := false
	println(a1 && a1)
	println(a1 || b1)
	println(! a1)

	//6.其他运算符: &取地址，*取内容 类似c++
	var aPointer *bool = &a1
	println("&a1=", aPointer)        //返回变量存储地址
	println("*aPointer=", *aPointer) //读取地址中的内容

	//字符串连接
	println("123" + "abc")

	//字符串数字转换
	//字符串转int
	num, err := strconv.Atoi("123")
	//字符串转int64
	var num64 int64
	num64, err = strconv.ParseInt("1234", 10, 64)
	if err != nil {
		println(err)
	}
	println("字符串->数字:", num, num64)
	str := strconv.Itoa(456)
	println("数字->字符串:", str)

	//读取当前文件名
	_, file, line, ok := runtime.Caller(0) //使用下划线告诉编译器抛弃返回值
	println(file, line, ok)
	base := path.Base(file)
	ext := path.Ext(file)
	println(base, ext, strings.TrimSuffix(base, ext))

	//println对小数打印有问题
	println(3.14)
	fmt.Println(3.14)

	var a3 int64 = 100
	var b3 int = 10
	//fmt.Println(a3 + b3)//编译不过
	fmt.Println(a3 + int64(b3)) //必须转成相同类型的数据才能操作
}

package main

import "fmt"

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

	//声明常量
	const ConstA = 1

	//声明变量
	var a = 10
	b := 20

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

	// 6.成员运算符:
	a2 := "a"
	b2 := "abcdefg"
	println("a in b =",a2 in b2) //判断a是否在b的里面，可以是字符串，或者是元组，序列，字典
	println("a not in b =",a2 not in b2)
}

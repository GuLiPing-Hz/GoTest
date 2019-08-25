/*
包名，一个应用程序只有一个main包，main包里需提供main函数
名字的长度没有逻辑限制，但是Go语言的风格是尽量使用短小的名字，对于局部变量尤其是
这样；你会经常看到i之类的短名字，而不是冗长的theLoopIndex命名。
通常来说，如果一个 名字的作用域比较大，生命周期也比较长，那么用长的名字将会更有意义
*/
package main

import (
	"fmt"
	"math"
	"path"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
)

/*

初学Go
常量，变量，运算符，读取脚本当前文件名

Go的安装：https://studygolang.com/dl

window用户可以使用msi安装，并且必须在环境变量中指定GOPATH,我把GoPath放在C盘，跟GoRoot位置一致

注释语法类似c //单行，多行就像我现在写的提示文字

*/

//import "fmt"最常用的一种形式
//import "./test"导入同一目录下test包中的内容
//import f "fmt"导入fmt，并给他启别名ｆ
//import . "fmt"，将fmt启用别名"."，这样就可以直接使用其内容，而不用再添加ｆｍｔ，如fmt.Println可以直接写成Println
//import  _ "fmt" 表示不使用该包，而是只是使用该包的init函数，并不显示的使用该包的其他内容。注意：这种形式的import，当import时就执行了fmt包中的init函数，而不能够使用该包的其他函数。

/**
%d 十进制整数
%x, %o, %b 十六进制，八进制，二进制整数。
%f, %g, %e 浮点数： 3.141593 3.141592653589793 3.141593e+00
%t 布尔：true或false
%c 字符（rune） (Unicode码点)
%s 字符串
%q 带双引号的字符串"abc"或带单引号的字符'c'
%v 变量的自然形式（natural format）
%T 变量的类型
%% 字面上的百分号标志（无操作数）

请注意fmt的两个使用技巧。
	通常Printf格式化字符串包含多个%参数时将会包含对应相同数量 的额外操作数，但是%之后的 [1] 副词告诉Printf函数再次使用第一个操作数。
	第二，%后 的 # 副词告诉Printf在用%o、%x或%X输出时生成0、0x或0X前缀,#号也只能搭配o,x,X使用

	小技巧控制输出的缩进。 %*s 中的 * 会在字符串之前填充一 些空格。
	fmt.Printf("%*s",2,"") 每次输出会先填充 2 数量的空格，再输出""
*/

//包一级变量声明 var 变量名字 类型 = 表达式
var (
	//包一级的各种类型的声明语句的顺序 无关紧要（译注：函数内部的名字则必须先声明之后才能使用）
	OutA = 1 //首字母大写可以让外部包访问该变量，函数等
	inA  = 2 //首字母小写表示只能在本包内的文件访问。
)

/**
一个类型声明语句创建了一个新的类型名称，和现有类型具有相同的底层结构。
新命名的类 型提供了一个方法，用来分隔不同概念的类型，这样即使它们底层类型相同也是不兼容的

类型声明语句一般出现在包一级，因此如果新创建的类型名字的首字符大写，则在外部包也 可以使用
*/
//定义新类型 关键字type 类型名字 底层类型
type newInt int //底层类型虽然一样，但是go是严格的类型语言，不同类型加减的时候需要强制转换。
func (imp newInt) String() string { //类型都会定义默认的String方法用于某个类型的输入格式
	return fmt.Sprintf("newInt:%d", int(imp))
}

//包一级函数声明
func main() {
	//最简单的打印-注释(单行)
	fmt.Println("Hello 世界") //两个语句写同一行时才需要分号
	fmt.Println(`Hello 
世界2`)

	//查看变量地址
	//查看变量类型
	var vGlp int = 1
	fmt.Printf("vGlp address = %p, type = %s\n", &vGlp, reflect.TypeOf(vGlp))

	charA := 'a'
	charB := '事'
	charC := '\n'
	fmt.Printf("charA = %d %[1]c %[1]q\n", charA)
	fmt.Printf("charB = %d %[1]c %[1]q %c\n", charB, 20107) //第一个打印的unicode码
	fmt.Printf("charC = %d %[1]c %[1]q\n", charC)

	//声明常量 - 常量精度256bit，只有常量可以是无类型的
	const ConstA = 1
	const ( //常量当枚举使用
		Unknown = 0
		Female  = 1
		Male    = 2
	)
	{
		//无类型常量自动转换成默认的类型
		i := 0      // untyped integer; implicit int(0)
		r := '\000' // untyped rune; implicit rune('\000')
		f := 0.0    // untyped floating-point; implicit float64(0.0)
		c := 0i     // untyped complex; implicit complex128(0i)
		fmt.Println(i, r, f, c)
	}

	// iota 特殊常量，可以认为是一个可以被编译器修改的常量
	// 在没一个const关键字出现是，被重置为0，
	// 然后再下一个const出现之前，每出现一次iota，其所代表的数字会自动增加1。
	const (
		cA = iota
		cB = iota
		cC = iota
	)
	fmt.Print("cA,cB,cC=")
	fmt.Println(cA, cB, cC)
	//上面可以简写为
	const (
		cA1 = iota
		cB1
		cC1
	)
	fmt.Print("cA1,cB1,cC1=")
	fmt.Println(cA1, cB1, cC1)

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
	fmt.Println("IOTA", a2, b2, c2, d2, e2, f2, g2, h2, i2)

	const (
		aa1 = g2 + iota
		aa2
		aa3
		aa4
	)
	fmt.Println("IOTA", aa1, aa2, aa3, aa4)

	//声明变量
	//1.声明一般类型
	var a = 10 //go对创建的值都赋予0值，字符串是空字符串，指针,函数是nil

	/**
	2.简短变量声明左边的变量可能并不是全部都是刚刚声明的。如 果有一些已经在相同的词法域声明过了（§2.7），
	那么简短变量声明语句对这些已经声明过 的变量就只有赋值行为了。
	*/
	b := 20 //省略var 简短变量声明

	/*
		指针 跟C语言类似，但是Go语言中的指针不支持加减操作，即只能对指针指向的变量加减运算。
		在Go语言中，返回函数中局部变量的地址也是安全的
		*p++ // 非常重要：只是增加p指向的变量的值，并不改变p指针！！！
	*/
	p := &a //申明一个指针，指向a的内存地址
	r := (*int)(p)
	fmt.Printf("p type = %T,%v,%03d,r type=%T\n", p, *p, a, r)
	//p++ 编译报错，没有该运算
	*p++
	fmt.Printf("*p++ type = %T,%v,%03d\n", p, *p, a)

	//创建指针的方法，使用new内建函数，跟C++差不多
	//表达式new(T)将创建一个T类型的匿名变 量，初始化为T类型的零值，然后返回变量地址，返回的指针类型为 *T
	q := new(int64)
	fmt.Printf("*q type = %T,%v,%03x\n", q, *q, a)

	var d newInt
	fmt.Printf("test newInt = %v\n", d)

	var z float64                      //float32类型的累计计算误差很容易扩散，并且float32能精确表示的正整数并不是很大，只有23位表示数字
	fmt.Println(z, -z, 1/z, -1/z, z/z) //0 -0 +Inf -Inf NaN
	fmt.Println(math.IsNaN(z/z), math.NaN(), 1/z == 1/z, -1/z == -1/z, z/z == z/z)

	//1.算术运算符:
	a = 10
	b = 8
	fmt.Println("a+b=", a+b)     //加
	fmt.Println("a-b=", a-b)     //减
	fmt.Println("a*b=", a*b)     //乘
	fmt.Println("b/a=", b/a)     //除
	fmt.Println("b%a=", b%a)     //模 模的符号依赖于被取模的数的符号
	fmt.Println("b%-a=", b % -a) //模 模的符号依赖于被取模的数的符号
	fmt.Println("-b%a=", -b%a)   //模 模的符号依赖于被取模的数的符号
	fmt.Println("-a=", -a)

	//2.比较操作符:
	fmt.Println("a==b", a == b)
	fmt.Println("a!=b", a != b)
	fmt.Println("a>b", a > b)
	fmt.Println("a<b", a < b)
	fmt.Println("a>=b", a >= b)
	fmt.Println("a<=b", a <= b)

	fmt.Println("a =", a)
	a++
	//a11 := a ++ //语法错误
	fmt.Println("a ++=", a)
	a--
	//a12 := a -- //语法错误
	fmt.Println("a --=", a)

	//3.赋值运算符：
	var c = 1
	c += a
	fmt.Println("c+=a =", c)
	c -= a
	fmt.Println("c-=a =", c)

	c *= 2
	fmt.Println(c)
	c /= 2
	fmt.Println(c)
	c %= 1
	fmt.Println(c)

	// 4.位运算符：
	a = 3                                                                     //二进制的表示  0000 0011
	b = 10                                                                    //二进制的表示  0000 1010
	fmt.Printf("a=%d(二进制%04[1]b),b=%d(二进制%04[2]b)\n", a, b)                   //按位与
	fmt.Printf("a&b = %d(二进制%04[1]b,八进制%#[1]o,十六进制%#[1]x,十六进制%#[1]X)\n", a&b) //按位与
	fmt.Printf("a|b =%d(二进制%04[1]b)\n", a|b)                                  //按位或
	fmt.Printf("a^b =%d(二进制%04[1]b)\n", a^b)                                  //按位异或
	fmt.Printf("^b =%d(二进制%04[1]b)\n", ^b)                                    //按位取反
	fmt.Printf("a&^b =%d(二进制%04[1]b)\n", a&^b)                                //按位清空
	fmt.Printf("3&^13 =%d(二进制%04[1]b)\n", 3&^13)                              //按位清空 其实就是先对后面的数按位反再跟前面的数按位与
	fmt.Printf("a<<2 =%d(二进制%04[1]b)\n", a<<2)                                //按位左移
	fmt.Printf("b>>2 =%d(二进制%04[1]b)\n", b>>2)                                //按位右移

	// 5.逻辑运算符: 短路行为；助记： && 对应逻辑乘法， || 对应逻辑加法，乘法比加法优先 级要高
	a1 := true // 这里必须大写
	b1 := false
	fmt.Println(a1 && a1)
	fmt.Println(a1 || b1)
	fmt.Println(! a1)

	//6.其他运算符: &取地址，*取内容 类似c++
	var aPointer *bool = &a1
	fmt.Println("&a1=", aPointer)        //返回变量存储地址
	fmt.Println("*aPointer=", *aPointer) //读取地址中的内容

	//字符串连接
	fmt.Println("123" + "abc")

	//字符串数字转换
	//字符串转int
	num, err := strconv.Atoi("123")
	//字符串转int64
	var num64 int64
	num64, err = strconv.ParseInt("1234", 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("字符串->数字:", num, num64)
	str := strconv.Itoa(456)
	tt := time.Now().Unix()
	fmt.Println("数字->字符串:", str, tt, strconv.FormatInt(tt, 10))
	bb := []byte(str)
	fmt.Println("字符串->[]byte", bb)
	fmt.Println("[]byte->字符串", string(bb))

	//读取当前文件名
	_, file, line, ok := runtime.Caller(0) //使用下划线告诉编译器抛弃返回值
	fmt.Println(file, line, ok)
	dir := path.Dir(file)
	base := path.Base(file)
	ext := path.Ext(file)
	fmt.Println(dir, base, ext, strings.TrimSuffix(base, ext))

	//更多其他基础测试
	//println对小数打印有问题
	fmt.Println(3.14)
	fmt.Println(3.14)

	var a3 int64 = 100
	var b3 int = 10
	//fmt.Println(a3 + b3)//编译不过
	fmt.Println(a3 + int64(b3)) //必须转成相同类型的数据才能操作

	var s1 string
	fmt.Println(s1 == "") //没有 s1 == nil

	a4, a5, a6, a7 := 0, 0, 0, 0
	fmt.Println("a4&a5&a6&a7 = ", a4&a5&a6&a7)

	type User struct {
		name  string
		score int
	}
	var users []User
	users = append(users, User{"a", 8})
	users = append(users, User{"b", 10})
	for _, user := range users {
		fmt.Println("name=", user.name, "score=", user.score)
	}

	fmt.Println("切片实验，修改前")
	var bytes [] byte
	bytes = append(bytes, 1, 2, 3, 4, 5, 6, 7, 8)
	for _, b := range bytes {
		fmt.Print(b, " ")
	}
	fmt.Println()
	pos := 2
	bytes1 := bytes[pos : pos+2] //获取其中的部分切片
	bytes1[0] = 30               //修改其中的内容
	bytes1[1] = 40
	fmt.Println("切片实验，修改后")
	for _, b := range bytes {
		fmt.Print(b, " ")
	}
	fmt.Println()

	//没有三目运算符
}

package main

import "unsafe"

/**
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

func joke() {
	c := 2     //局部变量
	println(c) //未使用的局部变量编译会报错
}

//joke()//这里无法调用

func main() {
	joke()

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

	//流程控制
	if true {
		println("True")
	}

	a = 1 //使用全局变量
	if a == 0 {
		println("a==0")
	} else if a == 1 {
		println("a==1")
	} else {
		println("!(a==0&&a==1)")
	}

	//a == 1 ? println("a==1"):println("a!=1")
}

/**
公斤(千克)/斤相互转换
 */
package pkglearn

import "fmt"

const (
	kg2J = 2
)

type Kg float32
type Jin float32

func (imp Kg) String() string {
	return fmt.Sprintf("%g(千克/公斤)", imp)
}

func (imp Jin) String() string {
	return fmt.Sprintf("%g(斤)", imp)
}

func Kg2Jin(v Kg) Jin {
	return Jin(v * 2)
}

func Jin2Kg(v Jin) Kg {
	return Kg(v / 2)
}

//包的初始化顺序
var a = b + c // a 第三个初始化, 为 3
var b = f()   // b 第二个初始化, 为 2, 通过调用 f (依赖c)
var c = 1     // c 第一个初始化, 为 1
func f() int { return c + 1 }

//练习题2.3
var pc [256]byte

/*
对于在包级别声明的变量，如果有初始化表达式则用表达式初始化，还有一些没有初始化表 达式的，
例如某些表格数据初始化并不是一个简单的赋值过程。在这种情况下，
我们可以用 一个特殊的init初始化函数来简化初始化工作。每个文件都可以包含多个init初始化函数
*/
func init() {
	/*
	@注意：这样的init初始化函数除了不能被调用或引用外，其他行为和普通函数类似。
	在每个文件中的 init初始化函数，在程序开始执行时按照它们声明的顺序被自动调用。
	*/
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] + pc[byte(x>>(1*8))] + pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] + pc[byte(x>>(4*8))] + pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] + pc[byte(x>>(7*8))])
}

func PopCount2_3(x uint64) int {
	var result int
	for i := uint(0); i < 8; i++ {
		temp := int(pc[byte(x>>(i*8))])
		result += int(temp)
	}
	return result
}

func PopCount2_4(x uint64) int {
	var result int
	for i := uint(0); i < 64; i++ {
		result += int((x >> i) & 0x01)
	}
	return result
}

func PopCount2_5(x uint64) int {
	var result int
	for x != 0 {
		x = x & (x - 1)
		result++
	}
	return result
}

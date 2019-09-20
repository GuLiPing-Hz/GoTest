package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/**
各个底层类型的对齐字节数

bool 1个字节
intN, uintN, floatN, complexN N/8个字节(例如float64是8个字节)
//上面的是按字节，下面按机器字(32位就对应4字节，64位对应8字节)

int, uint, uintptr 1个机器字
*T 1个机器字
string 2个机器字(data,len)
[]T 3个机器字(data,len,cap)
map 1个机器字
func 1个机器字
chan 1个机器字
interface 2个机器字(type,value)
*/
type X struct {
	//64位机器
	a bool    //1字节
	b float64 //对齐空7字节+8字节
	c int16   //2字节 +补齐6字节
}

type X2 struct {
	a bool    //1字节
	c int16   //对齐空1字节 2字节
	b float64 //对齐空4字节 + 8字节
}

func isIntX(v reflect.Value) bool {
	return v.Kind() == reflect.Int || v.Kind() == reflect.Int8 ||
		v.Kind() == reflect.Int16 || v.Kind() == reflect.Int32 ||
		v.Kind() == reflect.Int64
}

func isUintX(v reflect.Value) bool {
	return v.Kind() == reflect.Uint || v.Kind() == reflect.Uint8 ||
		v.Kind() == reflect.Uint16 || v.Kind() == reflect.Uint32 ||
		v.Kind() == reflect.Uint64
}

func lx13_1(x, y interface{}) bool {
	v1 := reflect.ValueOf(x)
	v2 := reflect.ValueOf(y)

	if !isIntX(v1) || !isIntX(v2) {
		return false
	}
	return v1.Int() == v2.Int()
}

func lx13_2_inner(v reflect.Value, dict map[uintptr]bool) bool {
	if v.CanAddr() {
		//uptr := uintptr(unsafe.Pointer(&v))
		uptr := v.UnsafeAddr()
		_, ok := dict[uptr]
		if ok {
			return true
		}
		dict[uptr] = true
	}

	fmt.Println("kind=", v.Kind())
	if v.Kind() == reflect.Struct {
		for i := 0; i < v.NumField(); i++ {
			fieldV := v.Field(i)
			if lx13_2_inner(fieldV, dict) {
				return true
			}
		}
		return false
	} else if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		return lx13_2_inner(v.Elem(), dict)
	} else {
		return false
	}
}

func lx13_2(x interface{}) bool {
	v := reflect.ValueOf(x)
	dict := make(map[uintptr]bool)
	return lx13_2_inner(v, dict)
}

//数据相同的两个结构，X2占用的内存更小。
func main() {
	x := X{}
	x2 := X2{}
	//打印实际占用字节大小
	fmt.Printf("Sizeof(x)=%d\n", unsafe.Sizeof(x))
	fmt.Printf("Sizeof(x2)=%d\n", unsafe.Sizeof(x2))
	fmt.Printf("Sizeof(x.a)=%d\n", unsafe.Sizeof(x.a))
	fmt.Printf("Sizeof(x.b)=%d\n", unsafe.Sizeof(x.b))
	fmt.Printf("Sizeof(x.c)=%d\n", unsafe.Sizeof(x.c))

	//打印对齐大小
	fmt.Printf("Alignof(x.a)=%d\n", unsafe.Alignof(x.a))
	fmt.Printf("Alignof(x.b)=%d\n", unsafe.Alignof(x.b))
	fmt.Printf("Alignof(x.c)=%d\n", unsafe.Alignof(x.c))

	//打印字段在结构中的偏移位置
	fmt.Printf("Offsetof(x.a)=%d\n", unsafe.Offsetof(x.a))
	fmt.Printf("Offsetof(x.b)=%d\n", unsafe.Offsetof(x.b))
	fmt.Printf("Offsetof(x.c)=%d\n", unsafe.Offsetof(x.c))
	fmt.Printf("Offsetof(x2.a)=%d\n", unsafe.Offsetof(x2.a))
	fmt.Printf("Offsetof(x2.b)=%d\n", unsafe.Offsetof(x2.b))
	fmt.Printf("Offsetof(x2.c)=%d\n", unsafe.Offsetof(x2.c))

	fmt.Println()
	var a = float64(1)

	//通过对unsafe.Pointer的桥接，我们可以将任意类型指针转换成其他类型的指针
	//unsafe.Pointer可以理解成c语言里面的void*指针
	var b = (*int64)(unsafe.Pointer(&a))
	fmt.Printf("a=%b,b=%016x\n", a, *b)

	fmt.Println("练习题13.1", reflect.DeepEqual(int16(5), int32(5)),
		lx13_1(int16(5), int32(5)))

	type Circle2 struct {
		val  int
		next *Circle2
	}

	circle := &Circle2{val: 10}
	circle.next = circle
	fmt.Println("练习题13.2 是否有环？", lx13_2(circle), lx13_2(*circle), lx13_2(a))
}

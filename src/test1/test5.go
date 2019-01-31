=package main

import (
	"fmt"
	"reflect"
	"math"
	"test1/codeerror"
	"errors"
)

//学习 go语言接口，实现，异常处理 panic recover
/**
定义接口
type interface_name interface {
	method_name1() [return_type]
method_name2() [return_type]
method_name3() [return_type]
...
method_namen() [return_type]
}

定义结构体
type struct_name struct {
	// variables
}

 实现接口方法
func (struct_name_variable struct_name) method_name1() [return_type] {
	// 方法实现
}
...
func (struct_name_variable struct_name) method_namen() [return_type] {
	//方法实现
}
*/
type Animal interface {
	getName() string

	//isKind() bool
	//
	//add(a, b int) int
	//
	//feed(food string)
}

type Person interface {
	getName() string
}

type Cat struct {
}

func (animal Cat) getName() string {
	return "Cat"
}

type Dog struct {
}

func (animal Dog) getName() string {
	return "Dog"
}

type Man struct {
	name string
}

type ManBePerson Man //同一个类型，但是可以作为不同的特性

func (a *Man) getName() string {
	return "Man"
}

//这里实现
func (a *ManBePerson) getName() string {
	return a.name
}

type ManBeAnimal Man

func (a *ManBeAnimal) getName() string {
	return "Human"
}

func testInterface() {
	//结构体可以通过new来创建，gc回收
	var animal1 Animal
	animal1 = new(Cat) //必须实现所有接口才能new
	fmt.Println("Animal name is", animal1.getName())

	var animal2 Animal
	animal2 = new(Dog)
	fmt.Println("Animal name is", animal2.getName())

	//结构体也能通过下面的形式创建
	fmt.Println("比较两个类型 ManBePerson == Man ", reflect.TypeOf(ManBePerson{}) == reflect.TypeOf(Man{}))
	var man = Man{"Jack"}
	var manP1 *ManBePerson = (*ManBePerson)(&man) //定义不同的特性
	var manP2 *ManBeAnimal = (*ManBeAnimal)(&man)

	//针对同一个对象，我们可以通过type定义不同的类型(本质是同一个类型)，实现不同的特性
	fmt.Println(man.getName(), manP1.getName(), manP2.getName())
}

func Sqrt(f float64) (float64, error) {
	result := math.Sqrt(f)
	if f < 0 {
		//提示原生错误
		//return math.Sqrt(f), errors.New("math: square root of negative number")
		return result, codeerror.New(1, "math: square root of negative number")
	}
	// 实现
	return result, nil
}

func testError() {
	fmt.Println("testError in")

	defer func() {
		fmt.Println("defer func in")
		p := recover() //recover用于还原现场程序，可继续执行，与panic配对使用
		if p != nil {
			if reflect.TypeOf(p) == reflect.TypeOf(&codeerror.CodeError{}) {
				//interface转换成指定类型的数据,类型不能弄错，否则生成的时候会报错
				var er = p.(*codeerror.CodeError)
				fmt.Println(er)
				fmt.Println("er.Reserve = ", er.Reserve)
			}
			fmt.Println(reflect.TypeOf(p), reflect.TypeOf(&codeerror.CodeError{}))
			fmt.Printf("Fatal error: %s\n", p)
		}

		fmt.Println("defer func out")
	}()

	//错误
	a, err := Sqrt(-1)
	fmt.Println("a=", a) //
	if err != nil {
		//一般都会传入error类型变量，panic用于终止程序继续执行，但是会运行之前保存的defer
		panic(err)

		/**
		对于panic，一般是不得不用的情况下才使用，否则建议不要使用过多的异常机制，
		小错误还是走小错误的方式，不要乱抛异常
		 */
	}

	//由于上面的panic 导致我们这里的代码无法继续执行
	fmt.Println("testError out")
}

/**
伪代码
//检查错误的方法1
func checkError(err error) {
    if err != nil {
        fmt.Println("Error is ", err)
        os.Exit(-1)
    }
}

func foo() {
    err := doStuff1()
    checkError(err)

    err = doStuff2()
    checkError(err)

    err = doStuff3()
    checkError(err)
}
 */

//检查错误的方法2
type Something struct {
	err   error
	index int
}

func (thing *Something) do() (int, error) {
	if thing.err != nil {
		return thing.index, thing.err
	}

	//do something
	thing.index++
	//操作thing.err

	return thing.index, thing.err
}

//还可以通过这样的方式定义错误
var (
	ErrInvalid    = errors.New("invalid argument")
	ErrPermission = errors.New("permission denied")
	ErrExist      = errors.New("file already exists")
	ErrNotExist   = errors.New("file does not exist")
)
/*
伪代码，处理不同的错误方式
err := os.XXX
if err == os.ErrInvalid {
    //handle invalid
}
 */

func main() {
	fmt.Println("main in")

	testInterface()
	testError()

	fmt.Println("main out")
}

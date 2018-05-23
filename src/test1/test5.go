package main

import (
	"fmt"
	"reflect"
	"math"
	"test1/codeerror"
)

//go语言接口
/**
定义接口
type interface_name interface {
	method_name1 [return_type]
method_name2 [return_type]
method_name3 [return_type]
...
method_namen [return_type]
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
		if p := recover(); p != nil {
			if reflect.TypeOf(p) == reflect.TypeOf(&codeerror.CodeError{}) {
				var er = p.(codeerror.CodeError) //interface转换成指定类型的数据
				fmt.Println(er.Reserve)
			}
			fmt.Println("reflect.TypeOf(p) == reflect.TypeOf(codeerror.CodeError{})",
				reflect.TypeOf(p) == reflect.TypeOf(codeerror.CodeError{}))
			fmt.Println(reflect.TypeOf(p), reflect.TypeOf(&codeerror.CodeError{}))
			fmt.Printf("Fatal error: %s\n", p)
		}
	}()

	//错误
	a, err := Sqrt(-1)
	fmt.Println("a=", a) //
	if err != nil {
		//fmt.Println(err)
		panic(err)
	}

	fmt.Println("testError out")
}

func main() {
	fmt.Println("main in")

	testInterface()
	testError()

	fmt.Println("main out")
}

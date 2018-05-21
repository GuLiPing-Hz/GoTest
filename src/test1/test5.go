package main

import "fmt"

//go语言接口
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

func (a *ManBePerson) getName() string {
	return a.name
}

type ManBeAnimal Man

func (a *ManBeAnimal) getName() string {
	return "Human"
}

func main() {
	var animal1 Animal
	animal1 = new(Cat) //必须实现所有接口才能new
	fmt.Println("Animal name is", animal1.getName())

	var animal2 Animal
	animal2 = new(Dog)
	fmt.Println("Animal name is", animal2.getName())

}

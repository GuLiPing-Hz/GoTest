package main

import (
	"fmt"
	"sync"
	"container/list"
	"reflect"
	"strings"
)

func test() {

	fmt.Println("test1")

	if false {
		goto end
	}

	fmt.Println("test2")

end: //用法类似c，good
	fmt.Println("the end")
}

type robotSerer struct {
	mx          sync.Mutex
	hostMethods list.List //频繁增加删除，还是用list实现吧
}

func (rs *robotSerer) Range(cb func(val interface{})) {
	fmt.Println("Range")
	rs.mx.Lock()
	fmt.Println(rs.mx)
	defer rs.mx.Unlock()

	for it := rs.hostMethods.Front(); it != nil; it = it.Next() {
		cb(it.Value)
	}
}

func testList() {
	fmt.Println("testList")
	rs := robotSerer{}
	fmt.Println(rs)

	rs.hostMethods.PushBack(1)
	rs.hostMethods.PushBack("a")
	rs.hostMethods.PushBack(int64(1))
	rs.hostMethods.PushFront(int8(10))

	fmt.Println(rs)

	rs.Range(func(val interface{}) {
		fmt.Println("range", val, reflect.TypeOf(val))
	})
}

type atomicSet struct {
	set sync.Map
}

func (as *atomicSet) Range(cb func(val interface{}) bool) {
	if cb == nil {
		return
	}

	as.set.Range(func(key, value interface{}) bool {
		return cb(key)
	})
}

func (as *atomicSet) Insert(val interface{}) {
	as.set.Store(val, true)
}

func (as *atomicSet) Remove(val interface{}) {
	as.set.Delete(val)
}

func testSet() {
	as := atomicSet{}
	as.Insert(1)
	as.Insert("1")
	as.Insert(int64(1))

	as.Range(func(val interface{}) bool {
		fmt.Println("range 1:", val, reflect.TypeOf(val))
		return true
	})

	as.Remove(1)
	as.Insert("1")

	fmt.Println(strings.Repeat("*", 60))
	as.Range(func(val interface{}) bool {
		fmt.Println("range 2:", val, reflect.TypeOf(val))
		return true
	})
	fmt.Println(strings.Repeat("*", 60))
	as.Range(func(val interface{}) bool {
		fmt.Println("range 3:", as.set)
		fmt.Println(val, reflect.TypeOf(val))
		if val == "1" { //遍历的时候删除某个元素
			fmt.Println("remove")
			as.Remove(val)
		}
		fmt.Println(strings.Repeat("=", 60))
		fmt.Println("range 4:", as.set)
		return true
	})
	fmt.Println(strings.Repeat("*", 60))
	as.Range(func(val interface{}) bool {
		fmt.Println("range 5:", val, reflect.TypeOf(val))
		return true
	})
}

func main() {
	test()
	testList()
	testSet()
}

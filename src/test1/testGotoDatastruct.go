package main

import (
	"fmt"
	"sync"
	"container/list"
	"reflect"
	"strings"
	"encoding/json"
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

type roomStatus struct {
	RoomId   string               //对于需要Json解析/反解析的数据，我们必须使用大写，否则json.Marshal无法获取数据
	RealCnt  int `json:"realCnt"` //指定json序列化的名字
	RobotCnt int
	TotalCnt int
}

func testJson() {
	room := roomStatus{"100", 2, 2, 4}
	bytes, err := json.Marshal(room)
	fmt.Println(bytes, err)

	var room2 roomStatus
	json.Unmarshal(bytes, &room2)
	fmt.Println(room2, err)

	var rooms []roomStatus
	var a1 = roomStatus{"101", 1, 1, 2}
	var a2 = roomStatus{"102", 1, 1, 2}
	var a3 = roomStatus{"103", 1, 1, 2}
	rooms = append(rooms, a1)
	rooms = append(rooms, a2)
	rooms = append(rooms, a3)

	var roomPointers []*roomStatus
	roomPointers = append(roomPointers, &a1)
	roomPointers = append(roomPointers, &a2)
	roomPointers = append(roomPointers, &a3)

	//go数据序列化成字节
	bytes1, err := json.Marshal(rooms)
	fmt.Println("json rooms1=", bytes1, err)

	bytes2, err := json.Marshal(roomPointers)
	fmt.Println("json rooms2=", bytes2, err)
	var rooms2 []roomStatus
	json.Unmarshal(bytes2, &rooms2)
	fmt.Println("roomPointers2 = ", rooms2)

	var roomPointers2 []*roomStatus
	json.Unmarshal(bytes2, &roomPointers2)
	fmt.Println("roomPointers2 = ", roomPointers2)

	//json通用数据解析
	var rooms3 interface{}
	json.Unmarshal(bytes1, &rooms3)
	fmt.Println(rooms3, err)
	rooms4 := rooms3.([]interface{}) //json通用数据解析成数组
	for k, v := range rooms4 {
		fmt.Println(k, v)

		v1 := v.(map[string]interface{}) //jso通用数据解析成map
		for k1, v1 := range v1 {
			fmt.Println(k1, v1)
		}
	}
}

var testMx = sync.Mutex{}

func testMutex() {
	testMx.Lock()
	defer testMx.Unlock()
	fmt.Println("testMutex 1")
	testMx.Lock()
	fmt.Println("testMutex 2")
}

type RobotUser struct {
	UID  int64  `json:"uid"`      //UID
	Name string `json:"nickname"` // 昵称
	//小写大头的字段 或者 `json:"-"` 表示不要序列化字段
}

func testSyncMap() {
	var m1 sync.Map

	m1.Store(1, &RobotUser{1, "a"})
	m1.Store(2, &RobotUser{2, "b"})
	m1.Store(3, &RobotUser{3, "c"})

	var ru4 *RobotUser = new(RobotUser)
	ru4.UID = 4
	ru4.Name = "d"
	m1.Store(4, ru4)

	v, ok := m1.Load(1)
	if ok {
		fmt.Println(v)

		var v1 = v.(*RobotUser)
		aa := RobotUser{1, "aa"}

		fmt.Println("尝试更新map数据 1", v, v1)

		//尝试更新map数据
		v = &aa
		v1 = &aa
		fmt.Println("尝试更新map数据 2", v, v1)
		v, _ := m1.Load(1)
		fmt.Println("结果重新load：使用指针更新map", v) //这个没用
		m1.Store(1, &aa)

		v, _ = m1.Load(1)
		fmt.Println("使用store更新map", v) //这个有效更新到map
	}

	fmt.Println("testSyncMap id 4")
	v, ok = m1.Load(4)
	if ok {
		fmt.Println(v)

		var v1 = v.(*RobotUser)
		aa := RobotUser{4, "dd"}

		fmt.Println("尝试更新map数据 1", v, v1)

		//尝试更新map数据
		v = &aa
		v1 = &aa
		fmt.Println("尝试更新map数据 2", v, v1)
		v, _ := m1.Load(4)
		fmt.Println("结果重新load：使用指针更新map", v) //这个没用
		m1.Store(1, &aa)

		v, _ = m1.Load(1)
		fmt.Println("使用store更新map", v) //这个有效更新到map
	}
}

func main() {
	//test()
	////测试go包数据结构List 频繁插入删除较优
	//testList()
	////测试自定义数据结构Set
	//testSet()
	////测试json解析
	testJson()
	//testMutex()
	testSyncMap()
}

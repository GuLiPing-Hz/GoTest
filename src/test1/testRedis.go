package main

import (
	"github.com/go-redis/redis"
	"fmt"
	"reflect"
	"errors"
)

/*
命令参数
停止redis redis-cli -a xxx shutdown
启动redis redis-server /usr/local/redis/etc/redis.conf
*/

type UserInfo struct {
	Uid  int64
	Name string
}

var (
	ErrParam = errors.New("invalid argument")
)

func converStruct2Map(data interface{}) (map[string]interface{}, error) {
	dataType := reflect.TypeOf(data) //获取数据类型
	dataVal := reflect.ValueOf(data) //获取到数据的值

	if dataType.Kind() == reflect.Ptr {
		//dataType2 := dataType.Elem()
		dataVal2 := dataVal.Elem()
		return converStruct2Map(dataVal2)
	} else
	if dataType.Kind() == reflect.Map {
		return data.(map[string]interface{}), nil
	} else
	if dataType.Kind() == reflect.Struct {
		var ret = make(map[string]interface{})
		for i := 0; i < dataType.NumField(); i++ {
			var sfKey reflect.StructField = dataType.Field(i)
			var sfval = dataVal.Field(i)

			//过滤结构体中的ptr，map，slice，array，struct，chan，func
			if sfval.Kind() == reflect.Invalid || sfval.Kind() == reflect.Uintptr ||
				sfval.Kind() == reflect.Complex64 || sfval.Kind() == reflect.Complex128 ||
				sfval.Kind() == reflect.Array || sfval.Kind() == reflect.Chan ||
				sfval.Kind() == reflect.Func || sfval.Kind() == reflect.Interface ||
				sfval.Kind() == reflect.Map || sfval.Kind() == reflect.Ptr ||
				sfval.Kind() == reflect.Slice || sfval.Kind() == reflect.Struct ||
				sfval.Kind() == reflect.UnsafePointer {
				continue
			}

			//fmt.Println(i, "sf=", sf)
			var key = sfKey.Tag.Get("json") //优先查找json Tag
			if len(key) == 0 {
				key = sfKey.Name
			}

			if sfval.Kind() == reflect.Bool {
				ret[key] = sfval.Bool()
			} else if sfval.Kind() == reflect.Int {
				ret[key] = sfval.Int()
			} else if sfval.Kind() == reflect.Uint {
				ret[key] = sfval.Uint()
			} else if sfval.Kind() == reflect.String {
				ret[key] = sfval.String()
			} else if sfval.Kind() == reflect.Float32 || sfval.Kind() == reflect.Float64 {
				ret[key] = sfval.Float()
			}
		}
		return ret, nil
	} else {
		return nil, ErrParam
	}
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "121.196.203.52:6379",
		Password: "C9BE6E8D-F2CF-4154-BD34-922844BEAC11",
		DB:       0, // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	if err != nil {
		return
	}

	err = client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist

	//hset 设置hashmap key field value
	//hmset 批量设置hashmap key field value field value
	var arra = make([]int, 0)
	arra = append(arra, 1)
	arra = append(arra, 2)

	var dict = make(map[string]interface{})
	/** *******切片类型不允许放入redis*/
	//dict["a"] = arra               //slice
	/** *******切片类型不允许放入redis*/
	//dict["b"] = []string{"你", "好"} //array
	dict["c"] = "name" //string
	/** *******函数类型不允许放入redis*/
	//dict["d"] = func() string {
	//	return "Hello Func"
	//}
	ui := UserInfo{1, "ddd"}
	fmt.Println(ui)
	/** *******结构体不允许放入redis*/
	//dict["e"] = ui
	/** *******结构体指针类型不允许放入redis*/
	//dict["f"] = &ui

	//hashmap 可以存用户信息
	status := client.HMSet("test", dict)
	if status.Err() != nil {
		fmt.Println("status=", status)
	}
	fmt.Println("OK", status.Val())

	status1 := client.HGetAll("test")
	fmt.Println("client.HGetAll", status1)

	//set 可以存在线列表
	fmt.Println("client.SAdd", client.SAdd("onlines", 10001, 10002))
	fmt.Println("client.SMembers", client.SMembers("onlines"))

	fmt.Println("client.Exists", client.Exists("onlines"))
	fmt.Println("client.Del", client.Del("onlines"))
	fmt.Println("client.Exists", client.Exists("onlines"))

	userKey := fmt.Sprintf("user_%d", ui.Uid)
	userMap, err := converStruct2Map(ui)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(client.HMSet(userKey, userMap))
	}

	client.Close()
}

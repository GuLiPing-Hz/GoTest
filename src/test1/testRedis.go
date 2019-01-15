package main

import (
	"github.com/go-redis/redis"
	"fmt"
	"reflect"
	"errors"
	"github.com/mitchellh/mapstructure"
	"time"
)

/*
命令参数
停止redis redis-cli -a xxx shutdown
启动redis redis-server /usr/local/redis/etc/redis.conf
*/

type UserInfo struct {
	Uid      int
	Age      uint
	Name     string
	Percent  float32
	Percent2 float64
	Flag     int8

	//redis中的信息版本
	VerInRedis int64
}

const (
	//redis 缓存信息的版本号，从数据库第一次取出来放进去是0
	VerInRedis = "VerInRedis"

	//redis写入标志位 请求写入的时候 incrBy 1,如果返回1表示可以写入，大于1表示已有其他服务器正在写入，要么等待要么直接返回
	WritingInRedis = "WritingInRedis"
)

var (
	ErrParam = errors.New("invalid argument")
)

/**
这个方法实现了 mapstructure 从Struct到Map的转换 的原理
 */
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
			} else if sfval.Kind() == reflect.Int || sfval.Kind() == reflect.Int8 ||
				sfval.Kind() == reflect.Int16 || sfval.Kind() == reflect.Int32 ||
				sfval.Kind() == reflect.Int64 {
				ret[key] = sfval.Int()
			} else if sfval.Kind() == reflect.Uint || sfval.Kind() == reflect.Uint8 ||
				sfval.Kind() == reflect.Uint16 || sfval.Kind() == reflect.Uint32 ||
				sfval.Kind() == reflect.Uint64 {
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

func Decode(input, output interface{}) error {
	return DecodeEx(input, output, false)
}

func DecodeRedis(input, output interface{}) error {
	return DecodeEx(input, output, true)
}

func DecodeEx(input, output interface{}, weakly bool) error {
	dataType := reflect.TypeOf(output) //获取数据类型
	if dataType.Kind() != reflect.Ptr {
		return fmt.Errorf("need Ptr")
	}

	config := &mapstructure.DecoderConfig{
		Metadata:         nil,
		Result:           output,
		TagName:          "json",
		WeaklyTypedInput: weakly,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	return decoder.Decode(input)
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

	// 模拟从数据库取出数据，版本号为0
	ui := UserInfo{1, 2018, "Jack", 0.9, 0.8, 1, 0}
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
	fmt.Println("client.SCard", client.SCard("onlines"))
	fmt.Println("client.SRem", client.SRem("onlines", 10002))
	fmt.Println("client.SCard", client.SCard("onlines"))

	fmt.Println("client.Exists", client.Exists("onlines"))
	fmt.Println("client.Del", client.Del("onlines"))
	fmt.Println("client.Exists", client.Exists("onlines"))

	userKey := fmt.Sprintf("user_%d", ui.Uid)
	//读取redis
	if exits, _ := client.Exists(userKey).Result(); exits == 1 {
		var failRetry = 3
		for {
			//先查询服务器当前的缓存版本
			sc := client.HIncrBy(userKey, VerInRedis, 0) //+0表示我们只是为了查看一下当前的值
			if sc.Err() == nil {
				if sc.Val() == ui.VerInRedis {
					//如果redis中的版本跟我们的一致，那么我们可以改写它。否则直接返回失败、或者读取里面的数据
					break
				} else {
					fmt.Println("redis中版本跟服务器不一致", sc.Val(), "<>", ui.VerInRedis)
					sc := client.HGetAll(userKey)
					if sc.Err() == nil {
						err := DecodeRedis(sc.Val(), &ui)
						if err == nil {
							break
						} else {
							fmt.Println("err=", err)
						}
					}

					failRetry --
					if failRetry == 0 {
						fmt.Println("从redis读取数据尝试超过最大次数3")
						break
					}
				}
			}
		}
	} else {
		ui.VerInRedis = 0
	}

	for {
		sc2 := client.HIncrBy(userKey, WritingInRedis, 1) //+1表示我们需要往里写入数据
		if sc2.Err() != nil {
			if sc2.Err() == redis.Nil {
				break
			}
			time.Sleep(time.Millisecond * 1) //等待1毫秒
			continue
		}

		if sc2.Val() != 1 {
			client.HIncrBy(userKey, WritingInRedis, -1) //还原回去
			time.Sleep(time.Millisecond * 1)            //等待1毫秒
			continue
		}

		//这时候我们就可以堂堂正正的写数据了。
		//写入redis
		status2 := client.HIncrBy(userKey, "VerInRedis", 0)
		if status2.Err() == nil {
			if status2.Val() == ui.VerInRedis {
				fmt.Println("redis版本跟服务器的一致，可以写入")

				ui.VerInRedis ++
				var usrMap map[string]interface{}
				//userMap, err := converStruct2Map(ui) //这个方法可以让我们学习到反射转换的原理
				err := Decode(ui, &usrMap)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(client.HMSet(userKey, usrMap))
				}
			} else {
				fmt.Println("redis版本跟服务器的一致，不能写入")
			}
		}
	}

	client.Close()
}

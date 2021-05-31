package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func GetRandSeed() int64 {
	var a = 0 //变量地址当做随机数
	var b = 0 //变量地址当做随机数
	aPtr, _ := strconv.ParseInt(fmt.Sprintf("%x", &a), 16, 64)
	bPtr, _ := strconv.ParseInt(fmt.Sprintf("%x", &b), 16, 64)

	return time.Now().Unix() * aPtr * bPtr
}

func main() {
	rand.Seed(time.Now().Unix())
	var cnt = make(map[int]int, 0)
	for i := 1; i < 10000; i++ {
		r := int(rand.Int31n(10))
		//fmt.Println("rand(", i, ")=", r)
		_, exist := cnt[r]
		if exist {
			cnt[r] ++
		} else {
			cnt[r] = 0
		}

		fmt.Println("GetRandSeed=", GetRandSeed())
	}
	fmt.Println(cnt)
}

package main

import (
	"fmt"
	"time"
)

func testTimer() {
	//After的使用
	fmt.Println("main")
	//time.NewTimer(time.Second*2).C
	t := time.After(time.Second * 3) //返回一个只读的信道
	fmt.Println("t=", t)
	val := <-t //我们读取它，
	fmt.Println("val=", val)

	//AfterFunc的使用
	done := make(chan bool)
	fmt.Println("AfterFunc 1")
	time.AfterFunc(time.Second*3, func() {
		fmt.Println("AfterFunc 2")
		done <- true
	})
	fmt.Println("AfterFunc 3")
	<-done
	fmt.Println("AfterFunc 4")

	//NewTicker的使用
	tick := time.NewTicker(time.Second * 2)
	cnt := 0
	fmt.Println("NewTicker 1")
	for {
		fmt.Println("NewTicker 2")
		curTime := <-tick.C
		fmt.Println("curTime=", curTime, cnt)
		cnt ++

		if cnt == 5 {
			tick.Stop()
			break // end
		}
	}

	fmt.Println("tick=", tick)
	tick.Stop()
	fmt.Println("tick=", tick)
	tick = nil
	//end:
	fmt.Println("NewTicker 3", tick)
}

func testTime() {
	fmt.Println("time1", time.Now().Format(time.ANSIC))
	fmt.Println("time2", time.Now().Format(time.UnixDate))
	fmt.Println("time3", time.Now().Format(time.RFC3339))
	fmt.Println("time4", time.Now().Format(time.RFC3339Nano))
	fmt.Println("time5", time.Now().Format("2006-01-02 15:04:05.000"))
}

func main() {
	fmt.Println("Unix time 单位秒", time.Now().Unix())

	testTime()
	testTimer()
}

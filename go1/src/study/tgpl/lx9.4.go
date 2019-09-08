package main

import "fmt"

func main() {
	ch1 := make(chan int32)
	ch2 := make(chan int32)
	for i := 0; i < 1000000000; i++ {
		fmt.Printf("创建第 %d 个goroutine\n", i)

		//创建第 3127762 个goroutine
		//runtime: VirtualAlloc of 1048576 bytes failed with errno=1455
		if i > 0 {
			ch1 = ch2
			ch2 = make(chan int32)
		}
		func(ch1, ch2 chan int32, i int32) {
			go func() {
				for {
					v := <-ch1
					fmt.Printf("第 %d 道工序 = %d \n", i, v)
					ch2 <- v
				}
			}()
		}(ch1, ch2, int32(i))
	}
}

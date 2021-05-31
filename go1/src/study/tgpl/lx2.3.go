package main

import (
	"fmt"
	"go1/src/pkglearn"
	"time"
)

func main() {
	fmt.Printf("tm1=%v\n", time.Now())
	//貌似表达式最慢。。
	fmt.Printf("tm2=%v,v=%d\n", time.Now(), pkglearn.PopCount(45982254555))

	//下面两个速度差不多。。
	fmt.Printf("tm3=%v,v=%d\n", time.Now(), pkglearn.PopCount2_3(45982254555))
	fmt.Printf("tm4=%v,v=%d\n", time.Now(), pkglearn.PopCount2_4(45982254555))
	fmt.Printf("tm5=%v,v=%d\n", time.Now(), pkglearn.PopCount2_5(45982254555))
}

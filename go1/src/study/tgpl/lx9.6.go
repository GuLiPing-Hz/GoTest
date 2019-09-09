package main

import (
	"flag"
	"fmt"
	"pkg/B"
	"runtime"
)

func main() {
	//fmt.Printf("A=%d\n", A.A) //内部包，这里无法访问，只能由同一级目录才能访问
	//fmt.Printf("a=%d\n", A.a) //小写没有导出，无法访问。
	fmt.Printf("B=%d\n", B.B) //外部包，这里可以访问
	//fmt.Printf("b=%d\n", B.b) //小写没有导出，无法访问。

	var maxgr int
	flag.IntVar(&maxgr, "maxgr", 4, "max goroutine")
	runtime.GOMAXPROCS(maxgr) //设置同时最大运行的goroutines数目
	for {
		go fmt.Print(0)
		fmt.Print(1)
	}
}

/*
GOMAXPROCS=1 go run lx9.6.go
GOMAXPROCS=1 go run lx9.6.go
*/


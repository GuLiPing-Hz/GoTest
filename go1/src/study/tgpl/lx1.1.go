//如果包名不为main的话，在这个包里写的main方法将无法被调用。
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	//the Go programming language
	//练习题1.1
	fmt.Println(os.Args)
	//练习题1.2
	for i, s := range os.Args {
		fmt.Println(i, s)
	}
	//练习题1.3
	var strs []string
	for i := 0; i < 5000; i++ {
		strs = append(strs, fmt.Sprintf("%d", i))
	}
	var log, sep string
	fmt.Println(time.Now(), "普通循环拼接")
	for _, s := range strs {
		log += sep + s //粗糙的字符串拼接。。
		sep = " "
	}
	fmt.Println(time.Now(), log[:100])
	log = ""
	fmt.Println(time.Now(), "strings.join拼接")
	log = strings.Join(strs, " ") //性能完胜啊。。
	fmt.Println(time.Now(), log[:100])
}

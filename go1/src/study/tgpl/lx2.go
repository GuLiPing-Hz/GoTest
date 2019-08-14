package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func stdinScanner() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() { // Scan 函数在读到一行时返回 true ，在 无输入时返回 false
		// ==> ctrl+c 结束输入

		//map中默认空值取出来是0
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			//go fmt 格式化
			//%d 十进制整数
			//%x, %o, %b 十六进制，八进制，二进制整数。
			//%f, %g, %e 浮点数： 3.141593 3.141592653589793 3.141593e+00
			//%t 布尔：true或false
			//%c 字符（rune） (Unicode码点)
			//%s 字符串
			//%q 带双引号的字符串"abc"或带单引号的字符'c'
			//%v 变量的自然形式（natural format）
			//%T 变量的类型
			//%% 字面上的百分号标志（无操作数）
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func main() {
	//stdinScanner()
	//练习题1.4
	counts := make(map[string]int)
	files := make(map[string]string)

	paramsFiles := os.Args[1:]
	for i := range paramsFiles {
		buf, err := ioutil.ReadFile(paramsFiles[i])
		if err != nil {
			continue
		}

		lines := strings.Split(string(buf), "\n")
		for j := range lines {
			counts[lines[j]] ++
			files[lines[j]] = paramsFiles[i]
		}
	}

	fmt.Println("result:")
	for i := range counts {
		if counts[i] > 1 {
			fmt.Println(i, counts[i], files[i])
		}
	}
}

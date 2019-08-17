package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func fetch(url string, ch chan<- string) {
	start := time.Now()
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}

	//练习题1.11
	client := &http.Client{}
	client.Timeout = time.Second * 10 //设置一个超时时间。
	resp, err := client.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("err=%s\t%s\n", err.Error(), url)
		return
	}

	cnt, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("err=%s\t%s\n", err.Error(), url)
		return
	}

	ch <- fmt.Sprintf("%0.2fs\t%d\t%s\n", time.Since(start).Seconds(), cnt, url)
}

//.\go_build_lx5_go.exe www.google.com www.baidu.com www.qq.com
func main() {

	args := os.Args[1:]
	ch := make(chan string) //无缓冲的channel
	for i := range args {
		go fetch(args[i], ch)
		//练习题1.10
		go fetch(args[i], ch)
	}

	for _ = range args {
		fmt.Print(<-ch)
		//练习题1.10
		fmt.Print(<-ch)
	}
}

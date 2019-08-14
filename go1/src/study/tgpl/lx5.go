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

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("err=%s\t%s", err.Error(), url)
		return
	}

	cnt, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("err=%s\t%s", err.Error(), url)
		return
	}

	ch <- fmt.Sprintf("%0.2fs\t%d\t%s", time.Since(start).Seconds(), cnt, url)
}

func main() {

	args := os.Args[1:]
	ch := make(chan string) //无缓冲的channel
	for i := range args {
		go fetch(args[i], ch)
		//练习1.10
		go fetch(args[i], ch)
	}

	for _ = range args {
		fmt.Print(<-ch)
		fmt.Print(<-ch)
	}
}

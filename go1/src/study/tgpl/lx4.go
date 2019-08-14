package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	url := os.Args[1]
	//练习1.8
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("err=%v\n", err.Error())
		os.Exit(1)
	}

	//练习1.7
	_, err = io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Printf("err=%v\n", err.Error())
		os.Exit(1)
	}

	//练习1.9
	fmt.Printf("http status:%s", resp.Status)
}

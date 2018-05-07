package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path:", r.URL.Path)
	fmt.Println("scheme:", r.URL.Scheme)
	fmt.Println(r.Form["url_long:"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "<h1>Hello astaxie!</h1>") //这个写入到w的是输出到客户端的
}

func sayMore(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Print("r.Form=")
	fmt.Print(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("")
	fmt.Print("r.PostForm=")
	fmt.Print(r.PostForm)
	fmt.Println("")
	fmt.Println("path:", r.URL.Path)
	fmt.Println("scheme:", r.URL.Scheme)
	strUrlLong := r.Form["url_long:"]
	fmt.Print("r.Form['a']=")
	fmt.Print(strUrlLong)
	fmt.Println("")
	fmt.Println("r.Form['a']=" + strUrlLong[0])
	for k, v := range r.Form {
		fmt.Print("key:", k)
		fmt.Print("val:", strings.Join(v, ";"))
	}
	fmt.Println("")
	fmt.Fprintf(w, "<h1>Hello More!</h1>") //这个写入到w的是输出到客户端的
}

func main() {
	http.HandleFunc("/", sayhelloName)       //设置访问的路由
	http.HandleFunc("/more/", sayMore)       //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

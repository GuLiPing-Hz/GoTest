package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
)

func printRequest(r *http.Request) {
	fmt.Println("r.Form=", r.Form)         //这些信息是输出到服务器端的打印信息 , Get参数
	fmt.Println("r.PostForm=", r.PostForm) //Post参数
	fmt.Println("path=", r.URL.Path)
	fmt.Println("scheme=", r.URL.Scheme)
	fmt.Println("method=", r.Method) //获取请求的方法

	fmt.Println("Http Get参数列表 begin:")
	for k, v := range r.Form {
		fmt.Println("Http Get["+k+"]=", strings.Join(v, " ; "))
	}
	fmt.Println("Http Get参数列表 end:")

	fmt.Println("Http Post参数列表 begin:")
	for k, v := range r.PostForm {
		fmt.Println("Http Post["+k+"]=", strings.Join(v, " ; "))
	}
	fmt.Println("Http Post参数列表 end:")

	arraA := r.Form["a"]
	fmt.Println("r.Form['a']=", arraA)
	if len(arraA) > 0 {
		fmt.Println("r.Form['a'][0]=", arraA[0])
	}
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println("******************************************sayhelloName")
	printRequest(r)

	fmt.Fprintf(w, "<h1>Hello 世界!</h1>") //这个写入到w的是输出到客户端的
}

func sayMore(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println("******************************************sayMore")
	printRequest(r)

	fmt.Fprintf(w, "<h1>Hello More!</h1>") //这个写入到w的是输出到客户端的
}

func sayMore1(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println("******************************************sayMore1")
	printRequest(r)

	fmt.Fprintf(w, "<h1>Hello More1!</h1>") //这个写入到w的是输出到客户端的
}

func main() {
	http.HandleFunc("/", sayhelloName)       //设置访问的路径
	http.HandleFunc("/more", sayMore)        //设置访问的路径
	http.HandleFunc("/more/", sayMore1)      //设置访问的路径
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

package main

import (
	"fmt"
	"net/http"
	"strings"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"io/ioutil"
)

//注意：在使用如下命令的时候，由于没有在window环境变量中添加GOPATH，导致一直下载失败，必须添加GOPATH环境变量才能安装
//go get -u github.com/go-sql-driver/mysql

//import "fmt"最常用的一种形式
//import "./test"导入同一目录下test包中的内容
//import f "fmt"导入fmt，并给他启别名ｆ
//import . "fmt"，将fmt启用别名"."，这样就可以直接使用其内容，而不用再添加ｆｍｔ，如fmt.Println可以直接写成Println
//import  _ "fmt" 表示不使用该包，而是只是使用该包的init函数，并不显示的使用该包的其他内容。注意：这种形式的import，当import时就执行了fmt包中的init函数，而不能够使用该包的其他函数。

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

	var arraA = r.Form["a"]
	//arraA := r.Form["a"]
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

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println("******************************************sayMore1")
	printRequest(r)

	var account = r.Form["account"]
	var pwd = r.Form["pwd"]

	if len(account) == 0 || len(pwd) == 0 {
		fmt.Fprintf(w, "<h1>参数提交异常</h1>") //这个写入到w的是输出到客户端的
	} else if account[0] == "" || pwd[0] == "" || len(pwd[0]) != 32 {
		fmt.Fprintf(w, "<h1>参数格式异常</h1>") //这个写入到w的是输出到客户端的
	} else {
		fmt.Fprintf(w, "<h1>欢迎注册,"+account[0]+"</h1>") //这个写入到w的是输出到客户端的

		var (
			dbhostsip  = "192.168.0.18:3306" //IP地址
			dbusername = "root"              //用户名
			dbpassword = "Glp470312"         //密码
			dbname     = "databasetest"      //数据库名
		)

		_, err := sql.Open("mysql", dbusername + ":"+
			dbpassword+ "@tcp("+ dbhostsip+ ")/"+ dbname+ "?charset=utf8mb4")
		checkErr(err)
	}
}

func testServer() {
	http.HandleFunc("/", sayhelloName)       //设置访问的路径
	http.HandleFunc("/more", sayMore)        //设置访问的路径
	http.HandleFunc("/more/", sayMore1)      //设置访问的路径
	http.HandleFunc("/register", register)   //设置访问的路径
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func testClient() {
	response, _ := http.Get("http://www.baidu.com")
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}

func main() {
	testClient()
	testServer()
}

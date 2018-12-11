package main

import (
	"fmt"
	"net/http"
	"strings"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"io/ioutil"
	"net/url"
	"encoding/json"
	"test1/tool"
	"crypto/sha1"
	"sort"
)

var DbMgr = tool.DBMgr{}
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
	fmt.Println("r.Form['a']=", arraA) //获取到的是数组
	var vA = r.Form.Get("a")           //最好使用Get函数获取一个，
	fmt.Println(`r.Form.Get("a")=`, vA)
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
	fmt.Println("******************************************register")
	printRequest(r)

	var account = r.Form.Get("account")
	var pwd = r.Form.Get("pwd")

	if len(account) == 0 || len(pwd) == 0 {
		fmt.Fprintf(w, "<h1>参数提交异常</h1>") //这个写入到w的是输出到客户端的
	} else if account == "" || pwd == "" || len(pwd) != 32 {
		fmt.Fprintf(w, "<h1>参数格式异常</h1>") //这个写入到w的是输出到客户端的
	} else {
		//fmt.Fprintf(w, "<h1>登录成功,"+account+"</h1>") //这个写入到w的是输出到客户端的

		//sqlDatas, err := DbMgr.LoadTable(`SELECT account From databasetest.tabuser WHERE account=? AND pwd=?`, account, pwd)
		sqlDatas, err := DbMgr.LoadTable(`SELECT account From databasetest.tabuser WHERE account=?`, account)
		if err != nil {
			fmt.Fprintf(w, "<h1>登录失败,err=%s</h1>", err.Error()) //这个写入到w的是输出到客户端的
			return
		}

		fmt.Println("sqlDatas =", sqlDatas)
		if len(sqlDatas) <= 0 {
			fmt.Fprintf(w, "<h1>登录失败! 账号不存在或者密码错误</h1>") //这个写入到w的是输出到客户端的
			return
		}

		fmt.Fprintf(w, "<h1>登录成功! 欢迎%s</h1>", account) //这个写入到w的是输出到客户端的
	}
}

func httpTestJson(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	printRequest(r)

	//fmt.Println(r.Body)

	ct := r.Header.Get("Content-Type")
	fmt.Println(ct)
	if ct != "application/json" {
		w.WriteHeader(http.StatusBadRequest) //http.StatusOK
		return
	}

	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close() //这里需要自己关闭

	fmt.Println("httpTestJson body=", string(body))

	var result = map[string]interface{}{}
	result["code"] = 0

	bs, _ := json.Marshal(result)
	fmt.Fprintf(w, string(bs))
}

func httpCreate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	printRequest(r)

	var sign = r.Header.Get("sign")

	var plain = ""
	var dicKeys []string
	for k, _ := range r.Form {
		dicKeys = append(dicKeys, k)
	}
	sort.Strings(dicKeys) //go默认升序，可以用reverse获取降序
	fmt.Println(dicKeys)
	for _, v := range dicKeys {
		plain += v + "=" + r.Form[v][0] + "&"
	}

	key := "ZsNXN8e4Mrpzk7GKDyRcdkk1BSJCkBIHtY6UWbQjFgBXPYt40I48bwoOdTIq6QiK"
	plain = key + plain[:len(plain)-1]
	fmt.Println(plain)
	//获取16进制的签名数据
	calcSign := fmt.Sprintf("%X", sha1.Sum([]byte(plain)))

	fmt.Println(sign, calcSign)
	fmt.Fprintf(w, "ok")
}

func testServer() {
	http.HandleFunc("/", sayhelloName)     //设置访问的路径
	http.HandleFunc("/more", sayMore)      //设置访问的路径
	http.HandleFunc("/more/", sayMore1)    //设置访问的路径
	http.HandleFunc("/register", register) //设置访问的路径
	http.HandleFunc("/testJson", httpTestJson)
	http.HandleFunc("/api/orders/create", httpCreate)
	http.HandleFunc("/api/orders/list", httpCreate)

	server := &http.Server{Addr: ":9090", Handler: nil} //设置监听的端口
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func testClient() {
	response, _ := http.Get("http://www.baidu.com")
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))

	//http.PostForm()

	client := &http.Client{}

	param := url.Values{}
	param.Add("a", "1")
	param.Add("b", "2")
	paramEnc := param.Encode()
	fmt.Println(paramEnc)
	//request, err := http.NewRequest("POST", "http://127.0.0.1:9090/more?"+paramEnc, nil)
	request, err := http.NewRequest("POST", "http://127.0.0.1:9090/more", strings.NewReader(param.Encode()))
	if err != nil {
		fmt.Println(err)
		return
	}

	//request.Header.Set("appkey", appKey)
	//request.Header.Set("nonce", nonce)
	//request.Header.Set("curtime", curTime)
	//request.Header.Set("checksum", checksum)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, _ = client.Do(request)
	defer response.Body.Close()
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		bodystr := string(body)
		fmt.Println(bodystr)
	}
}

func main() {
	var (
		dbhostsip  = "192.168.0.18:3306" //IP地址
		dbusername = "glp4703"           //用户名
		dbpassword = `glp3329`           //密码
		dbname     = "databasetest"      //数据库名
	)

	pwd := dbpassword //string(sha2.Write)
	//fmt.Println("pwd=", pwd)

	//可以通过mysql的命令进入数据库，查询当前的用户信息
	//select * from mysql.user
	//host指定了访问来源，
	//db, err := sql.Open("mysql", dbusername + ":"+
	//	pwd+ "@tcp("+ dbhostsip+ ")/"+ dbname+ "?charset=utf8mb4")
	var err error
	DbMgr.DbInst, err = sql.Open("mysql", dbusername + ":"+
		pwd+ "@tcp("+ dbhostsip+ ")/"+ dbname+ "?charset=utf8mb4")
	if err != nil {
		fmt.Printf("SQL Open Err=%s\n", err)
	}
	defer DbMgr.DbInst.Close()

	//go
	testServer()
	//time.Sleep(time.Second * 3)
	//testClient()
	//time.Sleep(time.Hour)
}

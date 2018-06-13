package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //go get github.com/go-sql-driver/mysql
	"test1/tool"
	"strings"
	"github.com/mitchellh/mapstructure"
	"reflect"
)

func testMySql(dbMgr *tool.DBMgr) {
	fmt.Printf("SQL Open OK %s\n", dbMgr.DbInst)

	// Execute the query
	rows, err := dbMgr.DbInst.Query(`SELECT * FROM user`)
	if err != nil {
		fmt.Println(err.Error()) // proper error handling instead of panic in your app
		return
	}

	// Get column names 获取列名
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Make a slice for the values 为值创建切片
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	for rows.Next() { //遍历表数据行
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("-----------------------------------")
	}
	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}

type DbUser struct {
	RowId int    `json:"rowId"`
	Uid   int64  `json:"uid"`
	Name  string `json:"name"`
}

func (du DbUser) SayHello(name string) { //如果类型为 *DbUser，那么反射将无法识别是否结构体的函数
	fmt.Println("Hello", du.Name, ",I'm", name)
}

func (du DbUser) PrintUID() {
	fmt.Println("DbUser.Uid=", du.Uid)
}

func (du DbUser) PrintName() {
	fmt.Println("DbUser.Name=", du.Name)
}

//通过反射来了解更多信息
func testReflect(data interface{}) {
	var dataType2 reflect.Type
	var dataVal2 reflect.Value
	dataType := reflect.TypeOf(data) //获取数据类型
	fmt.Println(dataType)

	dataVal := reflect.ValueOf(data) //获取到数据指针的值
	fmt.Println("dataVal=", dataVal)

	fmt.Println("dataType.Kind() =", dataType.Kind()) //获取数据种类
	fmt.Println("dataVal.Kind() =", dataVal.Kind())   //获取数据种类
	if dataType.Kind() == reflect.Ptr {
		fmt.Println("is Pointer")

		dataType2 = dataType.Elem() //获取数据指针指向的类型  必须是Kind为：Array, Chan, Map, Ptr, or Slice
		fmt.Println("dataType2=", dataType2)

		dataVal2 = dataVal.Elem() //获取数据指针指向的值
		fmt.Println("dataVal2=", dataVal2)

		fmt.Println("dataType2.Kind() =", dataType2.Kind()) //获取数据种类
		fmt.Println("dataVal2.Kind() =", dataVal2.Kind())   //获取数据种类

		if dataType2.Kind() == reflect.Struct { //判断数据是否是结构体
			fmt.Println("dataType2 查看结构体字段", dataType2.NumField(), dataVal2.NumField())
			for i := 0; i < dataType2.NumField(); i++ {
				var sf reflect.StructField = dataType2.Field(i)
				fmt.Println(i, "sf=", sf)
			}

			var args []reflect.Value
			args = append(args, dataVal2)

			fmt.Println("dataType2 查看结构体函数", dataType2.NumMethod(), dataVal2.NumMethod())
			for i := 0; i < dataType2.NumMethod(); i++ {
				var method reflect.Method = dataType2.Method(i)
				fmt.Println(i, "method=", method)
				if method.Func.Type().NumIn() == 1 { //需要几个参数
					//函数调用
					method.Func.Call(args) //跟Lua一样，调用结构体，第一个参数需要传入自己
				} else {

				}
			}

			fmt.Println("调用函数")
			//如果是通过type找函数的话，调用的时候需要带上本身的value
			methodSayHello, ok := dataType2.MethodByName("SayHello")
			if ok {
				args2 := append(args, reflect.ValueOf("Jack From type"))
				methodSayHello.Func.Call(args2)
			}

			//如果是通过value找函数的话，调用的时候不需要带上本身的value
			methodSayHello2 := dataVal2.MethodByName("SayHello")
			if methodSayHello2.IsValid() {
				var args2 []reflect.Value
				args2 = append(args2, reflect.ValueOf("Jack From value"))
				methodSayHello2.Call(args2)
			}

			//修改结构字段的值
			for i := 0; i < dataType2.NumField(); i++ {
				var sfType reflect.StructField = dataType2.Field(i)
				fmt.Println(i, "sfType=", sfType)
				var sfValue reflect.Value = dataVal2.Field(i)

				if sfType.Type.Kind() == reflect.Int {
					sfValue.SetInt(2)
				}
			}
		}
	}

}

func testDbHelp(dbMgr *tool.DBMgr) {
	fmt.Println(strings.Repeat("*", 100))
	data, err := dbMgr.LoadTable(`SELECT * FROM user`)
	if err != nil {
		fmt.Printf("err = %s\n", err.Error())
		return
	}

	fmt.Println(data)

	fmt.Println(strings.Repeat("*", 100))
	data, err = dbMgr.LoadTable(`SELECT * FROM user where uid=?`, 5)
	if err != nil {
		fmt.Printf("err = %s\n", err.Error())
		return
	}

	fmt.Println(data)

	var dbUser []DbUser
	var dbUserType = reflect.TypeOf(&dbUser)
	fmt.Println(dbUserType, dbUserType.Kind(), dbUserType.Elem(), "end")
	mapstructure.Decode(data, &dbUser)
	fmt.Println(dbUser, dbUser[0].Uid)
}

func main() {
	var (
		dbhostsip  = "192.168.0.18:3306" //IP地址
		dbusername = "glp4703"           //用户名
		dbpassword = `glp3329`           //密码
		dbname     = "databasetest"      //数据库名
	)

	pwd := dbpassword //string(sha2.Write)
	fmt.Println("pwd=", pwd)

	//可以通过mysql的命令进入数据库，查询当前的用户信息
	//select * from mysql.user
	//host指定了访问来源，
	//db, err := sql.Open("mysql", dbusername + ":"+
	//	pwd+ "@tcp("+ dbhostsip+ ")/"+ dbname+ "?charset=utf8mb4")
	dbMgr := tool.DBMgr{}
	var err error
	dbMgr.DbInst, err = sql.Open("mysql", dbusername + ":"+
		pwd+ "@tcp("+ dbhostsip+ ")/"+ dbname+ "?charset=utf8mb4")
	if err != nil {
		fmt.Printf("SQL Open Err=%s\n", err)
	}
	defer dbMgr.DbInst.Close()

	dbTemp := DbUser{1, 10, "Hello 世界"}
	testReflect(&dbTemp)
	fmt.Println(dbTemp)

	//testMySql(&dbMgr)
	//testDbHelp(&dbMgr)
}

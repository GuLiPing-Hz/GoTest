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

//通过反射来了解更多信息
func testReflect(data interface{}) {
	dataType := reflect.TypeOf(data)
	fmt.Println(dataType)

	fmt.Println("dataType.Kind() =", dataType.Kind())
	if dataType.Kind() == reflect.Ptr {
		fmt.Println("is Pointer")
	}

	fmt.Println(dataType.)
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
	//testMySql(&dbMgr)
	//testDbHelp(&dbMgr)
}

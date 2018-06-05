package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //go get github.com/go-sql-driver/mysql
)

func testMySql() {
	var (
		dbhostsip  = "192.168.0.18:3306" //IP地址
		dbusername = "root"              //用户名
		dbpassword = "Glp470312"         //密码
		dbname     = "databasetest"      //数据库名
	)

	//sha2 := sha256.New()
	pwd := dbpassword //string(sha2.Sum([]byte(dbpassword)))
	db, err := sql.Open("mysql", dbusername + ":"+
		pwd+ "@tcp("+ dbhostsip+ ")/"+ dbname+ "?charset=utf8mb4")
	if err != nil {
		fmt.Printf("SQL Open Err=%s\n", err)
	}
	defer db.Close()

	fmt.Printf("SQL Open OK %s\n", db)

	// Execute the query
	rows, err := db.Query(`SELECT * FROM user`)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	for rows.Next() {
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

func main() {
	testMySql()
}

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 数据库连接对象可以在多个函数中使用

// 在全局中声明一个db变量，好让所有的函数都能调用
var db *sql.DB

// initDB 初始化连接
func initDB() (err error) {
	// 1.数据库信息
	databaseInfo := "root:root123@tcp(192.168.3.121:3306)/sql_test"
	// 2.校验
	db, err = sql.Open("mysql", databaseInfo)
	if err != nil {
		return
	}
	// 3.连接
	err = db.Ping()
	if err != nil {
		return
	}

	db.SetMaxOpenConns(10) //设置数据库连接池的最大连接数
	db.SetMaxIdleConns(5)  //设置空闲时间最大连接数

	return
}

func main() {
	// 初始化数据库连接
	err := initDB()
	if err != nil {
		fmt.Printf("init DB falied, err:%v\n", err)
		return
	}
	fmt.Println("connect database success!")
}

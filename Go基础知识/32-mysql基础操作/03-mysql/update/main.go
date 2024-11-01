package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 在全局中声明一个db变量，好让所有的函数都能调用
var db *sql.DB

// initDB 初始化连接
func initDB() (err error) {
	//1.数据库信息
	databaseInfo := "root:root123@tcp(192.168.3.121:3306)/sql_test"
	//2.校验
	db, err = sql.Open("mysql", databaseInfo)
	if err != nil {
		return
	}
	//3.连接
	err = db.Ping()
	if err != nil {
		return
	}
	db.SetMaxOpenConns(10) //设置数据库连接池的最大连接数
	db.SetMaxIdleConns(5)  //设置空闲时间最大连接数
	return
}

// updateRow 修改一行数据
func updateRow(newAge int, id int) {
	sqlStr := `update user set age=? where id = ?`
	ret, err := db.Exec(sqlStr, newAge, id)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() //获取受影响的行数
	if err != nil {
		fmt.Printf("get any rows failed, err:%v\n", err)
	}
	fmt.Printf("update %d rows data!\n", n)
}

func main() {
	//初始化数据库连接
	err := initDB()
	if err != nil {
		fmt.Printf("init DB falied, err:%v\n", err)
		return
	}
	fmt.Println("connect database success!")

	//修改一行
	updateRow(9000, 1)

}

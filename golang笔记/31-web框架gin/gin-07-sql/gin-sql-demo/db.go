package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 声明一个全局的db变量
var db *sqlx.DB

// 初始化数据库
func InitDB() (err error) {
	// 1 数据库连接信息
	// 用户名:密码@tcp(IP:端口)/数据库名称
	databaseInfo := `root:root123@tcp(192.168.3.121:3306)/go`
	// 2 连接数据库
	db, err = sqlx.Connect("mysql", databaseInfo)
	if err != nil {
		return err
	}
	// 设置数据库连接池
	db.SetMaxOpenConns(16) // 设置数据库连接池最大的连接
	db.SetMaxIdleConns(4)  // 设置最大空闲连接数
	return
}

// 查询所有数据
func QueryAllData() ([]book, error) {
	sqlStr := `select id, title, price from book`

	var bookList []book
	err := db.Select(&bookList, sqlStr)
	if err != nil {
		return nil, err
	}

	return bookList, err
}

// 插入数据
func InsertBookData(bookName string, bookPrice int64) (err error) {
	sqlStr := `insert into book(title, price) value (?, ?)`
	_, err = db.Exec(sqlStr, bookName, bookPrice)
	if err != nil {
		return err
	}
	return
}

// 删除数据
func DeleteBookData(id string) (err error) {
	sqlStr := `delete from book where id = ?`
	_, err = db.Exec(sqlStr, id)
	if err != nil {
		return err
	}
	return
}

package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//声明全局db变量，方便各个函数调用
var db *sql.DB

//用来查询数据
type user struct {
	id       int
	name     string
	password string
}

//init 导入该包的时候就初始化数据库的连接
func init() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB falied, err:%v\n", err)
		return
	}
}

//initDB 连接数据库
func initDB() (err error) {
	databaseInfo := "root:root123@tcp(192.168.3.121:3306)/sql_test"
	db, err = sql.Open("mysql", databaseInfo)
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	return
}

//QueryOne 注册功能查询用户是否已经注册
func QueryOne(name string) bool {
	var u1 user
	sqlStr := `select id, name, password from userinfo where name = ?`
	//查询数据，如果没查询到，u1为空
	db.QueryRow(sqlStr, name).Scan(&u1.id, &u1.name, &u1.password)
	if u1.name == name {
		//找到了相同的数据.证明已经注册
		return false
	} else {
		return true
	}
}

//QueryLogin 登录功能账号密码合法性校验
func QueryLogin(name, password string) bool {
	//如果账号或者密码为空直接返回false
	if name == "" || password == "" {
		return false
	}
	var u1 user
	sqlStr := `select id, name, password from userinfo where name = ?`
	db.QueryRow(sqlStr, name).Scan(&u1.id, &u1.name, &u1.password)
	if u1.name == name && u1.password == password {
		return true
	}
	return false
}

//insertOne 注册功能插入数据
func InsertOne(name, password string) (bool, error) {
	sqlStr := `insert into userinfo(name,password) value(?,?)`
	_, err := db.Exec(sqlStr, name, password)
	if err != nil {
		return false, err
	}
	return true, nil
}

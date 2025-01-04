package main

/*
[sqlx库使用指南]
使用sqlx实现批量插入数据的例子，介绍了sqlx中可能被忽视了的 sqlx.In 和 DB.NamedExec 方法

sqlx可以认为是Go语言内置 database/sql 的超集
它在优秀的内置 database/sql 基础上提供了一组扩展
这些扩展中除了大家常用来查询的 Get 和 Select 外还有很多其他强大的功能
sqlx中的exec方法与原生sql中的exec使用基本一致 所以增删改应该跟原生的sql库差不多

安装方法:
go get github.com/jmoiron/sqlx
*/

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func main() {
	// 连接数据库初始化
	err := initDB()
	if err != nil {
		fmt.Printf("connect database failed, err:%v\n", err)
		return
	}

	// 查询一条数据
	queryRowDemo(2)

	// 查询多条数据
	queryMultiRowDemo(10)

	// 插入一条数据
	insertRowDemo()

	// 更新一条数据
	updateRowDemo()

	// 删除了一行数据
	deleteRowDemo()
}

// initDB 初始化数据连接
func initDB() (err error) {
	// 1.数据库连接信息
	// 用户名:密码@tcp(IP:端口)/数据库名字
	databaseInfo := `root:root123@tcp(127.0.0.1:3307)/sql_test`

	// 2.连接数据库
	db, err = sqlx.Connect("mysql", databaseInfo)
	if err != nil {
		return err
	}

	// 3.设置数据库连接池
	db.SetMaxOpenConns(16) // 设置数据库连接池的最大连接数
	db.SetMaxIdleConns(4)  // 设置最大空闲连接数

	return nil
}

// ################### 查询 ###############################
type user struct {
	ID   int //注意这里的字段要大写，因为db.Get db.Select等用了反射
	Name string
	Age  int
}

// queryRowDemo 查询单条数据Get
func queryRowDemo(id int) {
	sqlStr := `select id, name, age from user where id = ?`

	var u user
	err := db.Get(&u, sqlStr, id)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Printf("%#v\n", u)
	fmt.Printf("id:%d name:%s age:%d\n", u.ID, u.Name, u.Age)
}

// queryMultiRowDemo 查询多行数据 Select
func queryMultiRowDemo(id int) {
	sqlStr := `select id, name, age from user where id > ?`

	var users []user
	err := db.Select(&users, sqlStr, id) // 这个id是对应上面的 id > ? 中的问号
	if err != nil {
		fmt.Printf("select more rows failed, err:%v\n", err)
		return
	}

	// fmt.Printf("%#v\n", users)
	for _, v := range users {
		fmt.Printf("%#v\n", v)
	}
}

// ###################################### 插入 ################################
// insertRowDemo 插入数据
func insertRowDemo() {
	sqlStr := `insert into user(name,age) values (?,?)`
	ret, err := db.Exec(sqlStr, "沙河娜扎", 99)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}

	theID, err := ret.LastInsertId() // 新插入数据在数据库中的ID
	if err != nil {
		fmt.Printf("get lastinsertid failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert sucess, the id is %d.\n", theID)

	theRow, err := ret.RowsAffected() // 插入了多少行数据
	if err != nil {
		fmt.Printf("get rows failed, err:%v\n", err)
		return
	}
	fmt.Printf("插入了[%d]行数据\n", theRow)
}

// ###################################### 更新 #######################################
// updateRowDemo 更新数据
func updateRowDemo() {
	sqlStr := `update user set age=? where id=?`

	ret, err := db.Exec(sqlStr, 10000, 1)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}

	n, err := ret.RowsAffected() // 操作影响了多少行
	if err != nil {
		fmt.Printf("get rowsaffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("更新成功,影响了%d行数据\n", n)
}

// ###################################### 删除  #########################################
// deleteRowDemo 删除
func deleteRowDemo() {
	sqlStr := `delete from user where id = ?`

	ret, err := db.Exec(sqlStr, 12)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}

	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get rows failed, err:%v\n", err)
		return
	}

	fmt.Printf("删除了%d行数据\n", n)
}

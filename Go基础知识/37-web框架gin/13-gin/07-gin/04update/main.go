package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 删除

// 声明一个全局的db变量
var db *sqlx.DB

// initDB 初始化连接
func initDB() (err error) {
	// 1 数据库连接信息
	// 用户名:密码@tcp(IP:端口)/数据库名称
	databaseInfo := `root:root123@tcp(10.0.0.51:3306)/go`

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

// 修改数据
func updateRowDemo(password string, id int) {
	sqlStr := `update user set password=? where id=?`

	ret, err := db.Exec(sqlStr, password, id)
	if err != nil {
		fmt.Println(err)
		return
	}

	n, err := ret.RowsAffected() // 操作影响了多少行
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("更新数据成功，影响了 [%d] 行数据", n)
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Println(err)
		return
	}

	updateRowDemo("67890", 12)
}

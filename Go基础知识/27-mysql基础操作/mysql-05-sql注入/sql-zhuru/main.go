package main

// [SQL注入问题]
// 我们任何时候都不应该自己拼接SQL语句

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

// initDB 初始化数据连接
func initDB() (err error) {
	databaseInfo := "root:root123@tcp(127.0.0.1:3307)/sql_test"
	db, err = sqlx.Connect("mysql", databaseInfo)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(16)
	db.SetMaxIdleConns(4)

	return err
}

// user 用户结构体
type user struct {
	ID   int
	Name string
	Age  int
}

// SQL注入示例
func sqlInjectDemo(name string) {
	// 自己拼接SQL语句
	// 注意:我们任何时候都不应该自己拼接SQL语句！！！
	sqlStr := fmt.Sprintf("select id, name, age from user where name = '%s' ", name)
	fmt.Printf("SQL:%s\n", sqlStr)

	// 可以使用这种方法防止SQL注入
	// sqlStr := "select id, name, age from user where name = ? "

	var users []user
	err := db.Select(&users, sqlStr)
	if err != nil {
		fmt.Printf("exec SQL failed, err:%v\n", err)
		return
	}

	for _, u := range users {
		fmt.Printf("user:%#v\n", u)
	}
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("connect databae failed, err:%v\n", err)
		return
	}

	// sql注入 传入正常的SQL语句
	sqlInjectDemo("李四") // 客户端正常传入

	// sql注入 传入SQL注入语句
	sqlInjectDemo("xxx' or 1=1 #") // 把数据库整个表查询出来了
	/*
		alnk@Alnk-MacBook-Air sql-zhuru % go  build
		alnk@Alnk-MacBook-Air sql-zhuru % ./sql-zhuru
		SQL:select id, name, age from user where name = '李四'
		user:main.user{ID:3, Name:"李四", Age:22}
		user:main.user{ID:9, Name:"李四", Age:18}
		SQL:select id, name, age from user where name = 'xxx' or 1=1 #'
		user:main.user{ID:1, Name:"Alnk", Age:18}
		user:main.user{ID:3, Name:"李四", Age:22}
		user:main.user{ID:4, Name:"王相机", Age:32}
		user:main.user{ID:5, Name:"天说", Age:72}
		user:main.user{ID:6, Name:"白慧姐", Age:40}
		user:main.user{ID:7, Name:"六七强", Age:30}
		user:main.user{ID:9, Name:"李四", Age:18}
		user:main.user{ID:10, Name:"六七强", Age:30}
		user:main.user{ID:11, Name:"王相机", Age:32}
		user:main.user{ID:56, Name:"沙河娜扎", Age:99}
	*/

	// sql注入 传入SQL注入语句
	sqlInjectDemo("xxx' union select * from user #")
	/*
		alnk@Alnk-MacBook-Air sql-zhuru % go  build
		alnk@Alnk-MacBook-Air sql-zhuru % ./sql-zhuru
		SQL:select id, name, age from user where name = '李四'
		user:main.user{ID:3, Name:"李四", Age:22}
		user:main.user{ID:9, Name:"李四", Age:18}
		SQL:select id, name, age from user where name = 'xxx' or 1=1 #'
		user:main.user{ID:1, Name:"Alnk", Age:18}
		user:main.user{ID:3, Name:"李四", Age:22}
		user:main.user{ID:4, Name:"王相机", Age:32}
		user:main.user{ID:5, Name:"天说", Age:72}
		user:main.user{ID:6, Name:"白慧姐", Age:40}
		user:main.user{ID:7, Name:"六七强", Age:30}
		user:main.user{ID:9, Name:"李四", Age:18}
		user:main.user{ID:10, Name:"六七强", Age:30}
		user:main.user{ID:11, Name:"王相机", Age:32}
		user:main.user{ID:56, Name:"沙河娜扎", Age:99}
	*/
}

/*
[SQL注入问题]
我们任何时候都不应该自己拼接SQL语句！！！
*/

package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//连接mysql
var db *sqlx.DB

func initDB() (err error) {
	databaseInfo := "root:root123@tcp(192.168.3.121:3306)/sql_test"
	db, err = sqlx.Connect("mysql", databaseInfo)
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(5)
	return err
}

//sql注入
type user struct {
	ID   int
	Name string
	Age  int
}

//SQL注入示例
func sqlInjectDemo(name string) {
	//自己拼接SQL语句
	//注意:我们任何时候都不应该自己拼接SQL语句！！！
	sqlStr := fmt.Sprintf("select id, name, age from user where name = '%s' ", name)
	fmt.Printf("SQL:%s\n", sqlStr)

	//可以使用这种方法防止SQL注入，或者使用之前预编译的方法 03mysql_insert_demo
	//sqlStr := "select id, name, age from user where name = ? "

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

	//sql注入
	//sqlInjectDemo("王五") //客户端正常传入

	//sqlInjectDemo("xxx' or 1=1 #") //把数据库整个表查询出来了
	/*
		lichengguo@lichengguodeMacBook-Pro 09sql_injection_demo % ./09sql_injection_demo
		SQL:select id, name, age from user where name = 'xxx' or 1=1 #'
		user:main.user{ID:1, Name:"张三", Age:10000}
		user:main.user{ID:2, Name:"李四", Age:20}
		user:main.user{ID:3, Name:"王五", Age:90}
		user:main.user{ID:4, Name:"马六", Age:37}
		user:main.user{ID:5, Name:"马云", Age:58}
		user:main.user{ID:6, Name:"王相机", Age:32}
		user:main.user{ID:7, Name:"天说", Age:9000}
		user:main.user{ID:10, Name:"沙河小王子", Age:19}
		user:main.user{ID:11, Name:"沙河娜扎", Age:32}
	*/

	//sqlInjectDemo("xxx' union select * from user #")
	/*
		lichengguo@lichengguodeMacBook-Pro 09sql_injection_demo % ./09sql_injection_demo
		SQL:select id, name, age from user where name = 'xxx' union select * from user #'
		user:main.user{ID:1, Name:"张三", Age:10000}
		user:main.user{ID:2, Name:"李四", Age:20}
		user:main.user{ID:3, Name:"王五", Age:90}
		user:main.user{ID:4, Name:"马六", Age:37}
		user:main.user{ID:5, Name:"马云", Age:58}
		user:main.user{ID:6, Name:"王相机", Age:32}
		user:main.user{ID:7, Name:"天说", Age:9000}
		user:main.user{ID:10, Name:"沙河小王子", Age:19}
		user:main.user{ID:11, Name:"沙河娜扎", Age:32}
	*/
}

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

// 声明一个结构体，用来接收查询出来的数据
type user struct {
	id   int
	name string
	age  int
}

// queryOne 查看一行数据
func queryOne(id int) {
	var u1 user
	//1.写查询单条记录的sql语句
	sqlStr := `select id, name, age from user where id = ?`
	//2.执行并拿到结果
	//必须对db.QueryRow()调用Scan方法，因为该方法会释放数据库连接
	//Scan()从连接池里拿出一个连接出来去数据库查询数据
	db.QueryRow(sqlStr, id).Scan(&u1.id, &u1.name, &u1.age)
	//3.打印结果
	fmt.Printf("u1: %#v\n", u1)
}

// queryMore 查询多条数据
func queryMore(n int) (userInfo []user) {
	//1.sql语句
	sqlSstr := `select id, name, age from user where id > ?`
	//2.执行
	rows, err := db.Query(sqlSstr, n)
	if err != nil {
		fmt.Printf("exec %s query falied, err:%v\n", sqlSstr, err)
		return
	}
	//3.一定要关闭rows，节省数据库资源
	defer rows.Close()
	//4.循环读取数据
	for rows.Next() {
		var u1 user
		err := rows.Scan(&u1.id, &u1.name, &u1.age)
		if err != nil {
			fmt.Printf("Scan failed, err:%v\n", err)
			return
		}
		userInfo = append(userInfo, u1)
	}
	return userInfo
}

func main() {
	//初始化数据库连接
	err := initDB()
	if err != nil {
		fmt.Printf("init DB falied, err:%v\n", err)
		return
	}
	fmt.Println("connect database success!")

	//查1行
	queryOne(1)

	//查多行
	ret := queryMore(1)
	fmt.Printf("%#v\n", ret)
	fmt.Println(ret)
}

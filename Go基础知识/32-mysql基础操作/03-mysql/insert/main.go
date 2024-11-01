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

	// 执行插入函数
	//insertNo1()
	//insertNo2("李四", 18)
	//insertNo2("王五", 90)
	//insertNo2("马六", 37)
	//insertNo2("马云", 58)
	//insertNo3()
}

//insertNo1 插入数据方法一
func insertNo1() {
	// 1.写sql语句
	sqlStr := `insert into user(name, age) values("张三", 20)`
	// 2.执行exec
	ret, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	// 拿到插入的数据的在数据库表中的ID
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get id failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert ID:%d\n", id)
	// 拿到数据库中受影响的数据行数
	row, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Println("受影响的行数:", row)
}

//insertNo2 插入数据方法二
func insertNo2(name string, age int) {
	sqlStr := `insert into user(name, age) values(?, ?)`
	ret, err := db.Exec(sqlStr, name, age)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	// 拿到插入的数据的在数据库表中的ID
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get id failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert ID:%d\n", id)
	// 拿到数据库中受影响的数据行数
	row, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Println("受影响的行数:", row)
}

/*
预处理插入多条数据 这种方式速度会快一点

普通SQL语句执行过程：
	客户端对SQL语句进行占位符替换得到完整的SQL语句
	客户端发送完整SQL语句到MySQL服务端
	MySQL服务端执行完整的SQL语句并将结果返回给客户端

预处理执行过程：
	把SQL语句分成两部分，命令部分与数据部分
	先把命令部分发送给MySQL服务端，MySQL服务端进行SQL预处理
	然后把数据部分发送给MySQL服务端，MySQL服务端对SQL语句进行占位符替换
	MySQL服务端执行完整的SQL语句并将结果返回给客户端

为什么要预处理？
	优化MySQL服务器重复执行SQL的方法，可以提升服务器性能，提前让服务器编译，一次编译多次执行，节省后续编译的成本。
	避免SQL注入问题。
*/
func insertNo3() {
	// 1.拼写SQL语句
	sqlStr := `insert into user(name, age) values(?, ?)`
	// 2.预处理SQL语句,发送到MySQL服务器
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	// 3.关闭连接，释放连接池的资源
	defer stmt.Close()
	// 4.SQL数据
	var m = map[string]int{
		"六七强": 30,
		"王相机": 32,
		"天说":  72,
		"白慧姐": 40,
	}
	// 5.发送数据到MySQL服务器
	for k, v := range m {
		stmt.Exec(k, v)
	}
}

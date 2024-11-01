package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 创建基础数据
// CREATE TABLE `user` (
//  `id` int(50) NOT NULL AUTO_INCREMENT,
//  `username` varchar(50) DEFAULT NULL,
//  `password` varchar(50) DEFAULT NULL,
//  PRIMARY KEY (`id`)
//) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;
//
// insert  into `user`(`id`,`username`,`password`) values (1,'zs','123'),(2,'li','456'),(3,'ww','789'),(4,'zz','135'),(5,'tt','246');

// gin框架用的数据库连接也是 sqlx
// 可以参考day10的 go使用sqlx包管理mysql数据库

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

// 查询数据
type user struct {
	ID       int // 这里字段名要大写，因为db.GEt db.Select 用了反射
	Username string
	Password string
}

// 查询单条数据
func queryRowDemo(id int) {
	sqlStr := `select id, username, password from user where id = ?`
	var u user
	err := db.Get(&u, sqlStr, id)
	if err != nil {
		fmt.Println("get failed, ", err)
		return
	}
	fmt.Printf("%#v\n", u)
	fmt.Println(u.ID, u.Username, u.Password)
}

// 查询多条数据
func queryMultiRowDemo(id int) {
	sqlStr := `select id, username, password from user where id > ?`

	var users []user
	err := db.Select(&users, sqlStr, id)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range users {
		fmt.Printf("%#v\n", v)
	}
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	//queryRowDemo(1)
	queryMultiRowDemo(0)

}

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //这里只要导入即可，使用的是这个包的 init() 函数
)

/*
[Go语言操作MySQL]
Go语言中内置的database/sql包提供了保证SQL或类SQL数据库的泛用接口，并不提供具体的数据库驱动。
使用database/sql包时必须注入（至少）一个数据库驱动

下载数据库驱动
go get -u github.com/go-sql-driver/mysql
go get github.com/go-sql-driver/mysql
加-u和不加的区别
加上它可以利用网络来更新已有的代码包及其依赖包。
如果已经下载过一个代码包，但是这个代码包又有更新了，那么这时候可以直接用 -u 标记来更新本地的对应的代码包。
如果不加这个 -u 标记，执行 go get 一个已有的代码包，会发现命令什么都不执行。
只有加了 -u 标记，命令会去执行 git pull 命令拉取最新的代码包的最新版本，下载并安装
*/

func main() {
	// 1.拼接连接数据库的信息
	// 用户名:密码@tcp(ip:端口)/数据库的名字
	databaseInfo := "root:root123@tcp(192.168.3.121:3306)/sql_test"

	// 2.校验拼接信息格式是否正确，不会去校验数据库的账号密码是否正确
	db, err := sql.Open("mysql", databaseInfo)
	if err != nil {
		fmt.Printf("databaseInfo:%s invalid, err:%v\n", databaseInfo, err)
		return
	}

	// 3.连接数据库
	err = db.Ping() //尝试连接数据库
	if err != nil {
		fmt.Printf("Open %s failed, err:%v\n", databaseInfo, err)
		return
	}
	fmt.Println("connect database sucess!")
}

# 数据库

## MySQL

### 连接数据库

```go
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
```

```go
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 数据库连接对象可以在多个函数中使用

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
}
```



### 插入数据

```go
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
	//1.写sql语句
	sqlStr := `insert into user(name, age) values("张三", 20)`
	//2.执行exec
	ret, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	//拿到插入的数据的在数据库表中的ID
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get id failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert ID:%d\n", id)
	//拿到数据库中受影响的数据行数
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
	//拿到插入的数据的在数据库表中的ID
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get id failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert ID:%d\n", id)
	//拿到数据库中受影响的数据行数
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
	//1.拼写SQL语句
	sqlStr := `insert into user(name, age) values(?, ?)`
	//2.预处理SQL语句,发送到MySQL服务器
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	//3.关闭连接，释放连接池的资源
	defer stmt.Close()
	//4.SQL数据
	var m = map[string]int{
		"六七强": 30,
		"王相机": 32,
		"天说":  72,
		"白慧姐": 40,
	}
	//5.发送数据到MySQL服务器
	for k, v := range m {
		stmt.Exec(k, v)
	}
}
```



### 删除数据

```go
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//在全局中声明一个db变量，好让所有的函数都能调用
var db *sql.DB

//initDB 初始化连接
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

//deleteRow 删除一行数据
func deleteRow(id int) {
	sqlStr := `delete from user where id=?`
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete id:%d failed, err:%v\n", id, err)
		return
	}
	//拿到被删除的行数
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get delete id failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete %d row data sucess!\n", n)
}

func main() {
	//初始化数据库连接
	err := initDB()
	if err != nil {
		fmt.Printf("init DB falied, err:%v\n", err)
		return
	}
	fmt.Println("connect database success!")

	//删除一行
	deleteRow(15)
}
```



### 更新数据

```go
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

// updateRow 修改一行数据
func updateRow(newAge int, id int) {
	sqlStr := `update user set age=? where id = ?`
	ret, err := db.Exec(sqlStr, newAge, id)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() //获取受影响的行数
	if err != nil {
		fmt.Printf("get any rows failed, err:%v\n", err)
	}
	fmt.Printf("update %d rows data!\n", n)
}

func main() {
	//初始化数据库连接
	err := initDB()
	if err != nil {
		fmt.Printf("init DB falied, err:%v\n", err)
		return
	}
	fmt.Println("connect database success!")

	//修改一行
	updateRow(9000, 1)
}
```



### 查询数据

```go
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
```



### 事务

```go
/*
Go实现MySQL事务
事务：一个最小的不可再分的工作单元
通常一个事务对应一个完整的业务(例如银行账户转账业务，该业务就是一个最小的工作单元)，
同时这个完整的业务需要执行多次的DML(insert、update、delete)语句共同联合完成。
A转账给B，这里面就需要执行两次update操作

在MySQL中只有使用了Innodb数据库引擎的数据库或表才支持事务。
事务处理可以用来维护数据库的完整性，保证成批的SQL语句要么全部执行，要么全部不执行

通常事务必须满足4个条件（ACID）：
原子性（Atomicity，或称不可分割性）、
一致性（Consistency）、
隔离性（Isolation，又称独立性）、
持久性（Durability
*/

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// initDB 连接数据库
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

// transactionDemo 事务提交
func transactionDemo() {
	// 1.开启事务
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("begin failed, err:%v\n", err)
		return
	}

	//2.执行多个sql操作
	sqlStr1 := `update user set age=age-2 where id = 1`
	sqlStr2 := `update user set age=age+2 where id = 2`

	//3.执行sqlStr1
	ret, err := tx.Exec(sqlStr1)
	if err != nil {
		//如果执行失败就回滚
		tx.Rollback()
		fmt.Println("执行SQL1出错了，已经回滚")
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		//如果执行失败就回滚
		tx.Rollback()
		fmt.Println("get RowsAffected failed, err:", err)
		return
	}
	fmt.Printf("影响了%d行数据", n)

	//4.执行sqlStr2
	_, err = tx.Exec(sqlStr2)
	if err != nil {
		//如果执行失败就回滚
		tx.Rollback()
		fmt.Println("执行sqlStr2出错了，已经回滚")
		return
	}

	//5.上面两步sql都执行成功，那么就去提交
	err = tx.Commit()
	if err != nil {
		//提交失败 也要回滚
		tx.Rollback()
		fmt.Println("提交失败，回顾完成")
		return
	}
	fmt.Println("事务执行成功!")
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("connect database failed, err:%v\n", err)
		return
	}

	//调用事务执行函数
	transactionDemo()
}
```



### sqlx库

```go
/*
[sqlx库使用指南]

使用sqlx实现批量插入数据的例子，介绍了sqlx中可能被忽视了的 sqlx.In 和 DB.NamedExec 方法

sqlx可以认为是Go语言内置 database/sql 的超集
它在优秀的内置 database/sql 基础上提供了一组扩展
这些扩展中除了大家常用来查询的 Get 和 Select 外还有很多其他强大的功能

sqlx中的exec方法与原生sql中的exec使用基本一致，所以增删改应该跟原生的sql库差不多

安装方法:
	go get github.com/jmoiron/sqlx
*/

package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

//initDB 初始化数据连接
func initDB() (err error) {
	//1.数据库连接信息
	//用户名:密码@tcp(IP:端口)/数据库名字
	databaseInfo := `root:root123@tcp(192.168.3.121:3306)/sql_test`

	//2.连接数据库
	db, err = sqlx.Connect("mysql", databaseInfo)
	if err != nil {
		return
	}

	//3.设置数据库连接池
	db.SetMaxOpenConns(10) //设置数据库连接池的最大连接数
	db.SetMaxIdleConns(5)  //设置最大空闲连接数
	return
}

//##########################################################################################################
//###################################### 查询 start #########################################################
//##########################################################################################################
type user struct {
	ID   int //注意这里的字段要大写，因为db.Get db.Select等用了反射
	Name string
	Age  int
}

//queryRowDemo 查询单条数据Get
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

//queryMultiRowDemo 查询多行数据 Select
func queryMultiRowDemo(id int) {
	sqlStr := `select id, name, age from user where id > ?`

	var users []user
	err := db.Select(&users, sqlStr, id) //这个id是对应上面的 id > ? 中的问号
	if err != nil {
		fmt.Printf("select more rows failed, err:%v\n", err)
		return
	}

	//fmt.Printf("%#v\n", users)
	for _, v := range users {
		fmt.Printf("%#v\n", v)
	}
}

//##########################################################################################################
//###################################### 查询 end ###########################################################
//##########################################################################################################
//
//
//##########################################################################################################
//###################################### 插入 start #########################################################
//##########################################################################################################
//insertRowDemo 插入数据
func insertRowDemo() {
	sqlStr := `insert into user(name,age) values (?,?)`
	ret, err := db.Exec(sqlStr, "沙河娜扎", 99)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() //新插入数据在数据库中的ID
	if err != nil {
		fmt.Printf("get lastinsertid failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert sucess, the id is %d.\n", theID)

	theRow, err := ret.RowsAffected() //插入了多少行数据
	if err != nil {
		fmt.Printf("get rows failed, err:%v\n", err)
		return
	}
	fmt.Printf("插入了[%d]行数据\n", theRow)
}

//##########################################################################################################
//###################################### 插入 end #########################################################
//##########################################################################################################
//
//
//##########################################################################################################
//###################################### 更新 start #########################################################
//##########################################################################################################
//updateRowDemo 更新数据
func updateRowDemo() {
	sqlStr := `update user set age=? where id=?`

	ret, err := db.Exec(sqlStr, 10000, 1)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}

	n, err := ret.RowsAffected() //操作影响了多少行
	if err != nil {
		fmt.Printf("get rowsaffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("更新成功,影响了%d行数据\n", n)

}

//##########################################################################################################
//###################################### 更新 end ###########################################################
//##########################################################################################################
//
//
//##########################################################################################################
//###################################### 删除 start ########################################################
//##########################################################################################################
//deleteRowDemo 删除
func deleteRowDemo() {
	sqlStr := `delete from user where id = ?`

	ret, err := db.Exec(sqlStr, 12)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}

	n, err := ret.RowsAffected() //操作影响的行数
	if err != nil {
		fmt.Printf("get rows failed, err:%v\n", err)
		return
	}

	fmt.Printf("删除了%d行数据\n", n)
}

//##########################################################################################################
//###################################### 删除 end ###########################################################
//##########################################################################################################
//
//
func main() {
	//连接数据库初始化
	err := initDB()
	if err != nil {
		fmt.Printf("connect database failed, err:%v\n", err)
		return
	}

	//查询一条数据
	//queryRowDemo(2)

	//查询多条数据
	//queryMultiRowDemo(10)

	//插入一条数据
	//insertRowDemo()

	//更新一条数据
	//updateRowDemo()

	//删除了一行数据
	//deleteRowDemo()
}
```



### sql注入

```go
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
```



## Redis

```go
/*
安装: go get  github.com/go-redis/redis
*/

package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

//在全局中声明一个db变量，好让所有的函数都能调用
var rdb *redis.Client

//initClient 普通连接
func initClient() (err error) {
	//1.配置信息
	rdb = redis.NewClient(&redis.Options{
		Addr: "192.168.3.121:6379",
		DB:   0,
	})
	//2.尝试连接
	str, err := rdb.Ping().Result()
	if err != nil {
		return err
	}
	fmt.Println("str:", str) //str: PONG
	return
}

/*
连接Redis哨兵模式
var rdb *redis.Client

func initClient()(err error){
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "master",
		SentinelAddrs: []string{"x.x.x.x:26379", "xx.xx.xx.xx:26379", "xxx.xxx.xxx.xxx:26379"},
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
*/

/*
连接Redis集群
var rdb *redis.Client

func initClient()(err error){
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

*/

//set/get示例 redisExample
func redisExample() {
	//1.插入值
	err := rdb.Set("score", 100, 0).Err() //set(key,values,失效时间 0表示永不失效)
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
		return
	}
	//2.取值
	//可以取到值的情况
	val, err := rdb.Get("score").Result()
	if err != nil {
		fmt.Printf("get score value failed, err:%v\n", err)
		return
	}
	fmt.Println("keys:score , val:", val)

	//键值对不存在或者值为nil的情况
	val2, err := rdb.Get("name").Result()
	if err == redis.Nil {
		fmt.Println("name does not exist")
	} else if err != nil {
		fmt.Printf("get name failed, err:%v\n", err)
		return
	} else {
		fmt.Println("keys: name val: ", val2)
	}
}

//zset示例
func redisExample2() {
	zsetKey := "language_rank"
	languages := []redis.Z{
		redis.Z{Score: 90.0, Member: "Golang"}, //Socre:相当于值 Member:相当于键
		redis.Z{Score: 98.0, Member: "Java"},
		redis.Z{Score: 95.0, Member: "Python"},
		redis.Z{Score: 97.0, Member: "JavaScript"},
		redis.Z{Score: 99.0, Member: "C/C++"},
	}

	//zadd
	num, err := rdb.ZAdd(zsetKey, languages...).Result()
	if err != nil {
		fmt.Printf("zadd failed, err:%v\n", err)
		return
	}
	fmt.Printf("zadd %d success!\n", num)

	//把Golang的分数加10
	newScore, err := rdb.ZIncrBy(zsetKey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed, err:%v\n", err)
		return
	}
	fmt.Printf("Golang's score is %f now.\n", newScore)

	//取分数最高的3个
	ret, err := rdb.ZRevRangeWithScores(zsetKey, 0, 2).Result()
	if err != nil {
		return
	}
	for _, z := range ret {
		fmt.Println(z)
	}

	//取95-100分的
	op := redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = rdb.ZRangeByScoreWithScores(zsetKey, op).Result()
	if err != nil {
		fmt.Printf("zrangebyscore failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}

//inserRedis1W 测试redis
func inserRedis1W() {
	startTime := time.Now().Unix()
	for i := 0; i < 10000; i++ {
		keyName := fmt.Sprintf("name%00000d", i)
		err := rdb.Set(keyName, "我是很占内存的一串字符串!!!!!!", time.Second*60).Err() //设置了key过期的时间
		if err != nil {
			fmt.Printf("insert failed, err:%v\n", err)
			return
		}
	}
	fmt.Printf("耗时:%v\n", time.Now().Unix()-startTime) //耗时:41
}

/*
pipLine 测试
Pipeline 主要是一种网络优化。它本质上意味着客户端缓冲一堆命令并一次性将它们发送到服务器。
这些命令不能保证在事务中执行。这样做的好处是节省了每个命令的网络往返时间（RTT）
*/
func pipLine() {
	startTime := time.Now().Unix()

	pipe := rdb.Pipeline()

	for i := 0; i < 10000; i++ {
		keyName := fmt.Sprintf("name%00000d", i)
		pipe.Set(keyName, "我是很占内存的一串字符串!!!!!!", time.Second*60)
	}

	pipe.Exec()

	fmt.Printf("耗时:%v\n", time.Now().Unix()-startTime) //耗时:0
}

func main() {
	err := initClient()
	if err != nil {
		fmt.Printf("connect redis server failed, err:%v\n", err)
		return
	}
	fmt.Println("redis connect sucess!")

	//set/get示例
	redisExample()

	//zadd
	//redisExample2()

	//测试插入1万键值对
	//inserRedis1W()

	//pipLine 测试
	//pipLine()

}

/*
192.168.3.121:6379> zrange language_rank 0 100
1) "Golang"
2) "Python"
3) "JavaScript"
4) "Java"
5) "C/C++"

192.168.3.121:6379> zrange language_rank 0 100 withscores
 1) "Golang"
 2) "90"
 3) "Python"
 4) "95"
 5) "JavaScript"
 6) "97"
 7) "Java"
 8) "98"
 9) "C/C++"
10) "99"
*/
```



## 消息队列NSQ

简介

```go
github下载地址: https://github.com/nsqio/nsq/releases/tag/v1.2.0

简介
NSQ是目前比较流行的一个分布式的消息队列，本文主要介绍了NSQ及Go语言如何操作NSQ。

NSQ是Go语言编写的一个开源的实时分布式内存消息队列，其性能十分优异。 NSQ的优势：
    NSQ提倡分布式和分散的拓扑，没有单点故障，支持容错和高可用性，并提供可靠的消息交付保证
    NSQ支持横向扩展，没有任何集中式代理。
    NSQ易于配置和部署，并且内置了管理界面
    
NSQ组件
1.nsqd: nsqd是一个守护进程，它接收、排队并向客户端发送消息
启动nsqd，指定-broadcast-address=127.0.0.1来配置广播地址
./nsqd -broadcast-address=127.0.0.1

如果是在搭配nsqlookupd使用的模式下需要还指定nsqlookupd地址:
./nsqd -broadcast-address=127.0.0.1 -lookupd-tcp-address=127.0.0.1:4160
如果是部署了多个nsqlookupd节点的集群，那还可以指定多个-lookupd-tcp-address

2.nsqlookupd
nsqlookupd是维护所有nsqd状态、提供服务发现的守护进程。它能为消费者查找特定topic下的nsqd提供了运行时的自动发现服务。 
它不维持持久状态，也不需要与任何其他nsqlookupd实例协调以满足查询。
因此根据你系统的冗余要求尽可能多地部署nsqlookupd节点。
它们小豪的资源很少，可以与其他服务共存。我们的建议是为每个数据中心运行至少3个集群

3.nsqadmin
一个实时监控集群状态、执行各种管理任务的Web管理平台。 启动nsqadmin，指定nsqlookupd地址
./nsqadmin -lookupd-http-address=127.0.0.1:4161
我们可以使用浏览器打开http://127.0.0.1:4171/访问如下管理界面
```



目录结构

```go
├── README.md
├── consumer
│   └── main.go
└── producer
    └── main.go
```



consumer/main.go 文件

```go
// nsq_consumer/main.go
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nsqio/go-nsq"
)

// NSQ Consumer Demo

// MyHandler 是一个消费者类型
type MyHandler struct {
	Title string
}

// HandleMessage 是需要实现的处理消息的方法
func (m *MyHandler) HandleMessage(msg *nsq.Message) (err error) {
	fmt.Printf("%s recv from %v, msg:%v\n", m.Title, msg.NSQDAddress, string(msg.Body))
	return
}

// 初始化消费者
func initConsumer(topic string, channel string, address string) (err error) {
	config := nsq.NewConfig()
	config.LookupdPollInterval = 15 * time.Second
	c, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		fmt.Printf("create consumer failed, err:%v\n", err)
		return
	}
	consumer := &MyHandler{
		Title: "dsb1号",
	}
	c.AddHandler(consumer)

	//if err := c.ConnectToNSQD(address); err != nil { // 直接连NSQD
	if err := c.ConnectToNSQLookupd(address); err != nil { // 通过lookupd查询
		return err
	}
	return nil

}

func main() {
	err := initConsumer("topic_demo", "first", "127.0.0.1:4161")
	if err != nil {
		fmt.Printf("init consumer failed, err:%v\n", err)
		return
	}
	c := make(chan os.Signal)        // 定义一个信号的通道
	signal.Notify(c, syscall.SIGINT) // 转发键盘中断信号到c
	<-c                              // 阻塞
}
```



producer/main.go文件

```go
// nsq_producer/main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/nsqio/go-nsq"
)

// NSQ Producer Demo 生产者

var producer *nsq.Producer

// 初始化生产者
func initProducer(str string) (err error) {
	config := nsq.NewConfig()
	producer, err = nsq.NewProducer(str, config)
	if err != nil {
		fmt.Printf("create producer failed, err:%v\n", err)
		return err
	}
	return nil
}

func main() {
	nsqAddress := "127.0.0.1:4150"
	err := initProducer(nsqAddress)
	if err != nil {
		fmt.Printf("init producer failed, err:%v\n", err)
		return
	}

	reader := bufio.NewReader(os.Stdin) // 从标准输入读取
	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("read string from stdin failed, err:%v\n", err)
			continue
		}
		data = strings.TrimSpace(data)
		if strings.ToUpper(data) == "Q" { // 输入Q退出
			break
		}
		// 向 'topic_demo' publish 数据
		err = producer.Publish("topic_demo", []byte(data))
		if err != nil {
			fmt.Printf("publish msg to nsq failed, err:%v\n", err)
			continue
		}
	}
}
```

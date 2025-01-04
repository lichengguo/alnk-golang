package main

// 数据库操作[增删改查]、事务操作

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB // 在全局中声明一个db变量，好让所有的函数都能调用

func main() {
	// 初始化数据库连接
	err := initDB()
	if err != nil {
		panic(err)
	}

	// 执行插入函数
	insertNo1()
	insertNo2("李四", 18)
	insertNo3()

	// 执行删除函数
	deleteRow("张三")

	// 执行修改函数
	updateRow(16, 1)

	// 执行查询函数
	queryOne(1)
	queryMore(0)

	// 执行事务函数
	transactionDemo()
}

// ================== 数据库连接 ==================
// initDB 初始化数据库连接
func initDB() (err error) {
	// 1.数据库信息
	databaseInfo := "root:root123@tcp(127.0.0.1:3307)/sql_test"

	// 2.校验
	db, err = sql.Open("mysql", databaseInfo)
	if err != nil {
		return err
	}

	// 3.连接
	err = db.Ping()
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(16) // 设置数据库连接池的最大连接数
	db.SetMaxIdleConns(4)  // 设置空闲时间最大连接数

	return nil
}

// ================== 数据库插入操作 ==================
// insertNo1 插入数据方法一
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

// insertNo2 插入数据方法二
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

// insertNo3 插入数据方法三
func insertNo3() {
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
			优化MySQL服务器重复执行SQL的方法，可以提升服务器性能，提前让服务器编译，一次编译多次执行，节省后续编译的成本
			避免SQL注入问题
	*/

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

// ================== 数据库删除操作 ==================
// deleteRow 删除一行数据
func deleteRow(name string) {
	sqlStr := `delete from user where name=?`
	ret, err := db.Exec(sqlStr, name)
	if err != nil {
		fmt.Printf("delete name:%s failed, err:%v\n", name, err)
		return
	}

	// 拿到被删除的行数
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get delete id failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete %d row data sucess!\n", n)
}

// ================== 数据库修改操作 ==================
// updateRow 修改数据
func updateRow(newAge int, id int) {
	sqlStr := `update user set age=? where id = ?`
	ret, err := db.Exec(sqlStr, newAge, id)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}

	n, err := ret.RowsAffected() // 获取受影响的行数
	if err != nil {
		fmt.Printf("get any rows failed, err:%v\n", err)
	}
	fmt.Printf("update %d rows data!\n", n)
}

// ================== 数据库查询操作 ==================
// 声明一个结构体，用来接收查询出来的数据
type user struct {
	id   int
	name string
	age  int
}

// queryOne 查看一行数据
func queryOne(id int) {
	var u1 user
	// 1.写查询单条记录的sql语句
	sqlStr := `select id, name, age from user where id = ?`

	// 2.执行并拿到结果
	// 必须对db.QueryRow()调用Scan方法，因为该方法会释放数据库连接
	// Scan()从连接池里拿出一个连接出来去数据库查询数据
	db.QueryRow(sqlStr, id).Scan(&u1.id, &u1.name, &u1.age)

	// 3.打印结果
	fmt.Printf("u1: %#v\n", u1)
}

// queryMore 查询多条数据
func queryMore(n int) {
	// 1.sql语句
	sqlSstr := `select id, name, age from user where id > ?`

	// 2.执行
	rows, err := db.Query(sqlSstr, n)
	if err != nil {
		fmt.Printf("exec %s query falied, err:%v\n", sqlSstr, err)
		return
	}

	// 3.一定要关闭rows，节省数据库资源
	defer rows.Close()

	// 4.循环读取数据
	var userInfo []user
	for rows.Next() {
		var u1 user
		err := rows.Scan(&u1.id, &u1.name, &u1.age)
		if err != nil {
			fmt.Printf("Scan failed, err:%v\n", err)
			return
		}
		userInfo = append(userInfo, u1)
	}
	fmt.Printf("%#v\n", userInfo)
}

// ================== 事务操作 ==================
// transactionDemo 事务提交
func transactionDemo() {
	/*
	   Go实现MySQL事务
	   事务：一个最小的不可再分的工作单元
	   通常一个事务对应一个完整的业务(例如银行账户转账业务，该业务就是一个最小的工作单元)
	   同时这个完整的业务需要执行多次的DML(insert、update、delete)语句共同联合完成
	   A转账给B，这里面就需要执行两次update操作

	   在MySQL中只有使用了Innodb数据库引擎的数据库或表才支持事务
	   事务处理可以用来维护数据库的完整性，保证成批的SQL语句要么全部执行，要么全部不执行

	   通常事务必须满足4个条件（ACID）
	   原子性（Atomicity，或称不可分割性）
	   一致性（Consistency）
	   隔离性（Isolation，又称独立性）
	   持久性（Durability
	*/

	// 1.开启事务
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("begin failed, err:%v\n", err)
		return
	}

	// 2.执行多个sql操作
	sqlStr1 := `update user set age=age-2 where id = 1`
	sqlStr2 := `update user set age=age+2 where id = 3`

	// 3.执行sqlStr1
	ret, err := tx.Exec(sqlStr1)
	if err != nil {
		// 如果执行失败就回滚
		tx.Rollback()
		fmt.Println("执行SQL1出错了, 已经回滚")
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		// 如果执行失败就回滚
		tx.Rollback()
		fmt.Println("get RowsAffected failed, err:", err)
		return
	}
	fmt.Printf("影响了%d行数据\n", n)

	// 4.执行sqlStr2
	ret, err = tx.Exec(sqlStr2)
	if err != nil {
		// 如果执行失败就回滚
		tx.Rollback()
		fmt.Println("执行sqlStr2出错了, 已经回滚")
		return
	}
	n, err = ret.RowsAffected()
	if err != nil {
		// 如果执行失败就回滚
		tx.Rollback()
		fmt.Println("get RowsAffected failed, err:", err)
		return
	}
	fmt.Printf("影响了%d行数据\n", n)

	// 5.上面两步sql都执行成功，那么就去提交
	err = tx.Commit()
	if err != nil {
		// 提交失败 也要回滚
		tx.Rollback()
		fmt.Println("提交失败，回顾完成")
		return
	}

	fmt.Println("事务执行成功!")
}

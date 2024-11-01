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

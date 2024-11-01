package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 图书管理系统

func main() {
	// 初始化数据库
	err := InitDB()
	if err != nil {
		fmt.Println("init database failed, ", err)
		return
	}

	// 创建路由
	r := gin.Default()

	// 加载页面模板
	//r.LoadHTMLGlob("/Users/lichengguo/go/src/code.oldboyedu.com/gostudy/day14gin框架/08gin数据库/05数据库练习/templates/*")
	r.LoadHTMLGlob("./templates/*")

	// 绑定路由
	book := r.Group("/book")
	{
		book.GET("/list", bookList)      // 查询所有图书
		book.GET("/new", getAddBookHtml) // 添加图书
		book.POST("/new", addBook)       // 添加图书
		book.GET("/delete", deleteBook)  // 删除图书
	}

	// 启动程序
	r.Run(":8080")
}

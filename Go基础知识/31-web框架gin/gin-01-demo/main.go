package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// gin框架安装: go get -u github.com/gin-gonic/gin
// Gin是一个golang的微框架，封装比较优雅，API友好，源码注释比较明确，具有快速灵活，容错方便等特点
// 对于golang而言，web框架的依赖要远比Python，Java之类的要小。自身的net/http足够简单，性能也非常不错
// 借助框架开发，不仅可以省去很多常用的封装带来的时间，也有助于团队的编码风格和形成规范

func main() {
	// 禁用控制台颜色
	gin.DisableConsoleColor()

	// 日志记录到文件
	f, _ := os.OpenFile("gin.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	gin.DefaultWriter = io.MultiWriter(f)
	defer f.Close() // 程序退出关闭文件

	// 如果需要同时将日志写入文件和控制台，请使用以下代码
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// 1 创建路由
	r := gin.Default()

	// 2 绑定路由规则， 执行的函数
	// gin.Context 封装了 request 和 response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world!")
	})

	// 3 监听端口，默认在8080
	r.Run(":8080")
}

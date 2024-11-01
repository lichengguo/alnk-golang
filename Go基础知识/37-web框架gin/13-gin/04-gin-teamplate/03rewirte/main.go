package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 重定向

func main() {
	// 1 创建路由
	r := gin.Default()

	// 2 绑定路由
	// 2.1 外部重定向
	r.GET("/redirect1", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})

	// 2.2 内部重定向
	r.GET("/redirect2", func(c *gin.Context) {
		c.Request.URL.Path = "/redirect1"
		r.HandleContext(c)
	})

	// 3 启动
	r.Run(":8080")
}

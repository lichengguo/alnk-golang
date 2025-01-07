package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// gin 路由

func main() {
	// 1 创建路由
	// 默认使用了2个中间件 Logger(), Recovery()
	r := gin.Default()
	// 也可以创建不带中间件的路由
	// r := gin.New()

	// 2 绑定路由规则和执行的函数
	// gin.Context，封装了request和response
	// 2.1 GET请求
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world! GET!")
	})

	// 2.2 POST请求
	r.POST("/", postURL)

	// 2.3 PUT请求
	r.PUT("/", putURL)

	// 2.4 DELETE请求
	r.DELETE("/", deleteURL)

	// 3 监听端口
	r.Run(":8080")
}

// postURL POST请求
func postURL(c *gin.Context) {
	c.String(http.StatusOK, "POST is ok!")
}

// putURL PUT请求
func putURL(c *gin.Context) {
	c.String(http.StatusOK, "PUT is ok!")
}

// deleteURL DELETE请求
func deleteURL(c *gin.Context) {
	c.String(http.StatusOK, "DELETE is ok!")
}

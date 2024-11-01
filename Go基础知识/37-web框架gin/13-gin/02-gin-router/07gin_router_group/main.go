package main

import "github.com/gin-gonic/gin"

// routes group 是为了管理一些相同的URL

func main() {
	// 1 创建路由
	r := gin.Default()

	// 2 路由组，处理GET请求的路由组
	v1 := r.Group("/v1")
	// {} 是gin框架的书写规范
	{
		v1.GET("/index", index)
		v1.GET("/login", login)
	}

	// 3 监听
	r.Run(":8080")
}

func index(c *gin.Context) {
	c.String(200, "index is ok!")
}

func login(c *gin.Context) {
	c.String(200, "login is ok!")
}

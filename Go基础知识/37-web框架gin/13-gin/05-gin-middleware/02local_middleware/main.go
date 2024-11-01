package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 局部中间件
// 根路由后面定义的中间件

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("我是中间件...")
		c.Next()
		fmt.Println("中间件执行完毕...")
	}
}

func main() {
	r := gin.Default()

	// 局部中间件
	r.GET("/middleware", MiddleWare(), func(c *gin.Context) {
		fmt.Println("我是根路由中的函数")

		// 返回给前端的数据
		c.JSON(200, gin.H{"msg": "中间件"})
	})

	r.Run(":8080")
}

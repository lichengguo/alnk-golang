package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// 中间件练习
// 需求 定义程序计时中间件，然后定义2个路由，执行函数后应该打印统计的执行时间

// 定义中间件
func myTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()
		// 统计时间
		fmt.Println("程序耗时:", time.Since(t))
	}
}

func main() {
	// 创建路由
	r := gin.Default()

	// 注册中间件
	r.Use(myTime())

	// 绑定路由
	shoppingGroup := r.Group("/shopping")
	{
		shoppingGroup.GET("/index", shopIndexHandler)
		shoppingGroup.GET("/home", shopHomeHandler)
	}

	// 启动程序
	r.Run(":8080")
}

func shopIndexHandler(c *gin.Context) {
	time.Sleep(5 * time.Second)
}

func shopHomeHandler(c *gin.Context) {
	time.Sleep(3 * time.Second)
}

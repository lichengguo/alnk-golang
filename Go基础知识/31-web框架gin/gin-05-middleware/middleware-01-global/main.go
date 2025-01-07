package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// 全局中间件

// gin可以构建中间件，但他只对注册过的路由函数起作用
// 对于分组路由，嵌套使用中间件，可以限定中间件的作用范围
// 中间件分为全局中间件，单个路由中间件和群组中间件
// gin 的中间件必须是一个 gin.HandlerFunc 类型

// 定义一个中间件
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 计算函数执行了多长时间
		t := time.Now()
		fmt.Println("中间件开始执行了...")

		// 设置变量到Context的key中，然后调用中间件的函数可以通过Get()获取
		c.Set("request", "我是中间件")

		// c.Next() 是执行调用中间件的函数
		c.Next()

		// 调用中间件的函数执行完毕后的一些其他事项
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕了...", status)

		//
		t2 := time.Since(t)
		fmt.Println("函数执行了多久: ", t2)
	}
}

func main() {
	// 1 创建路由
	r := gin.Default()

	// 2 注册中间件
	r.Use(MiddleWare())
	{
		// 3 绑定路由
		r.GET("/middleware", func(c *gin.Context) {
			// 取中间件里面的值
			req, _ := c.Get("request")
			fmt.Println("中间件中设置的值是: ", req)

			// 返回给客户端的值
			c.JSON(200, gin.H{"request": req})
		})
	}

	// 4 启动
	_ = r.Run(":8080")
}

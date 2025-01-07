package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// 同步异步
// goroutine 机制可以方便的实现异步处理
// 另外，在启动新的 goroutine 时，不应该使用原始上下文，必须使用它的只读副本

func main() {
	// 1 创建路由
	r := gin.Default()

	// 异步
	r.GET("/long_async", func(c *gin.Context) {
		// 需要一个副本
		copyContext := c.Copy()

		// 在启动新的 goroutine 时，不应该使用原始上下文，必须使用它的只读副本
		go func() {
			time.Sleep(5 * time.Second)
			log.Println("异步执行:" + copyContext.Request.URL.Path)
		}()

		c.String(200, "异步执行") // 直接返回数据给客户端，不需要等待5秒
	})

	// 同步
	r.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		log.Println("同步执行:" + c.Request.URL.Path)
		c.String(200, "同步执行") // 返回数据给客户端，需要等待5秒钟，客户端才能拿到数据
	})

	// 启动程序
	r.Run(":8080")
}

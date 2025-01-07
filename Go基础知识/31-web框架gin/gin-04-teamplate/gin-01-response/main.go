package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

// 各种数据格式的响应到客户端

func main() {
	// 1 创建路由
	r := gin.Default()

	// 2.1 json格式的响应
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "json格式的响应",
			"status":  200,
		})
	})

	// 2.2 结构体格式响应
	r.GET("/someStruct", func(c *gin.Context) {
		var msg struct {
			Name    string
			Message string
			Number  int
		}
		msg.Name = "root"
		msg.Message = "结构体格式响应"
		msg.Number = 123

		c.JSON(200, msg)
	})

	// 2.3 XML
	r.GET("/someXML", func(c *gin.Context) {
		c.XML(200, gin.H{
			"message": "abc",
		})
	})

	// 2.4 YAML
	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(200, gin.H{
			"name": "zhangshan",
		})
	})

	// 2.5 protobuf 谷歌开发的高效存储读取的工具
	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		// 定义数据
		label := "label"
		// 传protobuf格式数据
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		c.ProtoBuf(200, data)
	})

	// 3 监听
	r.Run(":8080")
}

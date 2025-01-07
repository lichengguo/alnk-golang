package main

import "github.com/gin-gonic/gin"

// HTML 模板渲染

// gin支持加载HTML模板, 然后根据模板参数进行配置并返回相应的数据，本质上就是字符串替换
// LoadHTMLGlob()方法可以加载模板文件

func main() {
	// 1 创建路由
	r := gin.Default()

	// 2 加载模板
	r.LoadHTMLGlob("./templates/*")

	// 3 绑定路由
	r.GET("/index", func(c *gin.Context) {
		// 根据文件名渲染
		// 最终json将title替换
		c.HTML(200, "index.tmpl", gin.H{
			"title": "我是一个很大的标题",
		})
	})

	// 4 启动程序监听端口
	r.Run(":8080")
}

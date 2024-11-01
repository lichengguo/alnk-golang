package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Cookie练习
// 模拟实现权限验证中间件
// 有2个路由，login和home
// login用于设置cookie
// home是访问查看信息的请求
// 在请求home之前，先跑中间件代码，检验是否存在cookie

// 权限检验中间件
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookieValue, err := c.Cookie("key_cookie")
		// 有cookie的情况
		if err == nil {
			// 检验cookie的键
			if cookieValue == "123" {
				c.Next()
				return
			}
		}
		// 返回错误
		c.JSON(http.StatusUnauthorized, gin.H{"error": "权限校验未通过"})
		// 如果校验未通过，不再同调后续的函数
		c.Abort()
		return
	}
}

func main() {
	// 创建路由
	r := gin.Default()

	// 绑定路由
	r.GET("/login", func(c *gin.Context) {
		// 设置cookie
		c.SetCookie("key_cookie", "123", 180, "/", "127.0.0.1", false, true)
		// 返回信息
		c.String(200, "cookie设置成功")
	})

	// 访问home之前，先调用中间件中权限验证
	r.GET("/home", AuthMiddleWare(), func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "我是一个很漂亮的主页"})
	})

	r.Run(":8080")
}

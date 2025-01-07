package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET URL 数据解析和绑定

// 定义接收的结构体
type Login struct {
	// binding:"requeired" 修饰的字段，如果接收为空值，则报错，是必须字段
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	// 创建路由
	r := gin.Default()

	// 绑定路由
	r.GET("/loginUrl", loginUrl)

	// 监听
	r.Run(":8080")

}

func loginUrl(c *gin.Context) {
	// 声明接收的变量
	var loginData Login

	// 解析数据
	err := c.ShouldBindQuery(&loginData)
	if err != nil {
		// 如果不能解析，返回错误给客户端
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 判断账号密码是否正确
	if loginData.User != "root" || loginData.Password != "123" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "账号或者密码错误",
		})
		return
	}

	// 密码正确
	c.JSON(200, gin.H{
		"status": "200",
		"msg":    "登录成功",
	})
}

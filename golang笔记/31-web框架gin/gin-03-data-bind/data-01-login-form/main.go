package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 表单数据解析和绑定

// 定义接收的结构体
type Login struct {
	// binding:"requeired" 修饰的字段，如果接收为空值，则报错，是必须字段
	User     string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	// 创建路由
	r := gin.Default()

	// 绑定路由
	r.POST("/loginForm", loginForm)

	// 监听
	r.Run(":8080")

}

func loginForm(c *gin.Context) {
	// 声明接收的变量
	var formData Login

	// Bind() 解析并且绑定 form 格式
	// 是根据请求头中的 content-type 自动推断的
	err := c.ShouldBind(&formData)
	// err := c.Bind(&formData)
	if err != nil {
		// 这是返回json数据给客户端
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 判断账号密码是否正确
	if formData.User != "root" || formData.Password != "123" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "账号或者密码错误",
		})
		return
	}

	// 密码正确
	c.JSON(200, gin.H{
		"status": "200",
		"msg":    "登录成功Form",
	})
}

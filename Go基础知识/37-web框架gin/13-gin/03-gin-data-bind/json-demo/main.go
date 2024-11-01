package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// json数据解析和绑定
// 把客户端传递过来的json数据绑定到结构体

// 定义接收的结构体
type Login struct {
	// binding:"requeired" 修饰的字段，如果接收为空值，则报错，是必须字段
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func main() {
	// 创建路由
	r := gin.Default()

	// json 绑定
	r.POST("/loginJson", loginJson)

	// 监听
	r.Run(":8080")
}

func loginJson(c *gin.Context) {
	// 声明接收客户端传过来的json数据的变量
	var jsonData Login

	// 将request的body中的数据，自动按照 json 格式解析到结构体
	err := c.ShouldBindJSON(&jsonData)
	if err != nil {
		// 如果不能解析，返回错误给客户端
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 判断账号密码是否正确
	if jsonData.User != "root" || jsonData.Password != "123" {
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

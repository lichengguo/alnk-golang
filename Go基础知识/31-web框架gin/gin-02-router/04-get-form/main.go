package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 获取表单参数

// 表单传输为post请求，http常见的传输格式为四种
// application/json
// application/x-www-form-urlencoded
// application/xml
// multipart/form-data

// 表单参数可以通过PostForm()方法获取
// 该方法默认解析的是x-www-form-urlencoded或from-data格式的参数

func main() {
	r := gin.Default()
	r.POST("/", postURL) // 获取表单参数
	_ = r.Run(":8080")
}

func postURL(c *gin.Context) {
	// 表单参数设置默认值
	type1 := c.DefaultPostForm("type", "alert") // DefaultPostForm()若参数不存在，则返回默认值

	// 接收其他的Key-Value
	userName := c.PostForm("username")
	passWord := c.PostForm("password")

	// 多选框
	hobbys := c.PostFormArray("hobby")

	// 返回
	c.String(http.StatusOK, fmt.Sprintf("type is %s, username is %s, password is %s, hobbys is %v",
		type1, userName, passWord, hobbys))
}

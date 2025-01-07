package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// gin 上传单个文件

// multipart/form-data格式用于文件上传
// gin文件上传与原生的net/http方法类似，不同在于gin把原生的request封装到c.Request中

func main() {
	r := gin.Default()
	r.POST("/upload", upload)
	_ = r.Run(":8080")
}

// upload 上传单个文件
func upload(c *gin.Context) {
	// 表单取文件
	file, _ := c.FormFile("file")
	fmt.Println(file.Filename)

	// 传到项目的根目录，名字就用本身的
	// 建议统一转换名字，为了安全起见
	_ = c.SaveUploadedFile(file, file.Filename)

	// 返回给客户端
	c.String(200, fmt.Sprintf("%s upload!", file.Filename))
}

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// gin 上传多个文件

func main() {
	r := gin.Default()
	// 限制表单上传大小为8MB，默认为32MB
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", upload)
	_ = r.Run(":8080")
}

// upload 上传多个文件
func upload(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
		return
	}

	// 获取所有图片
	files := form.File["file"]
	fmt.Println("files", files)
	// 遍历所有图片
	for _, file := range files {
		// 逐个保存
		err := c.SaveUploadedFile(file, file.Filename)
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("save pic failed, %s", err.Error()))
			return
		}
	}

	c.String(200, fmt.Sprintf("upload ok %d file", len(files)))
}

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 表单信息
type form struct {
	// binding:"requeired" 修饰的字段，如果接收为空值，则报错，是必须字段
	Title string `form:"title" binding:"required"`
	Price int64  `form:"price" binding:"required"`
}

// bookList 展示所有图书
func bookList(c *gin.Context) {
	booList, err := QueryAllData()
	if err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "获取图书列表失败",
		})
		return
	}
	// 返回数据给客户端
	c.HTML(http.StatusOK, "book_list.html", gin.H{
		"code": 0,
		"data": booList,
	})
}

// getAddBookHtml 获取图书界面
func getAddBookHtml(c *gin.Context) {
	c.HTML(200, "new_book.html", "")
}

// addBook 添加图书
func addBook(c *gin.Context) {
	// 声明接受变量
	var formData form
	// Bind() 解析并且绑定 form 格式
	// 是根据请求头中的 content-type 自动推断的
	err := c.Bind(&formData)
	if err != nil {
		c.JSON(400, gin.H{
			"msg":   "添加图书失败",
			"error": err.Error(),
		})
		return
	}

	// 写入到数据库
	err = InsertBookData(formData.Title, formData.Price)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "不能写入数据库，添加图书失败",
			"error": err.Error(),
		})
		return
	}

	// 添加成功，直接跳转到lsit页面
	c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:8080/book/list")
}

// deleteBook 删除一本图书
func deleteBook(c *gin.Context) {
	// 获取要删除的ID
	id := c.Query("id")
	err := DeleteBookData(id)
	if err != nil {
		c.JSON(400, gin.H{"msg": "获取不到ID,删除图书失败"})
		return
	}

	// 删除图书以后直接跳转到list界面
	c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:8080/book/list")
}

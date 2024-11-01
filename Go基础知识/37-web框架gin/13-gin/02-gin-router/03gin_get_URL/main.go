package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// URL参数

func main() {
	r := gin.Default()
	r.GET("/welcome", getURL) //RUL参数
	_ = r.Run(":8080")
}

// getURL 获取URL中的参数
// http://127.0.0.1:8080/welcome?name=alnk&age=18
// 结果 name:alnk age:18
// http://127.0.0.1:8080/welcome?name=alnk&age=18&age=100
// 结果 name:alnk age:18
// http://127.0.0.1:8080/welcome?name=alnk
// 结果 name:alnk age:
// http://127.0.0.1:8080/welcome?age=18
// 结果 name:jack age:18
func getURL(c *gin.Context) {
	name := c.DefaultQuery("name", "jack") // DefaultQuery() 若参数不存在，则返回默认值
	age := c.Query("age")                  // Query()若不存在，返回空串
	fmt.Printf("name:%s age:%s\n", name, age)
	c.String(http.StatusOK, "OK!")
}

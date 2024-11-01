package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Restful风格的API
// gin支持Restful风格的API
// 即Representational State Transfer的缩写。
// 直接翻译的意思是"表现层状态转化"，是一种互联网应用程序的API设计理念：URL定位资源，用HTTP描述操作
// 1.获取	/blog/getXxx	   Get     blog/Xxx
// 2.添加	/blog/addXxx       POST    blog/Xxx
// 3.修改	/blog/updateXxx    PUT     blog/Xxx
// 4.删除	/blog/delXxxx      DELETE  blog/Xxx

// API参数
// 可以通过Context的Param方法来获取API参数

func main() {
	r := gin.Default()
	r.GET("/user/:name/*action", getAPI) //API参数: 匹配单个API参数 ，*会匹配后面所有的参数
	r.Run(":8080")
}

// getAPI 获取API参数
// http://127.0.0.1:8000/user/zhangsan/run
// 结果: zhangsan is /run
// http://127.0.0.1:8080/user/zhangshan/run/game
// 结果:zhangshan is /run/game
func getAPI(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	fmt.Println(name + " is " + action)
	c.String(http.StatusOK, name+" is "+action)
}

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// Cookie是什么
// HTTP是无状态协议，服务器不能记录浏览器的访问状态，也就是说服务器不能区分两次请求是否由同一个客户端发出
// cookie就是解决HTTP协议无状态的方案之一，中文是小甜饼的意思
// Cookie实际上就是服务器保存在浏览器上的一段信息。
// 浏览器有了Cookie之后，每次向服务器发送请求时都会同时将该信息发送给服务器，服务器收到请求后，就可以根据该信息处理请求
// Cookie由服务器创建，并发送给浏览器，最终由浏览器保存

// Cookie的用途
// 保持用户登录状态
// 京东购物车

// Cookie的缺点
// 不安全，明文
// 增加带宽消耗
// 可以被禁用
// cookie有上限

// 测试服务端发送cookie给客户端，客户端请求时携带cookie

func main() {
	// 创建路由
	r := gin.Default()

	// 绑定路由
	// 服务端给客户端cookie
	r.GET("/cookie", func(c *gin.Context) {
		// 获取客户端是否携带cookie
		cookie, err := c.Cookie("key_cookie")
		if err != nil {
			cookie = "NotSet"
			// 如果客户端没有携带cookie，那么就给客户端设置cookie
			// maxAge int 单位为秒 cookie存活时间
			// path cookie所在的目录
			// domain string 域名
			// secure 是否只能通过https访问
			// httpOnly bool 是否允许别人通过js获取自己的cookie
			c.SetCookie("key_cookie", "value_cookie", 60, "/", "127.0.0.1", false, true)
		}

		// 打印获取到的cookie的值
		fmt.Println("cookie的值是:", cookie)
	})

	// 启动程序
	r.Run(":8080")
}

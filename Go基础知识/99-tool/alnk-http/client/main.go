package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// 创建一个HTTP客户端
	client := &http.Client{}

	// 创建一个GET请求
	resp, err := client.Get("http://www.baidu.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 获取响应的状态码
	statusCode := resp.StatusCode
	fmt.Println("status code:", statusCode)

	// 获取响应的头部
	headers := resp.Header
	fmt.Println("headers:", headers)

	// 获取响应的内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// 输出响应的内容
	fmt.Println("body:", string(body))

}

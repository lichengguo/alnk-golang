package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 响应数据
	// w.Write([]byte("hello world!"))
	fmt.Fprintf(w, "hello world!!!")
}

func main() {
	// 注册路由
	http.HandleFunc("/hello", helloHandler)

	// 启动服务
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/register/", register) //注册
	http.HandleFunc("/", loginIndex)        //登录

	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		fmt.Printf("start http server failed, err:%v\n", err)
		return
	}
}

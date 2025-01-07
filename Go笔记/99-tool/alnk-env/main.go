package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	envStr := os.Environ() // 返回所有环境变量
	fmt.Println(envStr)

	fmt.Println("========= 遍历环境变量 =================")
	for _, env := range envStr {
		e := strings.Split(env, " ")
		fmt.Println(e)
	}

	// 插入环境变量
	os.Setenv("myname", "alnk")

	// 获取环境变量
	fmt.Println(os.Getenv("myname"))
}

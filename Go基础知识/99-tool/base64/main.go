package main

// base64编码
// base64解码

import (
	"encoding/base64"
	"fmt"
)

func main() {
	originStr := "hello world"
	fmt.Println("原始字符串:", originStr)

	// base64编码
	encodeStr := base64.StdEncoding.EncodeToString([]byte(originStr))
	fmt.Println("base64编码后的字符串:", encodeStr) // aGVsbG8gd29ybGQ=

	// base64解码
	decodeBytes, err := base64.StdEncoding.DecodeString(encodeStr)
	if err != nil {
		fmt.Println("解码失败:", err)
		return
	}
	fmt.Println("base64解码后的字符串:", string(decodeBytes))
}

package main

import (
	"fmt"
	"strings"
)

// 利用闭包实现的功能：增加指定的后缀。不存在则增加，存在则不改动

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	jpgFunc := makeSuffixFunc(".jpg")
	fmt.Println(jpgFunc("test1"))     //test1.jpg
	fmt.Println(jpgFunc("test2.jpg")) //test2.jpg

	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(txtFunc("a.txt")) //a.txt
	fmt.Println(txtFunc("b"))     //b.txt
}

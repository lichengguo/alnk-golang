package main

import (
	alnk "import-demo/calc"
	"fmt"
)

var x = 100

const pi = 3.14

func init() {
	fmt.Println("main init 自动执行！")
	fmt.Println(x, pi)
}

func main() {
	ret := alnk.Add(10, 20)
	fmt.Println(ret)
}

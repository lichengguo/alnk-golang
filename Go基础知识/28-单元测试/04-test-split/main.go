package main

import (
	"04-test-split/split_string"
	"fmt"
)

func main() {
	ret := split_string.Split("babcbef", "b")
	fmt.Println(ret)
}

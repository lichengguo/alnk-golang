package main

import (
	"fmt"
	"strconv"
)

/*
Go语言内置包之strconv
Go语言中strconv包实现了基本数据类型和其字符串表示的相互转换
strconv包实现了基本数据类型与其字符串表示的转换，主要有以下常用函数： Atoi()、Itoa()、parse系列、format系列、append系列
*/

func main() {
	// 从字符串中解析出对应的数据
	str := "10000"
	ret1, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return
	}
	fmt.Printf("%#v %T\n", ret1, ret1) //10000 int64

	// Atoi 字符串转换成int
	retInt, _ := strconv.Atoi(str)
	fmt.Printf("%#v %T\n", retInt, retInt) //10000 int

	// 从字符串中解析出布尔值
	boolStr := "true"
	boolValue, _ := strconv.ParseBool(boolStr)
	fmt.Printf("%#v %T\n", boolValue, boolValue) //true bool

	// 从字符串解析出浮点数
	floatStr := "1.234"
	floatValue, _ := strconv.ParseFloat(floatStr, 64)
	fmt.Printf("%#v %T\n", floatValue, floatValue) //1.234 float64

	// 把数字转换成字符串类型
	i := 97
	ret2 := string(i)
	fmt.Println(ret2) //a
	ret3 := fmt.Sprintf("%d", i)
	fmt.Printf("%#v\n", ret3) //"97"
	ret4 := strconv.Itoa(i)
	fmt.Printf("%#v %T\n", ret4, ret4) //"97" string
}

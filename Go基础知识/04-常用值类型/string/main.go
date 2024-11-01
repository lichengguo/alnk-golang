package main

import (
	"fmt"
	"strings"
)

// 字符串
// Go语言中字符串是用双引号包裹的
// Go语言中单引号包裹的是字符

// Go语言里的字符串的内部实现使用UTF-8编码

// s := "Hello 中国"

// 单独的字母、汉字、符号并且用单引号括起来的表示一个字符
// c1 := 'h'
// c2 := '1'
// c3 := '中'

// ASCII编码
// 字节：1字节=8Bit(8个二进制位)
// 1个字符'A'=1个字节

// 1个utf8编码的汉字 '中' = 一般占3个字节

func main() {
	// \ 反斜杠是具有特殊含义的，应该告诉程序写的 \ 就是一个单纯的 \
	// path := "D:\\Go\\src\\code.alnkedu.com\\studygo\\day01"
	// path := "'D:\\Go\\src\\code.alnkedu.com\\studygo\\day01'"
	// path := "\"D:\\Go\\src\\code.alnkedu.com\\studygo\\day01\""
	path := `"D:\Go\src\\code.alnkedu.com\studygo\day01\"`
	fmt.Println(path)

	s := "I'm ok"
	fmt.Println(s)

	// 多行字符串
	s2 := `
	aaa
bbb
		ccc
	`
	fmt.Println(s2)

	s3 := `D:\Go\src\code.alnkedu.com\studygo\day01`
	fmt.Println(s3)

	// 字符串相关操作
	// 统计字节数
	s4 := "hello中国"
	// 注意 len() 统计的字符串的时候，是统计字节数长度，而不是字符数，在utf8编码中，一个英文字符占用1个字节，一个中文字符一般占用3个字节
	fmt.Println(len(s4)) // 11

	// 字符串拼接
	name := "李响"
	world := "dsb"
	ss := name + world
	fmt.Println(ss)
	ss1 := fmt.Sprintf("%s - %s", name, world)
	fmt.Println(ss1)

	// strings
	// 分割
	ret := strings.Split(s3, "\\") // 注意这里的 \\ 前面的反斜杠是为了不让后面的反斜杠具有特殊意义
	fmt.Println(ret)

	// 包含
	fmt.Println("包含: ", strings.Contains(ss, "dsb"))
	fmt.Println("包含: ", strings.Contains(ss, "理想"))

	// 前缀
	fmt.Println("前缀: ", strings.HasPrefix(ss, "理想"))

	// 后缀
	fmt.Println("后缀: ", strings.HasSuffix(ss, "dsb"))

	// 查找字符串出现的索引
	s5 := "abcdeb"
	//     012345
	fmt.Println("索引: ", strings.Index(s5, "c"))
	fmt.Println("索引: ", strings.LastIndex(s5, "eb"))

	// 单独的字母、汉字、符号并且用单引号括起来的表示一个字符
	// 双引号的是字符串了
	s6 := 'c'
	fmt.Printf("s6: %T\n", s6) // int32 看不懂的话看下一个章节
	s7 := '中'
	fmt.Printf("s7: %T\n", s7) // int32
	s8 := "c"
	fmt.Printf("s8: %T\n", s8) // string

	// 切片拼接
	fmt.Println(strings.Join(ret, "+"))
	fmt.Println(strings.Join(ret, "-"))
}

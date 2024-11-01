package main

import "fmt"

func main() {
	// 回文判断 字符串从左往右读和从右往左读是一样的，那么就是回文。
	// 上海自来水来自海上 s[0]  s[len(s)-1]
	// 山西运煤车煤运西山

	ss := "a山西运煤车煤运西山a"
	r := make([]rune, 0, 255) //初始化 []rune切片 长度0，容量255

	// 把字符串中的字符拿出来放到一个切片中
	for _, c := range ss {
		r = append(r, c)
	}
	// 此时的切片r: [97 23665 35199 36816 29028 36710 29028 36816 35199 23665 97]

	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-1-i] {
			fmt.Printf("字符串 [%s] 不是回文\n", ss)
			return
		}
	}

	fmt.Printf("字符串 [%s] 是回文\n", ss)
}

package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		suffix      string // 需要合并的文件后缀 默认为yaml
		endFileName string // 合并后的文件名
	)

	// 命令行参数解析
	flag.StringVar(&suffix, "suffix", "yaml", "文件名的后缀,默认为yaml")                  // (变量名, 参数名, 默认值, 提示信息)
	flag.StringVar(&endFileName, "endFileName", "alnk-go-k8s", "默认为alnk-go-k8s") // (变量名, 参数名, 默认值, 提示信息)
	flag.Parse()                                                                 // 解析命令行参数

	if endFileName == "alnk-go-k8s" { // 如果只输入了匹配的后缀名，那么默认文件名需要加上后缀
		endFileName = endFileName + "." + suffix
	}

	fmt.Println(suffix, endFileName)
}

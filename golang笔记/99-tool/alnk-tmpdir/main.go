package main

import (
	"fmt"
	"os"
)

// 在程序运行时经常创建一些运行时用到，程序结束后就不再使用的数据
// 临时目录和文件对于上面的情况很有用，因为它不会随着时间的推移而污染文件系统

func main() {
	// 创建临时目录
	tmpDir, err := os.MkdirTemp("", "alnk-dir-")
	if err != nil {
		panic(err)
	}
	fmt.Println("Temp Dir: ", tmpDir) // Temp Dir:  /var/folders/6b/q9fjc5g95hl8t1n2hswyk8240000gn/T/alnk-1181480352

	// 创建临时文件
	tmpFile, err := os.CreateTemp(tmpDir, "alnk-file-")
	if err != nil {
		panic(err)
	}
	fmt.Println("Temp File: ", tmpFile.Name()) // Temp File:  /var/folders/6b/q9fjc5g95hl8t1n2hswyk8240000gn/T/alnk-1181480352/alnk-file-000000001

	// 临时文件写入数据
	_, err = tmpFile.Write([]byte("Hello, World!"))
	if err != nil {
		panic(err)
	}

	// 关闭临时文件
	err = tmpFile.Close()
	if err != nil {
		panic(err)
	}

	// 读取临时文件
	context, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		panic(err)
	}
	fmt.Println("Temp File Context: ", string(context)) // Temp File Context:  Hello, World!

	defer os.RemoveAll(tmpDir)      // 删除临时目录
	defer os.Remove(tmpFile.Name()) // 删除临时文件

}

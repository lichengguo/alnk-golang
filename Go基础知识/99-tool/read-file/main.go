package main

// Yaml文件合并

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

var wg sync.WaitGroup

// 获取当前目录下的文件信息
func getFile(suffix string) (fileNames []string, err error) {
	// 获取当前程序执行目录
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	// 获取目录下所有文件名称
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	// 删选文件名称 yaml
	for _, fileName := range files {
		if strings.HasSuffix(fileName.Name(), suffix) {
			fileNames = append(fileNames, fileName.Name())
		}
	}

	return fileNames, nil
}

// readFile 读取文件
func readFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

// writeFile 写入文件内容
func writeFile(fileObj *os.File, context []byte) error {
	// 写入内容
	wr := bufio.NewWriter(fileObj)
	_, err := wr.WriteString("---\n")
	if err != nil {
		return err
	}
	_, err = wr.WriteString(fmt.Sprintf("%s\n", context))
	if err != nil {
		return err
	}
	if err = wr.Flush(); err != nil { // 缓存写入到文件中
		return err
	}

	return nil
}

func main() {
	// 获取后缀yaml的文件
	fileNames, err := getFile("yaml")
	if err != nil {
		panic(err)
	}

	// 如文件数量庞大，可使用协程池对goroutine进行控制
	ch := make(chan []byte, len(fileNames))
	for _, v := range fileNames {
		wg.Add(1)
		go func(v string) {
			defer wg.Done()
			s, err := readFile(v)
			if err != nil {
				panic(err)
			}
			ch <- s
		}(v)
	}
	wg.Wait()

	close(ch) // 关闭通道

	// 获取一个文件句柄
	fileObj, err := os.OpenFile("alnk-go-k8s.yaml", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer fileObj.Close() // 关闭文件

	// 文件内容写入
	for v := range ch {
		err := writeFile(fileObj, v)
		if err != nil {
			panic(err)
		}
	}
}

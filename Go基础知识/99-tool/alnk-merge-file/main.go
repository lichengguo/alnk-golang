package main

// 文件合并工具
// merge-file 默认合并yaml文件，默认合并后的文件名为alnk-go-k8s.yaml
// merge-file -suffix yaml -endFileName alnk-go-k8s.yaml
// merge-file -suffix txt -endFileName endfilename.txt
// merge-file -suffix txt

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/panjf2000/ants"
)

// getFile 获取当前目录下的文件信息
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

	// 选定文件名称 yaml
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
func writeFile(fileObj *os.File, suffix string, context []byte) error {
	// 写入内容
	wr := bufio.NewWriter(fileObj)

	// 如果是yaml文件，每个文件之间添加分隔符
	if suffix == "yaml" {
		_, err := wr.WriteString("---\n")
		if err != nil {
			return err
		}
	}

	_, err := wr.WriteString(fmt.Sprintf("%s\n", context))
	if err != nil {
		return err
	}

	if err = wr.Flush(); err != nil { // 缓存写入到文件中
		return err
	}

	return nil
}

func main() {
	var (
		suffix      string // 需要合并的文件后缀 默认为yaml
		endFileName string // 合并后的文件名
	)

	// 命令行参数解析
	flag.StringVar(&suffix, "suffix", "yaml", "文件名的后缀,默认为yaml")
	flag.StringVar(&endFileName, "endFileName", "alnk-go-k8s", "默认为alnk-go-k8s")
	flag.Parse()
	if endFileName == "alnk-go-k8s" { // 如果只输入了匹配的后缀名，那么默认文件名需要加上后缀
		endFileName = endFileName + "." + suffix
	}

	// 获取程序执行目录下指定后缀文件名称
	fileNames, err := getFile(suffix)
	if err != nil {
		panic(err)
	}

	// 使用协程池对goroutine进行控制
	ch := make(chan []byte, len(fileNames))
	// 创建并发数为4的协程池
	pool, err := ants.NewPool(4)
	if err != nil {
		panic(err)
	}

	// 并发执行任务
	var wg sync.WaitGroup
	wg.Add(len(fileNames))
	for _, v := range fileNames {
		err := pool.Submit(func() {
			defer wg.Done()
			s, err := readFile(v)
			if err != nil {
				panic(err)
			}
			ch <- s
		})
		if err != nil {
			panic(err)
		}

	}
	// fmt.Println(runtime.NumGoroutine()) // 查看当前程序中的协程数量
	wg.Wait()      // 等待所有任务执行完成
	pool.Release() // 关闭协程池，释放资源
	close(ch)      // 关闭通道

	// 获取一个文件句柄
	fileObj, err := os.OpenFile(endFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer fileObj.Close() // 关闭文件

	// 文件内容写入
	for v := range ch {
		err := writeFile(fileObj, suffix, v)
		if err != nil {
			panic(err)
		}
	}

	fmt.Printf("文件合并完成,合并后的文件为:[%s]\n", endFileName)
}

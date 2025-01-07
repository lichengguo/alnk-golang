package main

import (
	"fmt"

	"github.com/nxadm/tail"
)

// tail包
// 监听读取日志文件，如果文件有更新，那么能够立即读取到，跟Linux中的tail命令很像
// https://godoc.org/github.com/hpcloud/tail

func main() {
	// 1.日志文件路径和名称
	fileName := "./my.log"

	// 2.tail的配置文件
	config := tail.Config{
		ReOpen:    true,                                 // Reopen recreated files (tail -F),跟Linux中的 tail -F 命令一样
		Follow:    true,                                 // 是否跟随, 跟Linux中的 tail -f 命令一样
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件哪个地方开始读
		MustExist: false,                                // false:文件不存在不报错;true:文件不存在则报错
		Poll:      true,                                 // 轮询文件更改而不是使用inotify
	}

	// 3.获取日志文件句柄对象
	tails, err := tail.TailFile(fileName, config)
	if err != nil {
		fmt.Printf("tail file failed, err:%v\n", err)
		return
	}

	//4.开始监听读取日志文件
	for line := range tails.Lines {
		fmt.Println(line.Text)
	}
}

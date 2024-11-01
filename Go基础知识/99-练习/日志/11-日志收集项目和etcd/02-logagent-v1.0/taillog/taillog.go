package taillog

import (
	"fmt"
	"github.com/nxadm/tail"
)

//实时读取日志文件内容到channel通道

//
var (
	tailObj *tail.Tail //声明全局日志文件句柄变量
)

//Init 初始化实时读取日志
func Init(fileName string) (err error) {
	//1.配置文件
	config := tail.Config{
		ReOpen:    true,                                 //Reopen recreated files (tail -F),跟Linux中的 tail -F 命令一样
		Follow:    true,                                 //是否跟随, 跟Linux中的 tail -f 命令一样
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, //从文件哪个地方开始读
		MustExist: false,                                //false:文件不存在不报错;true:文件不存在则报错
		Poll:      true,                                 //轮询文件更改而不是使用inotify
	}
	//2.获取日志文件句柄对象
	tailObj, err = tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail file failed, err:", err)
		return
	}
	return
}

//
func ReadChan() <-chan *tail.Line {
	return tailObj.Lines
}

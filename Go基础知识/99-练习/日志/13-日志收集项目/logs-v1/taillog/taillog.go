package taillog

import (
	"fmt"
	"github.com/nxadm/tail"
	"logs-v1/conf"
)

type LogData struct {
	Topic string
	Data  string
}

var logDataChan = make(chan *LogData, 50000) //日志内容存储通道

// Init 初始化函数
func Init(cfg []conf.LogFile) {
	//1.配置tail包配置文件
	tailConfig := tail.Config{
		ReOpen:    true,                                 //Reopen recreated files (tail -F),跟Linux中的 tail -F 命令一样
		Follow:    true,                                 //是否跟随, 跟Linux中的 tail -f 命令一样
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, //从文件哪个地方开始读
		MustExist: false,                                //false:文件不存在不报错;true:文件不存在则报错
		Poll:      true,                                 //轮询文件更改而不是使用inotify
	}

	//2.循环读取多个日志文件内容
	for _, config := range cfg {
		//2.1获取日志文件句柄对象
		tailObj, err := tail.TailFile(config.FilePath, tailConfig)
		if err != nil {
			fmt.Println("tail file failed, err:", err)
			return
		}
		//2.2每个日志文件开启一个goroutine去读取，然后写入到通道中
		go sendToChan(tailObj, config.Topic)
	}
	return
}

// sendToChan 读取日志文件内容，写入到通道
func sendToChan(tailObj *tail.Tail, topic string) {
	for line := range tailObj.Lines {
		logData := &LogData{
			Topic: topic,
			Data:  line.Text,
		}
		logDataChan <- logData
	}
}

// getInfoToChan 暴露一个可以读取通道数据的函数
func GetInfoToChan() <-chan *LogData {
	return logDataChan
}

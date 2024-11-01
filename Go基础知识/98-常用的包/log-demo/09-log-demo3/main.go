package main

import (
	"fmt"
	"log"
	"os"
)

/*
创建logger
log标准库中还提供了一个创建新logger对象的构造函数–New，支持我们创建自己的logger示例。New函数的签名如下
func New(out io.Writer, prefix string, flag int) *Logger
New创建一个Logger对象。其中，参数out设置日志信息写入的目的地。参数prefix会添加到生成的每一条日志前面。参数flag定义日志的属性（时间、文件等等）

Go内置的log库功能有限，例如无法满足记录不同级别日志的情况，我们在实际的项目中根据自己的需要选择使用第三方的日志库，如logrus、zap等
*/

func main() {
	//logger := log.New(os.Stdout, "<new sucess!>", log.Ldate|log.Ltime|log.Lshortfile)
	//logger.Println("这是自定义的logger记录日志")

	logFile, err := os.OpenFile("./log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	logger := log.New(logFile, "<new sucess!>", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("这是自定义的logger记录日志")
}

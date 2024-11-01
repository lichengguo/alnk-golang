package mylogs

/*
日志包
1.可以同时往终端和日志文件输入日志
2.日志分为5种级别[trace debug warning info error]

用法 ex：
//声明成全局变量，好让所有的函数都能调用
var logger = mylogs.NewLogger("info", "./logs/", "log.txt", true)
logger.Trace("这是一条Trace日志f1()")
logger.Debug("这是一条debug日志f1()...")
logger.Warning("warning日志f1()...")
logger.Info("info日志f1()")
logger.Error("错误日志f1()...")
*/

import (
	"fmt"
	"os"
)

// Logger 定义一个日志结构体
type Logger struct {
	Level    string //日志等级[trace debug warning info error]
	FilePath string //日志存放路径
	FileName string //日志存放文件名称
	Tag      bool   //true:往屏幕打印日志
}

// NewLogger Logger结构体构造函数
// tag参数 是否往终端输出日志
func NewLogger(level string, filepath, filename string, tag bool) Logger {
	return Logger{
		Level:    level,
		FilePath: filepath,
		FileName: filename,
		Tag:      tag,
	}
}

//方法
//Trace
func (l *Logger) Trace(logContent string) {
	l.writeLogFile("trace", logContent)
}

//Debug
func (l *Logger) Debug(logContent string) {
	l.writeLogFile("debug", logContent)
}

//Warning
func (l *Logger) Warning(logContent string) {
	l.writeLogFile("warning", logContent)
}

//Info
func (l *Logger) Info(logContent string) {
	l.writeLogFile("info", logContent)
}

//Error
func (l *Logger) Error(logContent string) {
	l.writeLogFile("error", logContent)
}

//写入日志方法
func (l *Logger) writeLogFile(lv string, logContent string) {
	lvl, _ := parseLogLevel(lv)
	level, ok := parseLogLevel(l.Level)

	if !ok {
		fmt.Println("日志配置的等级不正确")
		os.Exit(1)
	}

	if level <= lvl {
		l.writeFile(lv, logContent)
	}
}

//把日志内容写入到文件方法
func (l *Logger) writeFile(lv string, logContent string) {
	filePath := l.FilePath + l.FileName                 //拼接文件日志路径和日志名称
	logContent = fmt.Sprintf("[%s] %s", lv, logContent) //拼接日志内容

	if l.Tag { //往屏幕输出日志内容
		fmt.Println(logContent)
	}

	fileObj, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("log file not found")
		os.Exit(1)
	}

	defer fileObj.Close()

	fmt.Fprintln(fileObj, logContent)

}

//解析日志等级函数
func parseLogLevel(str1 string) (int, bool) {
	switch str1 {
	case "trace":
		return 1, true
	case "debug":
		return 2, true
	case "warning":
		return 3, true
	case "info":
		return 4, true
	case "error":
		return 5, true
	default:
		return 0, false
	}
}

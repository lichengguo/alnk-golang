package mylog

import (
	"errors"
	"path"
	"runtime"
	"strings"
)

/*
思路：直接通过 通道 当做中间件临时存储日志信息
一个主线程往里面写日志内容，另外一个线程负责从通道中把日志信息拿出来写入到文件

//var conlog mylog.ConLog = mylog.NewConLog("error", false)
//终端输出日志
//参数1: 日志等级[debug trace info warning error fatal]
//参数2: 是否记录日志 true:记录

//var fileLog mylog.FileLog = *mylog.NewFileLog("Debug", "./logs", "test.log", 5*1024*1024, true)
//往文件写日志(异步写入)
//参数1: 日志等级[debug trace info warning error fatal]
//参数2: 目录名称
//参数3: 文件名称
//参数4: 每个日志文件保存大小(单位:B)
//参数5: 是否记录日志 true:记录

//使用方法示例
1.往终端
var conlog mylog.ConLog = mylog.NewConLog("error", false)
conlog.Debug("f1(debug).......")

2.往文件
var fileLog mylog.FileLog = *mylog.NewFileLog("Debug", "./logs", "test.log", 5*1024*1024, true)
fileLog.Debug("f3(debug)..id: %d name: %s", id, name)
fileLog.Trace("f3(Trace).......")

*/

//日志等级常量
const (
	UNKNOW uint8 = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

//parseLogLevelToUint8 解析日志等级参数,把字符串转换为unit8类型
func parseLogLevelToUint8(s string) (uint8, error) {
	switch strings.ToLower(s) {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别!")
		return UNKNOW, err
	}
}

//parseLogLevelToString 把日志等级从uint8转为string
func parseLogLevelToInt8(lv uint8) (level string) {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "UNKNOW"
	}
}

//获取打印日志时的行号和文件名称信息
func getLogInfo(skip int) (fileName, funcName string, lineNo int) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return
	}
	funcName = runtime.FuncForPC(pc).Name() //函数名
	fileName = path.Base(file)              //文件名
	lineNo = line                           //行号
	return
}

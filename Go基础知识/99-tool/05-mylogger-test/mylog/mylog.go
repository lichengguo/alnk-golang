package mylog

import (
	"errors"
	"path"
	"runtime"
	"strings"
)

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
	//funcName = strings.Split(funcName, "/")[:]
	fileName = path.Base(file) //文件名
	lineNo = line              // 行号
	return
}

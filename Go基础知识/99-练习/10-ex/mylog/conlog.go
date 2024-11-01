package mylog

import (
	"fmt"
	"time"
)

//终端日志结构体
type ConLog struct {
	Level string //日志级别[debug trace info warning error fatal]
	Tag   bool   //是否在终端输出日志 true:输出
}

//NewConLog 结构体构造函数
func NewConLog(s string, t bool) ConLog {
	return ConLog{
		Level: s,
		Tag:   t,
	}
}

//log 往终端输出日志的方法
func (c ConLog) log(lvl uint8, msg string, a ...interface{}) {
	msg = fmt.Sprintf(msg, a...)
	level, err := parseLogLevelToUint8(c.Level)
	if err != nil {
		fmt.Println(err)
		return
	}
	//转换日志级别的数据类型，用于终端输出
	lvlString := parseLogLevelToInt8(lvl)
	// 获取行号文件名函数名等信息
	fileName, funcName, lineNo := getLogInfo(3)
	//日志等级判断和是否终端打印
	if level <= lvl && c.Tag {
		//获取当前格式化时间
		dateString := time.Now().Format("2006/01/02 15:04:05")
		//[时间][级别][文件名:函数名:行号][日志内容]
		fmt.Printf("[%s] [%s] [%s:%s:%d] [%s]\n", dateString, lvlString, fileName, funcName, lineNo, msg)
	}
}

//Debug 方法
func (c ConLog) Debug(msg string, a ...interface{}) {
	c.log(DEBUG, msg, a...)
}

//Trace 方法
func (c ConLog) Trace(msg string, a ...interface{}) {
	c.log(TRACE, msg, a...)
}

//Info 方法
func (c ConLog) Info(msg string, a ...interface{}) {
	c.log(INFO, msg, a...)
}

//Warning 方法
func (c ConLog) Warning(msg string, a ...interface{}) {
	c.log(WARNING, msg, a...)
}

//Error 方法
func (c ConLog) Error(msg string, a ...interface{}) {
	c.log(ERROR, msg, a...)
}

//
func (c ConLog) Fatal(msg string) {
	c.log(FATAL, msg)
}

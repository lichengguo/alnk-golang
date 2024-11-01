package main

import (
	"18-example/mylogs"
	"time"
)

// 声明成全局变量，好让所有的函数都能调用
var logger = mylogs.NewLogger("info", "./logs/", "log.txt", true)

func f1() {
	logger.Trace("这是一条Trace日志f1()")
	logger.Debug("这是一条debug日志f1()...")
	logger.Warning("warning日志f1()...")
	logger.Info("info日志f1()")
	logger.Error("错误日志f1()...")
}

func f2() {
	logger.Trace("这是一条Trace日志f2()")
	logger.Debug("这是一条debug日志f2()...")
	logger.Warning("warning日志f2()...")
	logger.Info("info日志f2()")
	logger.Error("错误日志f2()...")
}

func main() {
	for {
		f1()
		f2()
		time.Sleep(2 * time.Second)
	}
}

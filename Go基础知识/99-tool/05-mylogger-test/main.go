package main

import (
	"05-mylogger-test/mylog"
	//"time"
)

// 终端输出日志
// 参数1: 日志等级[debug trace info warning error fatal]
// 参数2: 是否记录日志 true:记录
var conlog mylog.ConLog = mylog.NewConLog("error", false)

// 往文件写日志
// 参数1: 日志等级[debug trace info warning error fatal]
// 参数2: 目录名称
// 参数3: 文件名称
// 参数4: 每个日志文件保存大小(单位:B)
// 参数5: 是否记录日志 true:记录
var fileLog mylog.FileLog = *mylog.NewFileLog("Debug", "./logs", "test.log", 1*1024*1024, true)

func f1() {
	conlog.Debug("f1(debug).......")
	conlog.Trace("f1(Trace).......")
	conlog.Info("f1(Info).......")
	conlog.Warning("f1(Warning).......")
	conlog.Error("f1(Error).......")
	conlog.Fatal("f1(Fatal).......")
}

func f2() {
	conlog.Debug("f2(debug).......")
	conlog.Trace("f2(Trace).......")
	conlog.Info("f2(Info).......")
	conlog.Warning("f2(Warning).......")
	conlog.Error("f2(Error).......")
	conlog.Fatal("f2(Fatal).......")
}

func f3() {
	for {
		fileLog.Debug("f3(debug).......")
		fileLog.Trace("f3(Trace).......")
		fileLog.Info("f3(Info).......")
		fileLog.Warning("f3(Warning).......")
		fileLog.Error("f3(Error).......")
		fileLog.Fatal("f3(Fatal).......")
		//time.Sleep(3 * time.Second)
	}
}

func main() {
	//f1()
	//
	//f2()
	//
	f3()
}

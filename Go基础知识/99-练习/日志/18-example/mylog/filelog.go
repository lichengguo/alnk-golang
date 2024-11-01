package mylog

import (
	"fmt"
	"os"
	"path"
	"time"
)

//logMsg 通道中的日志结构体
type logMsg struct {
	level      uint8  //日志等级
	fileName   string //日志名称
	dateString string //日志时间
	funcName   string //函数名
	lineNo     int    //行号
	msg        string //日志内容
}

//FileLog 用户调用日志包时候的结构体
type FileLog struct {
	Level      string       //日志级别[debug trace info warning error fatal]
	Tag        bool         //是否在文件输出日志 true:输出
	PathName   string       //日志文件路径
	FileName   string       //日志文件名称
	FileObj    *os.File     //日志文件句柄
	FileObjErr *os.File     //错误日志文件句柄
	FileSize   int64        //日志文件大小
	logChan    chan *logMsg //日志通道
}

//NewFileLog 结构体构造函数
func NewFileLog(level, pathName, fileName string, fileSize int64, t bool) *FileLog {
	//1.程序包被调用的时候，就要创建日志文件目录和日志文件错误文件
	fileObj := openLogFile(pathName, fileName) //一般日志文件
	errFileName := fileName + ".err"
	fileObjErr := openLogFile(pathName, errFileName) //错误日志文件

	//2.构造返回的结构体指针
	fl := &FileLog{
		Level:      level,
		Tag:        t,
		PathName:   pathName,
		FileName:   fileName,
		FileObj:    fileObj,               //文件句柄
		FileObjErr: fileObjErr,            //错误日志文件句柄
		FileSize:   fileSize,              //日志文件大小
		logChan:    make(chan *logMsg, 5), //通道初始化
	}
	//2.从通道中取出数据写入文件 开启goroutine
	go fl.logReadChanWriteFile()
	return fl
}

//######################################################################################################################
//##################################### 往通道中写入日志内容开始 ##########################################################
//######################################################################################################################
//logWriteChan 把日志写入到通道中
func (f *FileLog) logWriteChan(selfLevel uint8, msg string, a ...interface{}) {
	//1.判断日志等级，达标的写入通道
	//1.1用户自己设置的日志等级转化为uint8类型，好进行比较
	userLevel, err := parseLogLevelToUint8(f.Level)
	if err != nil {
		fmt.Printf("等级设置有问题,err:%v\n", err)
		return
	}
	//2.写入通道
	if userLevel <= selfLevel && f.Tag {
		//2.1拼接日志文件具体内容
		msg = fmt.Sprintf(msg, a...)
		// 获取行号文件名函数名等信息
		fileName, funcName, lineNo := getLogInfo(3)
		//获取当前格式化时间
		dateString := time.Now().Format("2006/01/02 15:04:05")
		//组装结构体内容，把日志的结构体指针写入通道，节省内存空间
		logTmp := &logMsg{
			level:      selfLevel,
			fileName:   fileName,
			dateString: dateString,
			funcName:   funcName,
			lineNo:     lineNo,
			msg:        msg,
		}
		select {
		case f.logChan <- logTmp:
		default:
			//如果通道中存储满了就丢掉日志
		}
	}
}

//6种级别的日志 写入到通道中
//Debug 方法
func (f *FileLog) Debug(msg string, a ...interface{}) {
	f.logWriteChan(DEBUG, msg, a...)
}

//Trace 方法
func (f *FileLog) Trace(msg string, a ...interface{}) {
	f.logWriteChan(TRACE, msg, a...)
}

//Info 方法
func (f *FileLog) Info(msg string, a ...interface{}) {
	f.logWriteChan(INFO, msg, a...)
}

//Warning 方法
func (f *FileLog) Warning(msg string, a ...interface{}) {
	f.logWriteChan(WARNING, msg, a...)
}

//Error 方法
func (f *FileLog) Error(msg string, a ...interface{}) {
	f.logWriteChan(ERROR, msg, a...)
}

//Fatal 方法
func (f *FileLog) Fatal(msg string, a ...interface{}) {
	f.logWriteChan(FATAL, msg, a...)
}

//######################################################################################################################
//##################################### 往通道中写入日志内容结束 ###########################################################
//######################################################################################################################

//######################################################################################################################
//##################################### 从通道中读取日志内容写入到日志文件中开始 ##############################################
//######################################################################################################################
//openLogFile 打开日志记录文件
func openLogFile(pathName, fileName string) *os.File {
	//日志文件全路径和名称
	logFileName := path.Join(pathName, fileName)
	fileObj, err := os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	//日志记录文件不能打开直接退出程序
	if err != nil {
		fmt.Println("open log file failed, err:%v\n", err)
		os.Exit(1)
	}
	return fileObj
}

//logReadChanWriteFile 从通道中读取日志写入文件
func (f *FileLog) logReadChanWriteFile() {
	for {
		//日志文件即将写入的时候判断一下日志文件的大小，看是否要切割日志文件
		if f.checkLogFileSize(f.FileObj) {
			//切割日志文件
			newFileObj, err := f.splitLogFile(f.FileObj)
			if err != nil {
				return
			}
			f.FileObj = newFileObj
			//time.Sleep(time.Second)
		}

		select {
		//从通道中读取日志
		case logTmp := <-f.logChan:
			//拼接要写入的日志内容
			levelString := parseLogLevelToInt8(logTmp.level) //把日志等级从uint8转为字符串类型
			//[时间][级别][文件名:函数名:行号][日志内容]
			logInfo := fmt.Sprintf("[%s] [%s] [%s:%s:%d] [%s]\n", logTmp.dateString, levelString, logTmp.fileName, logTmp.funcName, logTmp.lineNo, logTmp.msg)
			fmt.Fprintf(f.FileObj, logInfo) //写入到文件

			//如果日志等级大于ERROR，还要额外的写入到 err 文件中
			if logTmp.level >= ERROR {
				if f.checkLogFileSize(f.FileObjErr) {
					newFileObj, err := f.splitLogFile(f.FileObjErr)
					if err != nil {
						return
					}
					f.FileObjErr = newFileObj
				}
				//写入文件
				fmt.Fprintf(f.FileObjErr, logInfo)
			}
		default:
			//取不到日志的话，就休息500毫秒
			time.Sleep(time.Millisecond * 500)
		}
	}
}

//checkLogFileSize 检查日志文件的大小
func (f *FileLog) checkLogFileSize(file *os.File) bool {
	//获取当前日志文件的信息
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed22222, err:%v\n", err)
		return false
	}
	//如果当前文件大小 大于等于 日志文件的最大值，返回true
	return fileInfo.Size() >= f.FileSize

}

//splitLogFile 切割日志文件
func (f *FileLog) splitLogFile(file *os.File) (*os.File, error) {
	//1.获取需要切割的日志文件基础信息
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed, err:%v\n", err)
		return nil, err
	}
	//2.拿到当前的日志文件完整路径
	logName := path.Join(f.PathName, fileInfo.Name())
	//3.拼接成一个备份的名字
	nowStr := time.Now().Format("200601021504050000")
	newLogName := fmt.Sprintf("%s.bak%s", logName, nowStr)
	//4.关闭当前的日志文件
	file.Close()
	//5.备份
	os.Rename(logName, newLogName)
	//6.打开一个新的日志文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open new log file failed, err:%v\n", err)
		return nil, err
	}
	//7.将打开的新日志文件句柄（对象、指针）赋值给fileobj
	return fileObj, nil
}

//######################################################################################################################
//##################################### 从通道中读取日志内容写入到日志文件中结束 ##############################################
//######################################################################################################################

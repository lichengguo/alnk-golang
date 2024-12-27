package mylog

import (
	"fmt"
	"os"
	"path"
	"time"
)

//文件日志结构体
type FileLog struct {
	Level      string   //日志级别[debug trace info warning error fatal]
	Tag        bool     //是否在文件输出日志 true:输出
	PathName   string   //日志文件路径
	FileName   string   //日志文件名称
	FileObj    *os.File //日志文件句柄
	FileObjErr *os.File //错误日志文件句柄
	FileSize   int64    //日志文件大小

}

//NewFileLog 结构体构造函数
func NewFileLog(level, pathName, fileName string, fileSize int64, t bool) *FileLog {
	//程序包被调用的时候，就要创建日志文件目录和日志文件错误文件
	fileObj := initFile(pathName, fileName)
	errFileName := fileName + ".err"
	fileObjErr := initFile(pathName, errFileName)

	return &FileLog{
		Level:      level,
		Tag:        t,
		PathName:   pathName,
		FileName:   fileName,
		FileObj:    fileObj,    //文件句柄
		FileObjErr: fileObjErr, //错误日志文件句柄
		FileSize:   fileSize,   //日志文件大小
	}
}

//log 往文件输出日志的方法
func (f *FileLog) log(lvl uint8, msg string) {
	//把字符串类型的日志等级转换成uint8类型的日志等级
	level, err := parseLogLevelToUint8(f.Level) //实例化传进来的日志等级level
	if err != nil {
		fmt.Println(err)
		return
	}
	//转换日志级别的数据类型，用于文件日志输出
	lvlString := parseLogLevelToInt8(lvl)
	// 获取行号文件名函数名等信息
	fileName, funcName, lineNo := getLogInfo(3)
	//日志等级判断和是否日志文件输出
	if level <= lvl && f.Tag {
		//日志文件即将写入的时候判断一下日志文件的大小，看是否要切割日志文件
		if f.checkLogFileSize(f.FileObj) {
			newFileObj, err := f.splitLogFile(f.FileObj)
			if err != nil {
				return
			}
			f.FileObj = newFileObj
		}

		//获取当前格式化时间
		dateString := time.Now().Format("2006/01/02 15:04:05")
		//[时间][级别][文件名:函数名:行号][日志内容]
		fmt.Fprintf(f.FileObj, "[%s] [%s] [%s:%s:%d] [%s]\n", dateString, lvlString, fileName, funcName, lineNo, msg)
		//如果日志等级大于ERROR，还要额外的写入到 err 文件中
		if lvl >= ERROR {
			if f.checkLogFileSize(f.FileObjErr) {
				newFileObj, err := f.splitLogFile(f.FileObjErr)
				if err != nil {
					return
				}
				f.FileObjErr = newFileObj
			}
			fmt.Fprintf(f.FileObjErr, "[%s] [%s] [%s:%s:%d] [%s]\n", dateString, lvlString, fileName, funcName, lineNo, msg)
		}
	}
}

//打开一个日志文件
func initFile(pathName, fileName string) *os.File {
	logFileName := path.Join(pathName, fileName) //日志文件全路径和名称
	fileObj, err := os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:%v\n", err)
		os.Exit(1) //日志记录文件不能打开直接退出程序
	}
	//defer fileObj.Close() //关闭日志文件 不能再这里关闭文件，不然会写不进去
	return fileObj
}

//检查日志文件的大小
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

//切割
func (f *FileLog) splitLogFile(file *os.File) (*os.File, error) {
	//1.需要切割的日志文件
	nowStr := time.Now().Format("200601021504050000")
	fileInfo, err := file.Stat()

	if err != nil {
		fmt.Printf("get file info failed, err:%v\n", err)
		return nil, err
	}
	//2.拿到当前的日志文件完整路径
	logName := path.Join(f.PathName, fileInfo.Name())
	//3.拼接成一个备份的名字
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

//Debug 方法
func (f *FileLog) Debug(msg string) {
	f.log(DEBUG, msg)
}

//Trace 方法
func (f *FileLog) Trace(msg string) {
	f.log(TRACE, msg)
}

//Info 方法
func (f *FileLog) Info(msg string) {
	f.log(INFO, msg)
}

//Warning 方法
func (f *FileLog) Warning(msg string) {
	f.log(WARNING, msg)
}

//Error 方法
func (f *FileLog) Error(msg string) {
	f.log(ERROR, msg)
}

//Fatal 方法
func (f *FileLog) Fatal(msg string) {
	f.log(FATAL, msg)
}

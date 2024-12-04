### Go日志库Zap简单二次封装

#### 1. 在项目根目录或者项目其他目录下创建二次封装代码存放目录zaplog，其他目录名称也可以



#### 2. 新建config.go文件和zaplog文件，文件内容如下：

##### config.go

```go
// Package zaplog 封装zap日志库,配置文件
// @Author: Alnk
// @Description: 二次封装开源zap日志库
// @File: zaplog
// @Date: 2022/12/12 14:20
package zaplog

// 定义默认的常量
const (
	defaultBaseDirectoryName  = "logs"      // 日志根目录
	defaultInfoDirectoryName  = "info"      // info日志目录
	defaultWarnDirectoryName  = "warn"      // warn日志目录
	defaultErrorDirectoryName = "error"     // error日志目录
	defaultInfoFileName       = "info.log"  // info日志文件
	defaultWarnFileName       = "warn.log"  // warn日志文件
	defaultErrorFileName      = "error.log" // error日志文件
	defaultLogFileMaxSize     = 128         // 日志文件大小，单位：MB
	defaultLogFileMaxBackups  = 180         // 日志文件保留个数 多于180个文件后，清理比价旧的日志
	defaultLogFileMaxAge      = 1           // 日志文件一天一切隔
	defaultLogFileCompress    = false       // 日志文件是否压缩
	defaultLogPrintTag        = false       // true:在终端和文件同时输出日志; false:只在文件输出日志

)

// Config 配置文件结构体
type Config struct {
	BaseDirectoryName  string
	InfoDirectoryName  string
	WarnDirectoryName  string
	ErrorDirectoryName string
	InfoFileName       string
	WarnFileName       string
	ErrorFileName      string
	LogFileMaxSize     int
	LogFileMaxBackups  int
	LogFileMaxAge      int
	LogFileCompress    bool
	LogPrintTag        bool
}

// Option 定义配置选项函数
type Option func(*Config)

// SetBaseDirectoryName 自定义日志根目录
func SetBaseDirectoryName(name string) Option {
	return func(c *Config) {
		c.BaseDirectoryName = name
	}
}

// SetInfoDirectoryName 自定义info日志目录
func SetInfoDirectoryName(name string) Option {
	return func(c *Config) {
		c.InfoDirectoryName = name
	}
}

// SetWarnDirectoryName 自定义warn日志目录
func SetWarnDirectoryName(name string) Option {
	return func(c *Config) {
		c.WarnDirectoryName = name
	}
}

// SetErrorDirectoryName 自定义error日志目录
func SetErrorDirectoryName(name string) Option {
	return func(c *Config) {
		c.ErrorDirectoryName = name
	}
}

// SetInfoFileName 自定义info文件名
func SetInfoFileName(name string) Option {
	return func(c *Config) {
		c.InfoFileName = name
	}
}

// SetWarnFileName 自定义warn文件名
func SetWarnFileName(name string) Option {
	return func(c *Config) {
		c.WarnFileName = name
	}
}

// SetErrorFileName 自定义error文件名
func SetErrorFileName(name string) Option {
	return func(c *Config) {
		c.ErrorFileName = name
	}
}

// SetLogFileMaxSize 自定义日志文件大小
func SetLogFileMaxSize(size int) Option {
	return func(c *Config) {
		c.LogFileMaxSize = size
	}
}

// SetLogFileMaxBackups 自定义日志文件保留个数
func SetLogFileMaxBackups(size int) Option {
	return func(c *Config) {
		c.LogFileMaxBackups = size
	}
}

// SetLogFileMaxAge 自定义日志文件切隔间隔
func SetLogFileMaxAge(size int) Option {
	return func(c *Config) {
		c.LogFileMaxAge = size
	}
}

// SetLogFileCompress 自定义日志文件是否压缩
func SetLogFileCompress(compress bool) Option {
	return func(c *Config) {
		c.LogFileCompress = compress
	}
}

// SetLogPrintTag 自定义日志输出标记位 true:在终端和文件同时输出日志; false:只在文件输出日志
func SetLogPrintTag(tag bool) Option {
	return func(c *Config) {
		c.LogPrintTag = tag
	}
}

// NewConfig 应用函数选项配置
func NewConfig(opts ...Option) Config {
	// 初始化默认值
	defaultConfig := Config{
		BaseDirectoryName:  defaultBaseDirectoryName,
		InfoDirectoryName:  defaultInfoDirectoryName,
		WarnDirectoryName:  defaultWarnDirectoryName,
		ErrorDirectoryName: defaultErrorDirectoryName,
		InfoFileName:       defaultInfoFileName,
		WarnFileName:       defaultWarnFileName,
		ErrorFileName:      defaultErrorFileName,
		LogFileMaxSize:     defaultLogFileMaxSize,
		LogFileMaxBackups:  defaultLogFileMaxBackups,
		LogFileMaxAge:      defaultLogFileMaxAge,
		LogFileCompress:    defaultLogFileCompress,
		LogPrintTag:        defaultLogPrintTag,
	}

	// 依次调用opts函数列表中的函数，为结构体成员赋值
	for _, opt := range opts {
		opt(&defaultConfig)
	}

	return defaultConfig
}
```

##### zaplog.go

```go
// Package zaplog 封装zap日志库主程序文件
// @Author: Alnk
// @Description: 二次封装开源zap日志库
// @File: zaplog
// @Date: 2022/12/12 17:43
package zaplog

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger 全局变量，导出给调用者使用
var Logger *zap.Logger

// Init 初始化日志相关目录
func Init(config Config) error {
	// 创建日志根目录
	if _, err := os.Stat(config.BaseDirectoryName); os.IsNotExist(err) {
		err := os.MkdirAll(config.BaseDirectoryName, os.ModePerm)
		if err != nil {
			return fmt.Errorf("error creating directory, err: %v", err)
		}
	}

	// 创建日志子目录
	if err := os.MkdirAll(fmt.Sprintf("%s/%s", config.BaseDirectoryName, config.InfoDirectoryName), os.ModePerm); err != nil {
		return fmt.Errorf("error creating info directory, err: %v", err)
	}
	if err := os.MkdirAll(fmt.Sprintf("%s/%s", config.BaseDirectoryName, config.WarnDirectoryName), os.ModePerm); err != nil {
		return fmt.Errorf("error creating warn directory, err: %v", err)
	}
	if err := os.MkdirAll(fmt.Sprintf("%s/%s", config.BaseDirectoryName, config.ErrorDirectoryName), os.ModePerm); err != nil {
		return fmt.Errorf("error creating err directory, err: %v", err)
	}

	// 自定义初始化zap库
	initLogger(config)

	return nil
}

// getWriter 获取wirter文件写入
func getWriter(logBasePath, logLevelPath, LogFileName string, config Config) io.Writer {
	return &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s/%s", logBasePath, logLevelPath, LogFileName),
		MaxBackups: config.LogFileMaxBackups,
		MaxSize:    config.LogFileMaxSize,
		MaxAge:     config.LogFileMaxAge,
		Compress:   config.LogFileCompress,
	}
}

// initLog 初始化日志
func initLogger(c Config) {
	// 获取io.Writer实现
	infoWriter := getWriter(c.BaseDirectoryName, c.InfoDirectoryName, c.InfoFileName, c)
	warnWriter := getWriter(c.BaseDirectoryName, c.WarnDirectoryName, c.WarnFileName, c)
	errWriter := getWriter(c.BaseDirectoryName, c.ErrorDirectoryName, c.ErrorFileName, c)

	// 获取日志默认配置
	encoderConfig := zap.NewProductionEncoderConfig()

	// 自定义日志输出格式
	// 修改TimeKey
	encoderConfig.TimeKey = "time"
	// 修改MessageKey
	encoderConfig.MessageKey = "message"
	// 时间格式符合人类观看
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	// 日志等级大写INFO
	encoderConfig.EncodeLevel = func(lvl zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(lvl.CapitalString())
	}
	// 日志打印时所处代码位置
	encoderConfig.EncodeCaller = func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(caller.TrimmedPath())
	}
	// 加载自定义配置为json格式
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	// 自定义日志级别 info
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl <= zap.InfoLevel && lvl >= zap.DebugLevel
	})
	// 自定义日志级别 warn
	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl <= zap.WarnLevel && lvl > zap.InfoLevel
	})
	// 自定义日志级别 err
	errLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl <= zap.FatalLevel && lvl > zap.WarnLevel
	})

	// 日志文件输出位置
	var core zapcore.Core
	if c.LogPrintTag {
		//同时在文件和终端输出日志
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel), // info级别日志
			zapcore.NewCore(encoder, zapcore.AddSync(warnWriter), warnLevel), // warn级别日志
			zapcore.NewCore(encoder, zapcore.AddSync(errWriter), errLevel),   // error级别日志
			zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), zap.DebugLevel),
		)
	} else {
		//只在文件输出日志
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel), // info级别日志
			zapcore.NewCore(encoder, zapcore.AddSync(warnWriter), warnLevel), // warn级别日志
			zapcore.NewCore(encoder, zapcore.AddSync(errWriter), errLevel),   // error级别日志
		)
	}

	Logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}
```



#### 3. 导入zaplog包，进行初始化以及日志记录（示例）

```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/alnkzap/zaplog"
)

func main() {
    // 自定义配置，每个配置项的含义可以查看封装代码中的注释
	//c := NewConfig(
	//	SetBaseDirectoryName("alnklog"),
	//	SetInfoDirectoryName("alnkinfo"),
	//	SetWarnDirectoryName("alnkwarn"),
	//	SetErrorDirectoryName("alnkerr"),
	//	SetInfoFileName("alnk_info"),
	//	SetWarnFileName("alnk_warn"),
	//	SetErrorFileName("alnk_err"),
	//	SetLogFileMaxSize(1),
	//	SetLogFileMaxBackups(1),
	//	SetLogFileMaxAge(1),
	//	SetLogFileCompress(true),
	//	SetLogPrintTag(true),
	//)
    
    // 默认配置
	c := zaplog.NewConfig() 
	if err := zaplog.Init(c); err != nil {
		log.Println(err)
	}

	testLogger()

}

func testLogger() {
	startTime := time.Now()
    
	for i := 0; i < 1000000; i++ {
        // 标准记录
		zaplog.Logger.Info("Logger info")
		zaplog.Logger.Warn("Logger warn")
		zaplog.Logger.Error("Logger err")
        
        // 添加其他的key-value
        Logger.Info("Logger info", zap.Any("user", "alnk"))
	    Logger.Warn("Logger Warn", zap.Any("user", "alnk"))
	    Logger.Error("Logger Error", zap.Any("user", "alnk"))
	}
    
	fmt.Println("Logger执行时间: ", time.Since(startTime))
}

/* 输出记录
{"level":"ERROR","time":"2022-12-14 10:47:30.178","caller":"testing/testing.go:1439","message":"Logger err"}
{"level":"ERROR","time":"2022-12-14 10:47:30.179","caller":"testing/testing.go:1439","message":"Logger Error","user":"alnk"}

{"level":"INFO","time":"2022-12-14 10:47:30.131","caller":"testing/testing.go:1439","message":"Logger info"}
{"level":"INFO","time":"2022-12-14 10:47:30.179","caller":"testing/testing.go:1439","message":"Logger info","user":"alnk"}

{"level":"WARN","time":"2022-12-14 10:47:30.178","caller":"testing/testing.go:1439","message":"Logger warn"}
{"level":"WARN","time":"2022-12-14 10:47:30.179","caller":"testing/testing.go:1439","message":"Logger Warn","user":"alnk"}
*/
```

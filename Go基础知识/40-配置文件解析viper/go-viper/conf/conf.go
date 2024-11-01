package conf

import (
	"github.com/spf13/viper"
	"log"
)

var Conf *viper.Viper

// init 每个init函数在整个Go程序生命周期内仅会被执行一次
func init() {
	//fmt.Println("调用conf init()函数")
	Conf = viper.New()
	Conf.SetConfigName("config") // 配置文件名称（不包含扩展名）
	Conf.SetConfigType("yaml")   // 配置文件扩展名
	Conf.AddConfigPath("./conf") // 查找配置文件的路径

	if err := Conf.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}

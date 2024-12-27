package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//Configs 全局配置文件
type Configs struct {
	Configs Config
}

type Config struct {
	Kafka   Kafka
	LogFile []LogFile
}

type Kafka struct {
	Address       string
	GoroutineNums int //开启多少个goroutine去往kafka写数据
}

type LogFile struct {
	FilePath string
	Topic    string
}

//LoadConfigs 读取配置文件函数
func LoadConfigs(configPath string) (Config *Configs) {
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Printf("Load config file failed, err:%v\n", err)
		os.Exit(1)
	}
	//反序列化
	err = json.Unmarshal(file, &Config)
	if err != nil {
		fmt.Println("Unmarshal failed, err:%v\n", err)
		os.Exit(1)
	}
	return Config
}

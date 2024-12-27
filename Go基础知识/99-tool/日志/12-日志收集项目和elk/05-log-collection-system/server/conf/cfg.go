package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//Configs 全局配置
type Configs struct {
	Configs Config
}

type Config struct {
	Kafka Kafka
	Es    Es
}

type Kafka struct {
	Address  string
	ChanSize int //临时存放日志内容的通道的大小
	Topic    []string
}

type Es struct {
	Address string
	Nums    int //开启多少个goroutine往ES写入数据
}

//LoadCofig 读取配置文件
func LoadCofig(cfgPath string) (cfg *Configs) {
	f, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		fmt.Printf("open config file failed, err:%v\n", err)
		os.Exit(1)
	}

	err = json.Unmarshal(f, &cfg)
	if err != nil {
		fmt.Printf("Load config file failed, err:%v\n", err)
		os.Exit(1)
	}

	return cfg
}

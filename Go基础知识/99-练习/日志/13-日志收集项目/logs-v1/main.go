package main

import (
	"logs-v1/conf"
	"logs-v1/kafka"
	"logs-v1/taillog"
)

func main() {
	//0.加载配置文件
	cfg := conf.LoadConfigs("./conf/cfg.json")

	//1.初始化taillog包
	taillog.Init(cfg.Configs.LogFile)

	//2.初始化kafka包
	kafka.Init([]string{cfg.Configs.Kafka.Address}, cfg.Configs.Kafka.GoroutineNums)

	//3.hang住程序
	select {}
}

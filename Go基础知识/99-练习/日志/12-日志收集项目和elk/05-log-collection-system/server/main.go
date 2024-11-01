package main

import (
	"server/conf"
	"server/es"
	"server/kafka"
)

//log_transfer
//将日志从kafka服务中取出，然后存入到elasticsearch数据库中

func main() {
	//0.加载配置文件
	cfg := conf.LoadCofig("./conf/cfg.json")

	//1.初始化ES包
	es.Init(cfg.Configs.Es.Address, cfg.Configs.Es.Nums)

	//2.初始化kafka包
	kafka.Init([]string{cfg.Configs.Kafka.Address}, cfg.Configs.Kafka.Topic, cfg.Configs.Kafka.ChanSize)

	//hang住程序
	select {}
}

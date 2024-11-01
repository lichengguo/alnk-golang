package main

import (
	"04-log-transfer-v2.0/conf"
	"04-log-transfer-v2.0/es"
	"04-log-transfer-v2.0/kafka"
	"fmt"
	ini "gopkg.in/ini.v1"
	"strings"
)

//log transfer
//将日志从kafka服务端取出，然后存入到elasticsearch数据库中

func main() {
	//0.加载配置文件
	var cfg = new(conf.LogTransferCfg)
	err := ini.MapTo(cfg, "./conf/cfg.ini")
	if err != nil {
		fmt.Printf("init config file failed, err:%v\n", err)
		return
	}
	fmt.Println(cfg)

	//1.初始化ES
	//1.1初始化一个es连接的client
	//1.2对外提供一个可以往ES写入数据的函数，其实是一个通道channel
	err = es.Init(cfg.ESCfg.Address, cfg.ESCfg.ChanSize, cfg.ESCfg.Nums)
	if err != nil {
		fmt.Printf("init es server failed, err:%v\n", err)
		return
	}
	fmt.Println("init es server success!")

	//2.初始化kafka
	//2.1初始化一个消费kafka数据的连接的consumer
	//2.2每个分区的消费者分别取出数据 通过SentToESChan发送到es包中的通道
	topic := strings.Split(cfg.KafkaCfg.Topic, "|")
	fmt.Printf("topic:%s\n", topic)

	err = kafka.Init([]string{cfg.KafkaCfg.Address}, topic)
	if err != nil {
		fmt.Printf("init kafka consumer failed, err:%v\n", err)
		return
	}

	//阻塞程序，一直运行
	select {}

}

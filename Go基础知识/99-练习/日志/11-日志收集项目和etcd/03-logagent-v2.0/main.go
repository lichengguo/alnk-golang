package main

import (
	"03-logagent-v2.0/conf"
	"03-logagent-v2.0/etcd"
	"03-logagent-v2.0/kafka"
	"03-logagent-v2.0/taillog"
	"03-logagent-v2.0/utils"
	"fmt"
	"gopkg.in/ini.v1"
	"sync"
	"time"
)

//logagent程序入口

// cfg 指针类型 从配置文件加载kafka，etcd等配置
var cfg = new(conf.AppConf)

// main 入口函数
func main() {
	//0.加载配置文件到cfg这个全局变量
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Printf("load ini config file failed, err:%v\n", err)
		return
	}

	//1.初始化kafka
	err = kafka.Init([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.ChanMaxSize)
	if err != nil {
		fmt.Printf("init Kafka server failed,err:%v\n", err)
		return
	}
	fmt.Println("init kafka server success.")

	//2.初始化etcd
	err = etcd.Init(cfg.EtcdConf.Address, time.Duration(cfg.EtcdConf.Timeout)*time.Second)
	if err != nil {
		fmt.Printf("init etcd failed,err:%v\n", err)
		return
	}
	fmt.Println("init etcd success.")
	//从etcd中获取日志收集项的配置信息
	//先获取本机IP，etcd中根据不同的服务器ip，配置不一样
	ip, err := utils.GetOutboundIP()
	if err != nil {
		fmt.Printf("get local ip failed, err:%v\n", err)
		return
	}
	etcdConfKey := fmt.Sprintf(cfg.EtcdConf.Key, ip) //拼接去拿取ectd中配置的key
	//从etcd中获取对应的IP的配置
	logEntryConf, _ := etcd.GetConf(etcdConfKey)
	fmt.Printf("get config from etcd success, %v\n", logEntryConf)
	for index, value := range logEntryConf {
		fmt.Printf("index:%v value:%v\n", index, value)
	}

	//3. 收集日志发往Kafka
	taillog.Init(logEntryConf)
	// 从taillog包中获取对外暴露的通道
	newConfChan := taillog.NewConfChan()
	//派一个哨兵去监视日志收集项的变化（有变化及时通知我的logAgent实现热加载配置）
	var wg sync.WaitGroup
	wg.Add(1)
	go etcd.WatchConf(etcdConfKey, newConfChan) // 哨兵发现最新的配置信息会通知上面的那个通道
	wg.Wait()
}

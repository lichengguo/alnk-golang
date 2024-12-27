package main

import (
	"02-logagent-v1.0/conf"
	"02-logagent-v1.0/etcd"
	"02-logagent-v1.0/kafka"
	"fmt"
	ini "gopkg.in/ini.v1"
	"time"
)

//logagent程序入口

//声明一个全局变量，用来加载配置文件
var (
	cfg = new(conf.AppConf) //返回的是指针
)

//run 主函数
//func run() {
//	//1.实时读取日志
//	for {
//		select {
//		case line := <-taillog.ReadChan():
//			//2.发送到kafka服务端
//			kafka.SendToKafka(cfg.KafkaConf.Topic, line.Text)
//		default:
//			time.Sleep(time.Second) //当日志文件没有更新的时候就sleep 1秒
//		}
//	}
//}

func main() {
	//0.加载配置文件(高级用法 映射到结构体)
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Printf("load config file failed, err:%v\n", err)
		return
	}
	//fmt.Println(cfg)

	//1.初始化kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Address})
	if err != nil {
		fmt.Printf("connect kafka server failed, err:%v\n", err)
		return
	}
	fmt.Println("connect kafka server success!")
	//kafka.SendToKafka("alnk", "nice...")

	//2.初始化etcd
	//2.1 etcd连接初始化
	err = etcd.Init(cfg.EtcdConf.Address, time.Duration(cfg.EtcdConf.Timeout)*time.Second)
	if err != nil {
		fmt.Printf("init etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("init etcd success!")

	//2.2从etcd中获取日志收集项的配置信息
	logEntryConf, err := etcd.GetConf("/xxx")
	if err != nil {
		fmt.Printf("etcd GetConf failed, err:%v\n", err)
		return
	}
	fmt.Printf("get conf from etcd success, %v\n", logEntryConf)
	for index, value := range logEntryConf {
		fmt.Printf("index:%v value:%v\n", index, value)
	}

	//2.初始化打开日志文件
	//err = taillog.Init(cfg.TaillogConf.FileName)
	//if err != nil {
	//	fmt.Println("Init taillog package failed, err:", err)
	//	return
	//}
	//fmt.Println("init taillog package success!")
	//
	////run()
	//run()
}

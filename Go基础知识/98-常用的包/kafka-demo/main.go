package main

import (
	"fmt"
	"github.com/IBM/sarama"
)

// kafka
// 先看下00notes中的关于kafka的笔记
// sarama包连接kafka，向kafka发送消息
// 基于sarama第三方库开发kafka发送消息的客户端
// kafka启动
// 运行kafka需要使用Zookeeper，所以你需要先启动Zookeeper，如果你没有Zookeeper，你可以使用kafka自带打包和配置好的Zookeeper
// 进入到kafka的目录
// 启动zookeeper
// bin/zookeeper-server-start.sh config/zookeeper.properties
// 启动kafka
// bin/kafka-server-start.sh config/server.properties &

func main() {
	//sarama包的使用
	//1.配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          //发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner //新选出一个 partition
	config.Producer.Return.Successes = true                   //成功交付的消息将在success channel 返回

	//2.构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "web_log"
	msg.Value = sarama.StringEncoder("this is a test log...")

	//3.连接kafka
	client, err := sarama.NewSyncProducer([]string{"192.168.3.121:9092"}, config)
	if err != nil {
		fmt.Printf("connect kafka failed, err:%v\n", err)
		return
	}
	fmt.Println("connect kafka success!")
	defer client.Close() //延迟关闭kafka

	//4.发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Printf("send msg failed, err:%v\n", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
	fmt.Println("send msg to kafka success!")

}

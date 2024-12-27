package kafka

import (
	"fmt"
	"github.com/IBM/sarama"
)

//往kafka服务端写日志模块

// 声明一个全局的连接kafka的生产者client
var (
	client sarama.SyncProducer
)

// Init 初始化kafka client连接
func Init(addr []string) (err error) {
	//利用sarama包连接kafka服务端
	//1.配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          //发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner //新选出一个 partition
	config.Producer.Return.Successes = true                   //成功交付的消息将在success channel 返回

	//2.连接kafka
	client, err = sarama.NewSyncProducer(addr, config)
	if err != nil {
		fmt.Printf("producer closed, err:%v\n", err)
		return
	}
	return
}

// SendToKafka 发送消息到kafka服务端
func SendToKafka(topic, data string) {
	//构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)
	//发送消息到kafka
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
	fmt.Println("send msg to kafka success!")
}

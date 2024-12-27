package kafka

import (
	"fmt"
	"github.com/IBM/sarama"
	"logs-v1/taillog"
	"os"
	"time"
)

var client sarama.SyncProducer //声明一个全局的连接kafka的消息生产者client

func Init(addr []string, nums int) (err error) {
	//1.配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          //发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner //新选出一个 partition
	config.Producer.Return.Successes = true                   //成功交付的消息将在success channel 返回

	//2.连接kafka
	client, err = sarama.NewSyncProducer(addr, config)
	if err != nil {
		fmt.Printf("producer closed, err:%v\n", err)
		os.Exit(1)
	}

	//3.开启goroutine，实时的往kafka服务端写入数据
	//开启多个goroutine
	for i := 0; i < nums; i++ {
		go sendToKafka()
	}
	return
}

// sendToKafka 从通道中读取数据发送到kafka服务器
func sendToKafka() {
	for {
		select {
		case ld := <-taillog.GetInfoToChan(): //从通道中获取数据
			//构造一个消息
			msg := &sarama.ProducerMessage{}
			msg.Topic = ld.Topic
			msg.Value = sarama.StringEncoder(ld.Data)
			//发送消息到kafka
			_, _, err := client.SendMessage(msg)
			if err != nil {
				fmt.Println("send msg failed, err:", err)
				return
			}
			fmt.Printf("send log msg[%s] to kafka success!\n", ld.Data)
		default:
			time.Sleep(time.Millisecond * 500)
		}
	}
}

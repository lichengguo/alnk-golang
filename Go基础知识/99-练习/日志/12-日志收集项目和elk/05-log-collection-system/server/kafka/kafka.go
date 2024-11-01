package kafka

import (
	"fmt"
	"github.com/IBM/sarama"
	"os"
)

// logData 写入到channel通道中的数据格式
type logData struct {
	Topic string
	Data  string
}

// 声明一个全局通道
var ch chan *logData

// Init 初始化kafka
func Init(addr []string, topic []string, chanSize int) (err error) {
	//连接kafka服务器
	consumer, err := sarama.NewConsumer(addr, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		os.Exit(1)
	}
	fmt.Println("start consumer success!")

	//初始化通道的大小
	ch = make(chan *logData, chanSize)

	//消费多个kafka的不同的topic
	for _, topicValue := range topic {
		go consumerKafka(consumer, topicValue)
	}
	return
}

// consumerKafka 消费kafka的消息
func consumerKafka(consumer sarama.Consumer, topic string) {
	//根据topic取到所有的分区
	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		fmt.Printf("fail to get list of partition, err:%v\n", err)
		return
	}
	//一个topic可能有多个分区
	fmt.Println("分区列表:", partitionList)

	//遍历所有的分区
	for partition := range partitionList {
		//针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d, err:%v\n", partition, err)
			return
		}
		//异步从每个分区消费消息
		go func(pc sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Topic:%v Value:%v\n", msg.Partition, msg.Offset, msg.Key, msg.Topic, string(msg.Value))
				ld := logData{
					Topic: topic,
					Data:  string(msg.Value),
				}
				ch <- &ld //发送日志数据到通道
			}
		}(pc)
	}
}

// GetLogInfoToChan 暴露一个通道给外部使用
func GetLogInfoToChan() <-chan *logData {
	return ch
}

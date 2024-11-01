package main

import (
	"fmt"
	"github.com/IBM/sarama"
)

//从kafka服务端消费消息

func main() {
	//消费端连接kafka服务器
	consumer, err := sarama.NewConsumer([]string{"192.168.3.121:9092"}, nil)
	if err != nil {
		fmt.Printf("fail to start kafka consumer, err:%v\n", err)
		return
	}

	//根据topic取到所有的分区，topic下的数据会存在多个分区内
	partitonList, err := consumer.Partitions("web_log")
	if err != nil {
		fmt.Printf("fail to get list of partition, err:%v\n", err)
		return
	}
	fmt.Printf("分区列表:%v\n", partitonList)

	//循环遍历所有的分区
	for partition := range partitonList {
		pc, err := consumer.ConsumePartition("web_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d, err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		//异步从每个分区消费消息
		go func(pc sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Topic:%v Value:%v\n", msg.Partition, msg.Offset, msg.Key, msg.Topic, string(msg.Value))
			}
		}(pc)
		select {}
	}
}

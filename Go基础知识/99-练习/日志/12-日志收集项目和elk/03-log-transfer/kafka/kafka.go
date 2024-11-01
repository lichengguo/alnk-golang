package kafka

import (
	"03-log-transfer/es"
	"fmt"
	"github.com/IBM/sarama"
)

// Init 初始化kafka
func Init(addr []string, topic string) (err error) {
	//连接kafka服务器
	consumer, err := sarama.NewConsumer(addr, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return err
	}
	//根据topic取到所有的分区
	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		fmt.Printf("fail to get list of partition, err:%v\n", err)
		return err
	}
	fmt.Println("分区列表:", partitionList)

	//遍历所有的分区
	for partition := range partitionList {
		//针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d, err:%v\n", partition, err)
			return err
		}
		//defer pc.Close()

		//异步从每个分区消费消息
		go func(pc sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Topic:%v Value:%v\n", msg.Partition, msg.Offset, msg.Key, msg.Topic, string(msg.Value))
				ld := es.LogData{
					Topic: topic,
					Data:  string(msg.Value),
				}
				es.SentToESChan(&ld) //发送到es包中的通道
			}
		}(pc)
	}
	return
}

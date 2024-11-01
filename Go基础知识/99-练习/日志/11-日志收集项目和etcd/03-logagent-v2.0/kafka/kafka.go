package kafka

import (
	"fmt"
	"github.com/IBM/sarama"
	"time"
)

// logData 从日志文件收集日志存在通道中
type logData struct {
	topic string //kafka服务端存放的topic
	data  string //kafka服务端存放的日志内容
}

var (
	client      sarama.SyncProducer //声明一个全局的kafka client的连接
	logDataChan chan *logData
)

// Init 初始化kafka的client连接
func Init(addr []string, maxSize int) (err error) {
	//利用sarama第三方包连接kafka服务器
	//1.配置设置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出⼀个 partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	//2.连接kafka服务端
	client, err = sarama.NewSyncProducer(addr, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}

	//初始化logDataChan通道，用来给taillog包存放收集到日志
	//然后kafka包直接从logDataChan通道读取日志消息，发送到kafka服务端
	logDataChan = make(chan *logData, maxSize)

	//开启后台的goroutine，一直从通道中取取数据发往kafka服务器
	go sendToKafka()

	return
}

// sendToKafka 从logDataChan通道中取值发送到kafka服务器
func sendToKafka() {
	for {
		select {
		case ld := <-logDataChan:
			//构造⼀个消息
			msg := &sarama.ProducerMessage{}
			msg.Topic = ld.topic
			msg.Value = sarama.StringEncoder(ld.data)
			//发送到kafka
			pid, offset, err := client.SendMessage(msg)
			if err != nil {
				fmt.Println("send msg to kafka failed, err:", err)
				return
			}
			fmt.Printf("pid:%v offset:%v\n", pid, offset)
		default:
			time.Sleep(time.Millisecond * 50)
		}
	}
}

// SendToChan 给外部的包暴露一个可以往logDataChan通道中存放日志的函数
// topic 发送到kafka服务器的topic
// data 发送到kafka服务器的日志内容
func SendToChan(topic, data string) {
	msg := &logData{
		topic: topic,
		data:  data,
	}
	logDataChan <- msg
}

package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"strings"
	"time"
)

var (
	client *elastic.Client //声明一个全局的es数据库的客户端连接
	ch     chan *LogData
)

//LogData 写入到channel通道中的数据格式
type LogData struct {
	Topic string `json:"topic"`
	Data  string `json:"data"`
}

//Init 初始化ES连接
func Init(address string, chanSize, nums int) (err error) {
	//strings.HasPrefix 以什么开头
	if !strings.HasPrefix(address, "http://") {
		address = "http://" + address
	}

	//连接es数据库
	client, err = elastic.NewClient(elastic.SetURL(address))
	if err != nil {
		return err
	}
	fmt.Println("connect to es success!")

	//启动16个goroutine往es数据库写入数据
	ch = make(chan *LogData, chanSize)
	for i := 0; i < nums; i++ {
		go sendToES()
	}
	return
}

//sendToES 从通道ch中读取数据，写入到es数据库
func sendToES() {
	for {
		select {
		case msg := <-ch:
			put1, err := client.Index().Index(msg.Topic).BodyJson(msg).Do(context.Background())
			if err != nil {
				fmt.Printf("write to es server failed, err:%v\n", err)
				continue
			}
			fmt.Printf("_index:%v _id:%v\n", put1.Index, put1.Id)
		default:
			time.Sleep(time.Second * 1)
		}
	}
}

//sentToESChan 暴露一个函数，给其他模块使用
//把数据存入到通道中
func SentToESChan(msg *LogData) {
	ch <- msg
}

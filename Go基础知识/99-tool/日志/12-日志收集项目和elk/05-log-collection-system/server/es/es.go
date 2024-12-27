package es

import (
	"context"
	"fmt"
	elastic "github.com/olivere/elastic/v7"
	"os"
	"server/kafka"
	"strings"
	"time"
)

var client *elastic.Client //声明一个全局的es数据库的客户端连接

// Init 初始化ES包
func Init(address string, nums int) (err error) {
	//strings.HasPrefix 以什么开头
	if !strings.HasPrefix(address, "http://") {
		address = "http://" + address
	}

	//连接es数据库
	client, err = elastic.NewClient(elastic.SetURL(address))
	if err != nil {
		fmt.Printf("connect to es server failed, err:%v\n", err)
		os.Exit(1)
	}
	fmt.Println("connect to es server success!")

	//启动多个个goroutine往es数据库写入数据
	for i := 0; i < nums; i++ {
		go sendToES()
	}

	return
}

// sendToES 从kafka包中的GetLogInfoToChan函数获取数据，写入到es数据库
func sendToES() {
	for {
		select {
		case msg := <-kafka.GetLogInfoToChan():
			//写入es数据库
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

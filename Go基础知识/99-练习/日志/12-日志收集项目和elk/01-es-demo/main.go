package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

//go语言操作elasticsearch数据库
//官网文档: https://godoc.org/github.com/olivere/elastic
//我们使用第三方库 https://github.com/olivere/elastic 来连接ES并进行操作
//注意下载与你的ES相同版本的client，
//例如我们这里使用的ES是7.2.1的版本，那么我们下载的client也要与之对应为github.com/olivere/elastic/v7

//Student 需要插入到es数据库中的数据结构
type Student struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

func main() {
	//1.连接es服务器
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		// Handle error
		fmt.Printf("connect to es server failed, err:%v\n", err)
		return
	}
	fmt.Printf("connect to es server success!\n")

	//2.插入数据
	p1 := Student{
		Name:    "tangbohu444",
		Age:     1000,
		Married: false,
	}
	put1, err := client.Index().Index("japan").Type("student").BodyJson(p1).Do(context.Background())
	if err != nil {
		fmt.Printf("insert msg to es server failed, err:%v\n", err)
		return
	}
	fmt.Printf("数据插入成功, _id:%s _index:%s, _type:%s\n", put1.Id, put1.Index, put1.Type)
}

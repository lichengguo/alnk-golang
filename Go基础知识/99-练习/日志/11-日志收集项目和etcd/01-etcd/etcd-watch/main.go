package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

//etcd watch
//实时的，用来获取更改的通知

func main() {
	//1.连接etcd服务器
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.3.121:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd server sucess!")
	defer cli.Close() //延迟关闭连接

	//2.watch
	//派一个类似redis哨兵一样的东西，一直监视着某个key的变化（例如：新增，修改，删除）
	ch := cli.Watch(context.Background(), "/xxx") //返回一个通道
	//从通道中尝试取值（监视的信息）
	for wresp := range ch {
		for _, evt := range wresp.Events {
			fmt.Printf("Type:%v key:%v value:%v\n", evt.Type, string(evt.Kv.Key), string(evt.Kv.Value))
		}
	}

}

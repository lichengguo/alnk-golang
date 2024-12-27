package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

//从etcd服务端拉取数据
//etcd get demo

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

	//2.get
	//从etcd服务端获取数据

	//2.1设置获取消息的超时时间
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//2.2获取消息
	resp, err := cli.Get(ctx, "/logagent/192.168.1.6/collect_config")
	cancel() //超时自动关闭释放资源
	if err != nil {
		fmt.Printf("get msg from etcd server failed, err:%v\n", err)
		return
	}
	//2.3读取消息内容输出到终端
	for _, ev := range resp.Kvs {
		fmt.Printf("key:%s value:%s\n", ev.Key, ev.Value)
	}

}

package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

//发送消息到etcd服务器
//put msg to etcd
//client etcd/clientv3
/*
1.etcd服务端服务的安装
直接看下面这个链接的安装介绍
https://github.com/etcd-io/etcd/releases/tag/v3.4.13
注意：启动时修改监听的ip，默认监听127.0.0.1
命令:
./etcd --listen-client-urls http://0.0.0.0:2379 -advertise-client-urls http://0.0.0.0:2379 --listen-peer-urls http://0.0.0.0:2380

2.安装 etcd/clientv3 第三方扩展失败
https://www.icode9.com/content-4-689215.html
解决办法:
直接在go.mod文件中添加
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
然后重新 go get go.etcd.io/etcd/clientv3
如果还有报错就
replace github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.4
然后在go mod tidy 一下

*/

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

	//2.发送消息到etcd服务器
	//2.1设置发送消息的超时时间
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//2.2value值
	//value := `[{"path":"./nginx.log","topic":"web_log"},{"path":"./redis.log","topic":"redis_log"},{"path":"./mysql.log","topic":"mysql_log"}]`
	value := `[{"path":"./testlog/nginx.log","topic":"web_log"},{"path":"./testlog/mysql.log","topic":"mysql_log"}]`
	//2.3发送消息
	_, err = cli.Put(ctx, "/logagent/192.168.1.6/collect_config", value)
	cancel() //超时自动关闭释放资源
	if err != nil {
		fmt.Printf("put msg to etcd server falied, err:%v\n", err)
		return
	}

	//3.收集日志发往kafka

}

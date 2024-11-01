package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

//客户端

//
var wg sync.WaitGroup

//定义一个消息接收通道
var respChan = make(chan *respData, 1)

type respData struct {
	resp *http.Response
	err  error
}

//客户端超时取消示例
func doCall(ctx context.Context) {
	// 组装client请求头
	transport := http.Transport{
		//请求频繁可定义全局的client对象并启用长连接
		//请求不频繁使用短连接
		DisableKeepAlives: true, //这里是使用短连接
	}
	client := http.Client{
		Transport: &transport,
	}

	// 新创建一个GET请求
	req, err := http.NewRequest("GET", "http://127.0.0.1:8000", nil)
	if err != nil {
		fmt.Printf("new request failed, err:%v\n", err)
		return
	}
	// 使用带有超时的ctx创建一个新的请求
	req = req.WithContext(ctx)

	// 启动一个goroutine去连接服务器
	wg.Add(1)
	go func() {
		resp, err := client.Do(req) //向服务器发送请求
		if err != nil {
			fmt.Printf("client.do resp:%v, err:%v\n", resp, err)
		}
		rd := &respData{
			resp: resp,
			err:  err,
		}
		respChan <- rd
		wg.Done()
	}()
	defer wg.Wait()

	select {
	case <-ctx.Done():
		fmt.Println("call api timeout")
	case result := <-respChan:
		fmt.Println("call server api success")
		if result.err != nil {
			fmt.Printf("call server api failed, err:%v\n", err)
			return
		}
		defer result.resp.Body.Close()
		data, _ := ioutil.ReadAll(result.resp.Body)
		fmt.Printf("resp:%v\n", string(data))
	}
}

func main() {
	// 定义一个100毫秒的超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	// 调用cancel，释放goroutine资源
	defer cancel()

	// 调用请求函数
	doCall(ctx)
}

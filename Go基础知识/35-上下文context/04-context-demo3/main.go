package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 如何优雅的控制子goroutine退出？
// context方式 官方版的方案

var wg sync.WaitGroup

func f1(ctx context.Context) {
	defer wg.Done()
	go f2(ctx)
FORLOOP:
	for {
		fmt.Println("f1...")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): //等待上级通知
			break FORLOOP
		default:
		}
	}
}

func f2(ctx context.Context) {
FORLOOP:
	for {
		fmt.Println("f2...")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			break FORLOOP
		default:
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go f1(ctx)
	time.Sleep(time.Second * 5)
	//如何通知子goroutine退出
	cancel() //通知子goroutine结束
	wg.Wait()
	fmt.Println("over...")
}

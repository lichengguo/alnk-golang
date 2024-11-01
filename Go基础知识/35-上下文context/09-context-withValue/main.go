package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// context with系列方法
// WithValue

// func WithValue(parent Context, key, val interface{}) Context
// WithValue 返回父节点的副本，其中与key关联的值为val

// 仅对API和进程间传递请求域的数据使用上下文值，而不是使用它来传递可选参数给函数。

// 所提供的键必须是可比较的，并且不应该是string类型或任何其他内置类型，以避免使用上下文在包之间发生冲突。
// WithValue的用户应该为键定义自己的类型。
// 为了避免在分配给interface{}时进行分配，上下文键通常是具有具体类型struct{}。
// 或者导出的上下文关键变量的静态类型应该是指针或接口

type TraceCode string

var wg sync.WaitGroup

func worker(ctx context.Context) {
	key := TraceCode("TRACE_CODE")
	traceCode, ok := ctx.Value(key).(string)
	if !ok {
		fmt.Println("invalid trace code")
	}

LOOP:
	for {
		fmt.Printf("worker log, trace code:%s\n", traceCode)
		time.Sleep(time.Millisecond * 10)
		select {
		case <-ctx.Done(): //50毫秒后自动调用，结束这个goroutine
			break LOOP
		default:
		}
	}
	fmt.Println("worker done!")
	wg.Done()
}

func main() {
	// 设置一个50毫秒的超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)

	// 在系统的入口设置trace code 传递给后续启动的goroutine 实现日志数据聚合
	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "12345678")

	wg.Add(1)
	go worker(ctx)

	time.Sleep(5 * time.Second)
	cancel()

	wg.Wait()
	fmt.Println("main over!")
}

// 推荐以参数的方式显示传递Context
// 以Context作为参数的函数方法，应该把Context作为第一个参数。
// 给一个函数方法传递Context的时候，不要传递nil，如果不知道传递什么，就使用context.TODO()
// Context的Value相关方法应该传递请求域的必要数据，不应该用于传递可选参数
// Context是线程安全的，可以放心的在多个goroutine中传递

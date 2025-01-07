#### Go的context使用

##### 为什么需要context

```go
/*
在Go http包的Server中，每一个请求在都有一个对应的goroutine去处理
请求处理函数通常会启动额外的goroutine用来访问后端服务，比如数据库和RPC服务
用来处理一个请求的goroutine通常需要访问一些与请求特定的数据，比如终端用户的身份认证信息、验证相关的token、请求的截止时间
当一个请求被取消或超时时，所有用来处理该请求的goroutine都应该迅速退出，然后系统才能释放这些goroutine占用的资源

*/

//为什么需要Context
//1-基本示例
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

//初始的例子

func worker() {
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
	}

	//退出
	wg.Done()
}

func main() {
	wg.Add(1)
	go worker()

	//如何优雅的实现结束子goroutine
	wg.Wait()

	fmt.Println("over")
}

```



```go
//2-全局变量方式
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var exit bool

//全局变量方式存在的问题：
//1. 使用全局变量在跨包调用时不容易统一
//2. 如果worker中再启动goroutine，就不太好控制了

func worker() {
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		if exit {
			break
		}
	}

	wg.Done()
}

func main() {
	wg.Add(1)
	go worker()
	time.Sleep(time.Second * 3) //sleep3秒避免程序过快退出

	exit = true //修改全局变量实现子goroutine退出
	wg.Wait()
	fmt.Println("over")
}

```



```go
//3-通道方式
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

//管道方式存在的问题：
//1. 使用全局变量在跨包调用时不容易实现规范和统一，需要维护一个共用的channel

func worker(exitCahn chan struct{}) {
LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)

		select {
		case <-exitCahn: //等待接收上级通知
			break LOOP
		default:

		}
	}

	wg.Done()
}

func main() {
	var exitChan = make(chan struct{})
	wg.Add(1)
	go worker(exitChan)

	time.Sleep(time.Second * 3) //sleep3秒以免程序过快退出

	exitChan <- struct{}{} //给子goroutine发送退出信号
	close(exitChan)

	wg.Wait()

	fmt.Println("over")
}

```



```go
//4-官方版的方案
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(ctx context.Context) {
LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): //等待上级通知
			break LOOP
		default:
		}
	}

	wg.Done()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 3)

	cancel() //通知子goroutine结束

	wg.Wait()

	fmt.Println("over")
}

```



```go
//5-当子goroutine又开启另外一个goroutine时，只需要将ctx传入即可
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(ctx context.Context) {
	go worker2(ctx)

LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): //等待上级通知
			break LOOP
		default:
		}
	}

	wg.Done()
}

func worker2(ctx context.Context) {
LOOP:
	for {
		fmt.Println("worker2")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): //等待上级通知
			break LOOP
		default:
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 3)

	cancel() //通知子goroutine结束

	wg.Wait()

	fmt.Println("over")
}

```



##### 常用场景

```go
//context常用场景
//1-并发控制子goroutine退出
//2-超时
//3-传递上下文

//1-控制子goroutine退出
package main

import (
	"context"
	"fmt"
	"time"
)

func test1(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("收到信号，test1结束")
			return
		default:
			fmt.Println("test1...")     //执行业务逻辑
			time.Sleep(1 * time.Second) //不让打印速度太快
		}
	}
}

func test2(ctx context.Context) {
	go test3(ctx)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("收到信号，test2结束")
			return
		default:
			fmt.Println("test2...")     //执行业务逻辑
			time.Sleep(1 * time.Second) //不让打印速度太快
		}
	}
}

func test3(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("收到信号，test3结束")
			return
		default:
			fmt.Println("test3...")     //执行业务逻辑
			time.Sleep(1 * time.Second) //不让打印速度太快
		}
	}
}

func main() {
	//父context(利用根context得到)
	ctx, cancel := context.WithCancel(context.Background())

	//父context的子协程
	go test1(ctx)
	go test2(ctx)

	time.Sleep(5 * time.Second) //让子协程执行5秒

	fmt.Println("主进程结束了，通知子goroutine结束")
	cancel()                    //取消的信号，结束test1、test2和test2的子协程test3函数的运行
	time.Sleep(2 * time.Second) //这里等待2秒是为了让协程的退出信息打印出来
}

```

```go
//context常用场景
//1-并发控制子goroutine退出
//2-超时
//3-传递上下文

//1-控制子goroutine退出
package main

import (
	"context"
	"fmt"
	"time"
)

func test1(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("收到信号，test1结束")
			return
		default:
			fmt.Println("test1...")     //执行业务逻辑
			time.Sleep(1 * time.Second) //不让打印速度太快
		}
	}
}

func test2(ctx context.Context) {
	go test3(ctx)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("收到信号，test2结束")
			return
		default:
			fmt.Println("test2...")     //执行业务逻辑
			time.Sleep(1 * time.Second) //不让打印速度太快
		}
	}
}

func test3(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("收到信号，test3结束")
			return
		default:
			fmt.Println("test3...")     //执行业务逻辑
			time.Sleep(1 * time.Second) //不让打印速度太快
		}
	}
}

func main() {
	//父context(利用根context得到)
	ctx, cancel := context.WithCancel(context.Background())

	//父context的子协程
	go test1(ctx)

	//=======================================================================
	//子context，注意：这里虽然也返回了cancel的函数对象，但是未使用
	ctx2, _ := context.WithCancel(ctx)
	//子context的子协程
	go test2(ctx2)
	//=======================================================================

	time.Sleep(5 * time.Second) //让子协程执行5秒

	fmt.Println("主进程结束了，通知子goroutine结束")
	cancel()                    //取消的信号，结束test1、test2和test2的子协程test3函数的运行
	time.Sleep(2 * time.Second) //这里等待2秒是为了让协程的退出信息打印出来
}

```



```go
//context常用场景
//1-并发控制子goroutine退出
//2-超时
//3-传递上下文

//2-超时
package main

import (
	"context"
	"fmt"
	"time"
)

// NewContextWithTimeOut 创建一个带超时context， 三秒后退出执行
func NewContextWithTimeOut() (context.Context, context.CancelFunc) {
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	//return ctx, cancel

	return context.WithTimeout(context.Background(), time.Second*3)
}

//HttpHandler 处理程序
func HttpHandler() {
	ctx, cancel := NewContextWithTimeOut()
	defer cancel()
	Create(ctx)

}

//Create 业务逻辑代码
func Create(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("超时了，主进程叫我退出了...")
			fmt.Println(ctx.Err())
			return
		default:
			time.Sleep(time.Millisecond * 500)
			fmt.Println("正在努力的处理业务逻辑中......")
		}
	}
}

func main() {
	HttpHandler()
}

```



```go
//context常用场景
//1-并发控制子goroutine退出
//2-超时
//3-传递上下文

//3-传递上下文
package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"strings"
	"time"
)

type MyKEY string

const KEY MyKEY = "trace_id"

//NewRequestID 返回一个UUID
func NewRequestID() MyKEY {
	return MyKEY(strings.Replace(uuid.New().String(), "-", "", -1))
}

//NewContextWithTraceID 创建一个携带trace_id 的ctx
func NewContextWithTraceID() context.Context {
	return context.WithValue(context.Background(), KEY, NewRequestID())
}

//GetContextValue 获取设置的key对应的值,并断言
func GetContextValue(ctx context.Context, k MyKEY) MyKEY {
	if v, ok := ctx.Value(k).(MyKEY); ok {
		fmt.Println("打印k: " + k)
		fmt.Printf("打印v: %v\n", v)
		return v
	}

	return ""
}

//PrintLog 打印
func PrintLog(ctx context.Context, message string) {
	fmt.Printf("%s|info|trace_id=%s|%s\n", time.Now().Format("2006-01-02 15:04:05"), GetContextValue(ctx, KEY), message)
}

//ProcessEnter 打印值
func ProcessEnter(ctx context.Context) {
	PrintLog(ctx, "Golang")
}

func main() {
	ProcessEnter(NewContextWithTraceID())
}

```

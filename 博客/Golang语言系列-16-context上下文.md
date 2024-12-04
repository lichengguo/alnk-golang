# context上下文

## 控制子goroutine退出

### 全局变量方式

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// 如何优雅的控制子goroutine退出？
// 全局变量方式

// 存在的问题：
// 1.使用全局变量在跨包调用时不容易统一
// 2.如果f中再启动goroutine，就不太好控制了
var wg sync.WaitGroup
var notify bool // 标志位，控制子goroutine退出

func f() {
	defer wg.Done()

	for {
		fmt.Println("fffffffff")
		time.Sleep(time.Second)
		if notify {
			break
		}
	}
}

func main() {
	wg.Add(1)
	go f()
	time.Sleep(time.Second * 5) //sleep 5秒以免程序过快退出
	// 如何通知子goroutine退出
	notify = true //修改全局变量实现子goroutine的退出
	wg.Wait()
}
```



### 管道方式

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// 如何优雅的控制子goroutine退出？
// 通道方式

// 管道方式存在的问题：
// 1. 使用全局变量在跨包调用时不容易实现规范和统一，需要维护一个共用的channel

var wg sync.WaitGroup
var exitChan = make(chan bool, 1) // 退出子goroutine标志位

func f() {
	defer wg.Done()

FORLOOP:
	for {
		fmt.Println("ffffff")
		time.Sleep(time.Second * 1)
		select {
		case <-exitChan: //等待接收上级通知
			break FORLOOP
		default:
		}
	}
}

func main() {
	wg.Add(1)
	go f()
	time.Sleep(time.Second * 5) //sleep 5秒以免程序过快退出
	//如何通知子goroutine退出
	exitChan <- true //给子goroutine发送退出信号
	close(exitChan)
	wg.Wait()
	fmt.Println("over")
}
```



### context方式

```go
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
```



## context接口

```go
context.Context是一个接口，该接口定义了四个需要实现的方法。具体签名如下：

type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}

其中：
    Deadline方法需要返回当前Context被取消的时间，也就是完成工作的截止时间（deadline）；

    Done方法需要返回一个Channel，这个Channel会在当前工作完成或者上下文被取消之后关闭，多次调用Done方法会返回同一个Channel；

    Err方法会返回当前Context结束的原因，它只会在Done返回的Channel被关闭时才会返回非空的值；
        如果当前Context被取消就会返回Canceled错误；
        如果当前Context超时就会返回Dead line Exceeded错误；

    Value方法会从Context中返回键对应的值，对于同一个上下文来说，多次调用Value 并传入相同的Key会返回相同的结果，该方法仅用于传递跨API和进程间跟请求域的数据；



Background()和TODO()
    Go内置两个函数：Background()和TODO()，这两个函数分别返回一个实现了Context接口的background和todo。
    我们代码中最开始都是以这两个内置的上下文对象作为最顶层的partent context，衍生出更多的子上下文对象

    Background()主要用于main函数、初始化以及测试代码中，作为Context这个树结构的最顶层的Context，也就是根Context

    TODO()它目前还不知道具体的使用场景，如果我们不知道该使用什么Context的时候，可以使用这个。

    background和todo本质上都是emptyCtx结构体类型，是一个不可取消，没有设置截止时间，没有携带任何值的Context

```



## WithCancel

```go
package main

import (
	"context"
	"fmt"
)

// context with系列方法
// WithCancel
// func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
// WithCancel返回带有新Done通道的父节点的副本。
// 当调用返回的cancel函数或当关闭父上下文的Done通道时，将关闭返回上下文的Done通道，无论先发生什么情况

// 取消此上下文将释放与其关联的资源，因此代码应该在此上下文中运行的操作完成后立即调用cancel

// gen ...
func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return //return结束该goroutine防止泄露
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel() //当我们取完需要的整数后调用cancel

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}

}

// 上面的示例代码中，gen函数在单独的goroutine中生成整数并将它们发送到返回的通道。
// gen的调用者在使用生成的整数之后需要取消上下文，以免gen启动的内部goroutine发生泄漏
```



## WithDeadline

```go
package main

import (
	"context"
	"fmt"
	"time"
)

// context with系列方法

// WithDeadline
// func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)

// 返回父上下文的副本，并将deadline调整为不迟于d。
// 如果父上下文的deadline已经早于d，则WithDeadline(parent, d)在语义上等同于父上下文。
// 当截止日过期时，当调用返回的cancel函数时，或者当父上下文的Done通道关闭时，返回上下文的Done通道将被关闭，以最先发生的情况为准。

// 取消此上下文将释放与其关联的资源，因此代码应该在此上下文中运行的操作完成后立即调用cancel

func main() {
	d := time.Now().Add(time.Millisecond * 50) //当前时间往后50毫秒
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// 尽管ctx会过期，但在任何情况下调用它的cancel函数都是很好的
	// 如果不这样做，可能会使上下文及其父类存活时间超过必要的时间，造成资源的浪费
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println("err:", ctx.Err())
	}
}

// 上面的代码中，定义了一个50毫秒之后过期的deadline
// 然后我们调用context.WithDeadline(context.Background(), d)得到一个上下文（ctx）和一个取消函数（cancel）
// 然后使用一个select让主程序陷入等待：
// 等待1秒后打印overslept退出 或者 等待ctx过期后退出。
// 因为ctx 50毫秒后就过期，所以ctx.Done()会先接收到值，上面的代码会打印ctx.Err()取消原因
```



## WithTimeout

```go
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// context with系列方法

// WithTimeout
// func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)

// 取消此上下文将释放与其相关的资源，
// 因此代码应该在此上下文中运行的操作完成后立即调用cancel，
// 通常用于数据库或者网络连接的超时控制

var wg sync.WaitGroup

func worker(ctx context.Context) {
LOOP:
	for {
		fmt.Println("db connecting...")
		time.Sleep(time.Millisecond * 10) //假设正常连接数据库耗时10毫秒
		select {
		case <-ctx.Done(): //50毫秒后调用
			break LOOP
		default:
		}
	}
	fmt.Println("work done!")
	wg.Done()
}

func main() {
	// 设置一个50毫秒的超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)

	wg.Add(1)
	go worker(ctx)

	time.Sleep(time.Second * 5)

	cancel() //调用cancel，释放上下文资源

	wg.Wait()

	fmt.Println("over!")
}
```



## WithValue

```go
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
```



## 练习 模拟http请求超时

Server/main.go

```go
package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// server端，随机出现慢响应

func indexHandler(w http.ResponseWriter, r *http.Request) {
	number := rand.Intn(2) //产生0和1的随机整数

	if number == 0 {
		time.Sleep(time.Second * 10) //耗时10秒的慢响应
		fmt.Fprintf(w, "slow response")
		return
	}

	fmt.Fprintf(w, "quick response") //正常响应
}

func main() {
	http.HandleFunc("/", indexHandler)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
```



client/main.go

```go
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

	// 使用带有超时的ctx创建一个新的client request
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
```

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
使用goroutine和channel实现一个计算int64随机数 各个位数和 的程序。
	1. 开启一个goroutine循环生成int64类型的随机数，发送到jobChan
	2. 开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
	3. 主goroutine从resultChan取出结果并打印到终端输出
*/

// job ...
type job struct {
	value int64
}

// result ...
type result struct {
	job *job
	sum int64
}

// 这个例子中，因为一直死循环接收通道中的值，所以无缓冲区也可以。不过建议还是要设置合理的缓冲区
var jobChan = make(chan *job)
var resultChan = make(chan *result)
var wg sync.WaitGroup

// 生产随机数
func zhoulin(zl chan<- *job) { //接收一个通道，这个通道的类型是 *job
	defer wg.Done()
	// 循环生成int64类型的随机数，发送到jobChan
	for {
		x := rand.Int63()        //获取一个随机数 类型为int64
		newJob := &job{value: x} //把job类型的结构体的指针值给newjob
		zl <- newJob             //把newjob这个指针存放到zl这个同道中人
		time.Sleep(time.Millisecond * 500)
	}
}

// 消费随机数
func baodelu(zl <-chan *job, resultChan chan<- *result) {
	defer wg.Done()
	// 从jobChan中取出随机数计算 个位数 的和，将结果发送到resultChan通道
	for {
		job := <-zl //取出来的是一个 *job结构体
		sum := int64(0)
		n := job.value
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		newResult := &result{
			job: job,
			sum: sum,
		}
		resultChan <- newResult //把结果存入到通道
	}
}

func main() {
	// 开启一个goroutine，生产随机数
	wg.Add(1)
	go zhoulin(jobChan) //jobChan是一个通道，通道的值类型是 *job

	// 开启24个goroutine从通道中取值
	wg.Add(24)
	for i := 0; i < 24; i++ {
		go baodelu(jobChan, resultChan)
	}

	// 主goroutine从resulteChan取出结果并打印到终端
	for result := range resultChan {
		fmt.Printf("value:%d sum:%d\n", result.job.value, result.sum)
	}

	wg.Wait()
}

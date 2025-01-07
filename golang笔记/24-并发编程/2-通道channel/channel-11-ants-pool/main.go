package main

// 协程池
// 防止项目滥用协程，导致系统资源耗尽

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"

	"github.com/panjf2000/ants/v2"
)

const totalTask = 10

func main() {
	fmt.Println("start to work. GOMAXPROCS:", runtime.GOMAXPROCS(0))

	ant, _ := ants.NewPool(3) // 创建一个协程池，最大容量为3
	defer ant.Release()       // 释放协程池

	var wg sync.WaitGroup
	results := make([]int, totalTask)
	for i := 0; i < totalTask; i++ {
		idx := i   // 闭包问题 不能在闭包中直接使用循环变量 必须复制出一份
		input := i // 闭包问题 不能在闭包中直接使用循环变量 必须复制出一份
		wg.Add(1)
		err := ant.Submit(func() { // 提交任务 池满则阻塞
			defer wg.Done()
			// TODO 任务逻辑
			rd := rand.New(rand.NewSource(time.Now().UnixNano()))
			time.Sleep(time.Duration(rd.Intn(10)) * time.Millisecond)
			fmt.Println("task index:", input)
			results[idx] = input * input
		})
		if err != nil {
			fmt.Println("faild to submit task, err: ", err)
			wg.Done() // 任务提交失败，防止wg.Wait()一直阻塞
			break     // 任务提交失败，直接退出
		}
	}

	wg.Wait()

	fmt.Println("done, results: ", results)

}

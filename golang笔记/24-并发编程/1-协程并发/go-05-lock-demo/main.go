package main

import (
	"fmt"
	"sync"
)

/*
[并发安全和锁]

[互斥锁 sync.Mutex]
互斥锁是完全互斥的

为什么需要锁？
有时候在Go代码中可能会存在多个goroutine同时操作一个资源（临界区），这种情况会发生竞态问题（数据竞态）
类比现实生活中的例子有十字路口被各个方向的的汽车竞争；还有火车上的卫生间被车厢里的人竞争

使用互斥锁能够保证同一时间有且只有一个goroutine进入临界区，其他的goroutine则在等待锁
当互斥锁释放后，等待的goroutine才可以获取锁进入临界区，
多个goroutine同时等待一个锁时，唤醒的策略是随机的
*/

var x int64
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	for i := 0; i < 50000; i++ {
		// 有锁的情况
		lock.Lock() // 加锁
		x = x + 1
		lock.Unlock() // 解锁

		// 没锁的情况
		// x = x + 1
	}
	wg.Done()
}
func main() {
	// 启用2个goroutine计算
	wg.Add(2)
	go add()
	go add()
	wg.Wait()

	// 打印结果
	fmt.Println(x)
}

/*
不加锁多次运算的结果：51630 51769 51284

加锁多次运算的结果：100000 100000 100000
*/

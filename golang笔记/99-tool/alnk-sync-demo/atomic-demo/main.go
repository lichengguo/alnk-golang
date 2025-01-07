package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/*
[atomic 原子操作]
代码中的加锁操作因为涉及内核态的上下文切换会比较耗时、代价比较高
针对基本数据类型我们还可以使用原子操作来保证并发安全
因为原子操作是Go语言提供的方法它在用户态就可以完成，因此性能比加锁操作更好
Go语言中原子操作由内置的标准库sync/atomic提供

其他的原子操作的方法见图

这个例子没看出原子操作比加锁的性能好呀？？？
*/

var (
	x    int64
	wg   sync.WaitGroup
	lock sync.Mutex
)

// addAtomic 原子操作
func addAtomic() {
	defer wg.Done()
	//原子操作
	atomic.AddInt64(&x, 1)
}

// addLock 加锁操作
func addLock() {
	defer wg.Done()
	//互斥锁
	lock.Lock()
	x++
	lock.Unlock()
}

func main() {
	//原子操作执行耗时
	start := time.Now()
	for i := 0; i < 10000000; i++ {
		wg.Add(1)
		go addAtomic()
	}
	wg.Wait()
	fmt.Println(x)
	fmt.Println(time.Now().Sub(start)) //3.185264818s

	//加锁操作执行耗时
	//start1 := time.Now()
	//for i := 0; i < 10000000; i++ {
	//	wg.Add(1)
	//	go addLock()
	//}
	//wg.Wait()
	//fmt.Println(x)
	//fmt.Println(time.Now().Sub(start1)) //3.14448854s
}

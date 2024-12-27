package main

import (
	"fmt"
	"sync"
	"time"
)

/*
[sync.RWMutex 读写互斥锁]
互斥锁是完全互斥的，但是有很多实际的场景下是读多写少的
当我们并发的去读取一个资源不涉及资源修改的时候是没有必要加锁的，这种场景下使用读写锁是更好的一种选择
读写锁在Go语言中使用sync包中的RWMutex类型

读写锁分为两种：读锁和写锁
当一个goroutine获取读锁之后，其他的goroutine如果是获取读锁会继续获得锁,如果是获取写锁就会等待
当一个goroutine获取写锁之后，其他的goroutine无论是获取读锁还是写锁都会等待

读的goroutine来了获取的是读锁,后续的goroutine能读不能写
写的goroutine来了获取的是写锁,后续的goroutine不管是读还是写都要等待获取锁

需要注意的是读写锁非常适合读多写少的场景，如果读和写的操作差别不大，读写锁的优势就发挥不出来
*/

var (
	x      = 0
	wg     sync.WaitGroup
	lock   sync.Mutex   // 互斥锁
	rwLock sync.RWMutex // 读写互斥锁
)

// 读操作
func read() {
	defer wg.Done()
	// lock.Lock() // 互斥锁
	rwLock.RLock() // 读写锁的读锁
	time.Sleep(time.Millisecond * 2)
	// lock.Unlock() // 互斥锁
	rwLock.RUnlock() // 读写锁的读锁
}

// 写操作
func write() {
	defer wg.Done()
	// lock.Lock() // 互斥锁
	rwLock.Lock() // 读写锁的写锁
	x = x + 1
	time.Sleep(time.Millisecond * 20)
	// lock.Unlock() // 互斥锁
	rwLock.Unlock() // 读写锁的写锁
}

func main() {
	// 模拟读多写少的操作，写操作耗时多，但是次数少；读操作耗时少，但是次数多
	start := time.Now() //开始时间

	// 读的量是写的量的100倍
	// 写操作
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	time.Sleep(time.Second * 1) // 由于写太慢了，等写入完成以后再去读

	// 读操作
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()

	fmt.Printf("总耗时:%s\n", time.Since(start))
}

/*
结果
使用互斥锁耗时  		 总耗时:3.416285879 s
使用读写互斥锁耗时 	     总耗时:1.008857085s
*/

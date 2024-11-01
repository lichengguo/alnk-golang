package main

import (
	"fmt"
	"strconv"
	"sync"
)

// Go语言中内置的map不是并发安全的，超过一定个数的并发写入肯定报错

var m = make(map[string]int)

var lock sync.Mutex
var wg sync.WaitGroup

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func(n int) {
			key := strconv.Itoa(n) //转换成字符串类型的数字
			lock.Lock()            //互斥锁
			set(key, n)            //调用set函数
			lock.Unlock()          //互斥锁
			fmt.Printf("k=%v, v=%v\n", key, get(key))
			wg.Done()
		}(i)
	}

	wg.Wait()
}

/*
执行结果会报错
*/

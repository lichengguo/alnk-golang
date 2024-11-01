package main

import (
	"fmt"
	"strconv"
	"sync"
)

/*
[sync.Map]
Go语言的sync包中提供了一个开箱即用的并发安全版map， sync.Map
开箱即用表示不用像内置的map一样使用make函数初始化就能直接使用。
同时sync.Map内置了诸如Store、Load、LoadOrStore、Delete、Range等操作方法
*/

var (
	m2  = sync.Map{}
	wg1 sync.WaitGroup
)

func main() {
	for i := 0; i < 200; i++ {
		wg1.Add(1)

		go func(n int) {
			key := strconv.Itoa(n)   //int --> string
			m2.Store(key, n)         //写入
			value, _ := m2.Load(key) //读取
			fmt.Printf("k=%#v, v=%#v\n", key, value)
			wg1.Done()
		}(i)
	}
	wg1.Wait()
}

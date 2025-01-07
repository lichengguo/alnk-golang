package main

import "fmt"

func main() {

	// 创建一个channel
	ch := make(chan int)

	// 开启一个goroutine
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		// 关闭channel
		close(ch)
	}()

	// 通过range遍历channel
	for v := range ch {
		fmt.Println("接收到的数据", v)
	}

	// 主函数继续执行
	fmt.Println("主函数继续执行...")
}

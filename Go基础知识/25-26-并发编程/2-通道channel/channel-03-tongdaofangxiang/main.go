package main

import "fmt"

// 通道方向
// 当使用通道作为函数的参数时，可以指定这个通道是否为只读或只写
// 该特性可以提升程序的类型安全

func main() {
	// 创建一个双向通道
	ch := make(chan int)

	// 将通道通过不同函数的参数限制为发送或接收操作
	go sendData(ch, 10)
	go receiveData(ch)

	// 等待goroutine执行完毕后打印输入的内容
	var input string
	fmt.Scanln(&input)
	fmt.Printf("输出字符串: %s\n", input)
}

// sendData 通道只能用于发送数据的函数(只写)
func sendData(ch chan<- int, data int) {
	fmt.Println("发送", data)
	ch <- data
}

// receiveData 通道只能用于接收数据的函数(只读)
func receiveData(ch <-chan int) {
	data := <-ch
	fmt.Println("接收", data)
}

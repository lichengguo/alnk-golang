package main

// 定时器是当你想要在未来某一刻执行一次时使用的。而打点器 则是为你想要以固定的时间间隔重复执行而准备的。 这里是一个打点器的例子，它将定时的执行，直到我们将它停止。

import (
    "fmt"
    "time"
)

func main() {
    // 创建一个定时器，设置定时器的间隔为1秒
    ticker := time.NewTicker(1 * time.Second)

    // 启动一个匿名函数作为协程，用于处理定时器触发的事件
    go func() {
        // 循环遍历定时器的通道
        for {
            select {
            case t := <-ticker.C: // 定时器触发时
                fmt.Printf("%v | 定时器触发了\n", t)
            }
        }
    }()

    // 主函数继续执行其他操作
    fmt.Println("主函数继续执行...")

    // 等待10秒后停止定时器
    time.Sleep(10 * time.Second)
    ticker.Stop()

    // 打印出定时器停止的信息
    fmt.Println("定时器已停止")

    // 主函数继续执行其他操作
    fmt.Println("主函数继续执行...")
}

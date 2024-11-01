package main

// 我们经常需要在未来的某个时间点运行 Go 代码，或者每隔一定时间重复运行代码。 
// Go 内置的 定时器 和 打点器 特性让这些变得很简单。

import (
    "fmt"
    "time"
)

func main() {
    // 创建一个定时器，设置定时器触发时间为2秒后
    timer1 := time.NewTimer(2 * time.Second)

    // 手动阻塞主线程，等待定时器触发
    fmt.Println("等待定时器1触发...")

    // 通过 <-timer.C 从定时器的通道中接收数据，这将导致主线程阻塞，
    // 直到定时器触发或者定时器被停止
    <-timer1.C

    // 定时器触发后打印出相应的信息
    fmt.Println("定时器1触发了")

    // 手动停止定时器
    timer1.Stop()

    // 打印出定时器停止的信息
    fmt.Println("定时器1已停止")

    timer2 := time.NewTimer(time.Second)
    go func() {
        <-timer2.C
        fmt.Println("定时器2触发了")
    }()

    // 如果你需要的仅仅是单纯的等待，使用 time.Sleep 就够了。 
    // 使用定时器的原因之一就是，你可以在定时器触发之前将其取消。
    stop := timer2.Stop()
    if stop {
        fmt.Println("定时器2已停止")
    }

    time.Sleep(2 * time.Second)
    // 主函数继续执行其他操作
    fmt.Println("主函数继续执行...")
}

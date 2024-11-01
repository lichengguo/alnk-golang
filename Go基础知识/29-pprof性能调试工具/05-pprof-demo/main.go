package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

/*
[go性能优化]

在计算机性能调试领域里，profiling 是指对应用程序的画像，画像就是应用程序使用 CPU 和内存的情况。
Go语言是一个对性能特别看重的语言，因此语言中自带了 profiling 的库

Go语言项目中的性能优化主要有以下几个方面:
1.CPU profile：报告程序的 CPU 使用情况，按照一定频率去采集应用程序在 CPU 和寄存器上面的数据
2.Memory Profile（Heap Profile）：报告程序的内存使用情况
3.Block Profiling：报告 goroutines 不在运行状态的情况，可以用来分析和查找死锁等性能瓶颈
4.Goroutine Profiling：报告 goroutines 的使用情况，有哪些 goroutine，它们的调用关系是怎样的

采集性能数据
Go语言内置了获取程序的运行数据的工具，包括以下两个标准库:
1.runtime/pprof：采集工具型应用运行数据进行分析
2.net/http/pprof：采集服务型应用运行时数据进行分析

pprof开启后，每隔一段时间（10ms）就会收集下当前的堆栈信息，获取各个函数占用的CPU以及内存资源；
最后通过对这些采样数据进行分析，形成一个性能分析报告。
注意，我们只应该在性能测试的时候才在代码中引入pprof，因为性能测试会占用系统资源


[工具型应用]
如果你的应用程序是运行一段时间就结束退出类型。那么最好的办法是在应用退出的时候把 profiling 的报告保存到文件中，进行分析。
对于这种情况，可以使用runtime/pprof库。
首先在代码中导入runtime/pprof工具 import "runtime/pprof"

[CPU性能分析]
开启CPU性能分析 pprof.StartCPUProfile(w io.Writer)
停止CPU性能分析 pprof.StopCPUProfile()
应用执行结束后，就会生成一个文件，保存了我们的 CPU profiling 数据。得到采样数据之后，使用go tool pprof工具进行CPU性能分析

[内存性能分析]
记录程序的堆栈信息 pprof.WriteHeapProfile(w io.Writer)
得到采样数据之后，使用 go tool pprof 工具进行内存性能分析。
go tool pprof 默认是使用 -inuse_space 进行统计，还可以使用 -inuse-objects 查看分配对象的数量


[服务型应用]
如果你的应用程序是一直运行的，比如 web 应用，那么可以使用net/http/pprof库，它能够在提供 HTTP 服务进行分析
如果使用了默认的http.DefaultServeMux（通常是代码直接使用 http.ListenAndServe(“0.0.0.0:8000”, nil)），
只需要在你的web server端代码中按如下方式导入net/http/pprof
import _ "net/http/pprof"

*/

//一段有问题的代码
func logicCode() {
	var c chan int
	for {
		select {
		case v := <-c: //没初始化的chan，会阻塞，所以会走default分支
			fmt.Printf("recv from chan, value:%v\n", v)
		default:
			//time.Sleep(time.Millisecond * 500)
		}
	}
}

func main() {
	var isCPUprof bool
	var isMemProf bool

	flag.BoolVar(&isCPUprof, "cpu", false, "turn cpu pprof on")
	flag.BoolVar(&isMemProf, "mem", false, "turn mem pprof on")
	flag.Parse()

	//是否开启cpu检测
	if isCPUprof {
		f1, err := os.Create("./cpu.pprof") //在当前目录下创建一个cpu.pprof的文件
		if err != nil {
			fmt.Printf("create cpu pprof failed, err:%v\n", err)
			return
		}
		pprof.StartCPUProfile(f1) //往文件中记录cpu的信息
		defer func() {
			pprof.StopCPUProfile() //停止cpu性能分析
			f1.Close()             //关闭文件
		}()
	}

	//调用logicCode函数，模拟业务代码
	for i := 0; i < runtime.NumCPU(); i++ { //跑满整个cpu
		go logicCode()
	}
	time.Sleep(time.Second * 20) //模拟业务耗时时间

	if isMemProf {
		f2, err := os.Create("./mem.pprof")
		defer f2.Close()
		if err != nil {
			fmt.Printf("create mem pprof failed, err:%v\n", err)
			return
		}
		pprof.WriteHeapProfile(f2)
	}
}

/*
开始cpu性能分析
lichengguo@lichengguodeMacBook-Pro 05pprof_demo % ./05pprof_demo -cpu=true

使用go工具链里的pprof来分析一下
lichengguo@lichengguodeMacBook-Pro 05pprof_demo % go tool pprof cpu.pprof

执行上面的代码会进入交互界面如下：
Type: cpu
Time: Sep 2, 2020 at 11:31am (CST)
Duration: 20.12s, Total samples = 54.43s (270.46%)
Entering interactive mode (type "help" for commands, "o" for options)

在交互界面输入top3来查看程序中占用CPU前3位的函数
(pprof) top3
Showing nodes accounting for 49.97s, 91.81% of 54.43s total
Dropped 18 nodes (cum <= 0.27s)
Showing top 3 nodes out of 4
      flat  flat%   sum%        cum   cum%
       20s 36.74% 36.74%     41.81s 76.81%  runtime.selectnbrecv
    17.49s 32.13% 68.88%     18.84s 34.61%  runtime.chanrecv
    12.48s 22.93% 91.81%     54.33s 99.82%  main.logicCode
(pprof)

其中：
flat：当前函数占用CPU的耗时
flat：:当前函数占用CPU的耗时百分比
sun%：函数占用CPU的耗时累计百分比
cum：当前函数加上调用当前函数的函数占用CPU的总耗时
cum%：当前函数加上调用当前函数的函数占用CPU的总耗时百分比
最后一列：函数名称

还可以使用 list 命令查看具体的函数分析，例如执行 list logicCode 查看我们编写的函数的详细分析
(pprof) list logicCode
Total: 54.43s
ROUTINE ======================== main.logicCode in /Users/lichengguo/go/src/code.oldboyedu.com/gostudy/day09/05pprof_demo/main.go
    12.48s     54.33s (flat, cum) 99.82% of Total
         .          .     60://一段有问题的代码
         .          .     61:func logicCode() {
         .          .     62:   var c chan int
         .          .     63:   for {
         .          .     64:           select {
    12.48s     54.33s     65:           case v := <-c: //没初始化的chan，会阻塞，所以会走default分支
         .          .     66:                   fmt.Printf("recv from chan, value:%v\n", v)
         .          .     67:           default:
         .          .     68:           }
         .          .     69:   }
         .          .     70:}
(pprof)
通过分析发现大部分CPU资源被66行占用，我们分析出select语句中的default没有内容会导致上面的case v:=<-c:一直执行。
我们在default分支添加一行time.Sleep(time.Second)即可

[图形化]
MAC安装软件:brew install graphviz
(pprof) web
*/


## 单元测试

### 字符串切割函数

```go
package split_string

import (
	"fmt"
	"strings"
)

// Split:切割字符串
// example:
// abc, b --> [a c]
func Split(str string, sep string) []string {
	// 优化代码，初始化的时候指定长度和容量，避免在append的时候去动态扩容，影响性能
	var ret = make([]string, 0, strings.Count(str, sep)+1) //切片的make参数:类型、长度、容量

	index := strings.Index(str, sep)
	for index >= 0 {
		ret = append(ret, str[:index])
		str = str[index+len(sep):] //注意这不是切片,这是字符串切割
		index = strings.Index(str, sep)
	}
	ret = append(ret, str)

	// 为了让测试率达不到100%，只是试验而已，以后可以不用写这个if
	if index == -5 {
		fmt.Println("No!!!")
	}

	return ret
}

// 用来做性能比较测试的例子
// Fib是一个计算第n个斐波那契数的函数
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
```



### 单元测试

```go
package split_string

import (
	"reflect"
	"testing"
)

/*
【单元测试】
每个测试函数必须导入testing包
测试函数的名字必须以Test开头，可选的后缀名必须以大写字母开头,举几个例子：
func TestAdd(t *testing.T){ ... }
func TestSum(t *testing.T){ ... }
func TestLog(t *testing.T){ ... }
其中参数t用于报告测试失败和附加的日志信息。 testing.T

在命令行执行命令进行单元测试
在 *_test.go 文件所在的目录，执行命令:
	go test
结果:
	PASS
	ok      code.oldboyedu.com/gostudy/day09/04test_split/split_string      0.005s

缺点:
    1.代码没复用
*/

//func TestSplit(t *testing.T) { // 测试函数名必须以Test开头，必须接收一个*testing.T类型参数
//	got := Split("babcbef", "b")         // 程序输出的结果
//	want := []string{"", "a", "c", "ef"} // 期望的结果
//
//	// 因为slice不能比较直接，借助反射包中的方法比较
//	if !reflect.DeepEqual(got, want) {
//		t.Errorf("want:%#v but got:%#v\n", want, got) // 测试失败输出错误提示
//	}
//}

//func TestS2plit(t *testing.T) {
//	got := Split("a:b:c", ":")
//	want := []string{"a", "b", "c"}
//
//	if !reflect.DeepEqual(got, want) {
//		t.Errorf("want:%#v but got:%#v\n", want, got)
//	}
//}

func TestS3plit(t *testing.T) {
	got := Split("abcef", "bc")
	want := []string{"a", "efc"} //这里故意出错

	if !reflect.DeepEqual(got, want) {
		//测试用例失败
		t.Errorf("want:%#v but got:%#v\n", want, got)
	}
}

/*
--- FAIL: TestS3plit (0.00s)
    split_test.go:39: want:[]string{"a", "efc"} but got:[]string{"a", "ef"}
FAIL
exit status 1
FAIL    code.oldboyedu.com/gostudy/day09/04test_split/split_string      0.005s
*/
```



### 测试组

```go
package split_string

import (
	"reflect"
	"testing"
)

/*
【测试组】
缺点：
	1.不够灵活
	2.测试如果不通过，报错不明显
	3.不能跑单个测试用例
*/

func TestSplit(t *testing.T) {
	// 定义一个测试用例类型
	type test struct {
		input string   //输出
		sep   string   //分隔符
		want  []string //期望结果
	}

	// 定义一个存储测试用例的切片
	tests := []test{
		{input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		{input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		{input: "abcd", sep: "bc", want: []string{"a", "d"}},
		{input: "沙河有沙又有河", sep: "沙", want: []string{"", "河有", "又有河"}},
	}

	// 遍历切片，逐一执行测试用例
	for _, tc := range tests {
		got := Split(tc.input, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("got:%#v but want:%#v\n", got, tc.want)
		}

	}
}

/*
% go test
PASS
ok      code.oldboyedu.com/gostudy/day09/04test_split/split_string      0.006s
*/
```



### 子测试

```go
package split_string

import (
	"reflect"
	"testing"
)

/*
【子测试】

优点:
	1.报错明显
	2.可以跑单个测试用例
	3.输出结果更加灵活和详细
*/

/*
其他一些命令:
1.测试覆盖率
go test -cover

2.测试结果同时输出到文件
go test -cover -coverprofile=c.out

3.使用浏览器来打开, 绿色标记的语句块表示被覆盖了，而红色的表示没有被覆盖
go tool cover -html=c.out
*/

func TestSplit(t *testing.T) {
	// 定义一个测试用例类型
	type test struct {
		input string   //输入
		sep   string   //分隔符
		want  []string //期望结果
	}

	// 测试用例使用map存储
	tests := map[string]test{
		"simple":    {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong_sep": {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"more_sep":  {input: "沙河有沙又有河", sep: "沙", want: []string{"", "河有", "又有河"}},
	}

	// 遍历切片，逐一执行测试用例
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("name:%s want:%#v, got:%#v\n", name, tc.want, got) //将测试用例的name格式化输出
			}
		})
	}
}

/*
lichengguo@lichengguodeMacBook-Pro split_string % go test -v
=== RUN   TestSplit
=== RUN   TestSplit/simple
=== RUN   TestSplit/wrong_sep
=== RUN   TestSplit/more_sep
--- PASS: TestSplit (0.00s)
    --- PASS: TestSplit/simple (0.00s)
    --- PASS: TestSplit/wrong_sep (0.00s)
    --- PASS: TestSplit/more_sep (0.00s)
PASS
ok      code.oldboyedu.com/gostudy/day09/04test_split/split_string      0.005s

指定某个子测试用例，例如
lichengguo@lichengguodeMacBook-Pro split_string % go test -v -run=Split/simple
=== RUN   TestSplit
=== RUN   TestSplit/simple
--- PASS: TestSplit (0.00s)
    --- PASS: TestSplit/simple (0.00s)
PASS
ok      code.oldboyedu.com/gostudy/day09/04test_split/split_string      0.005s
*/
```



### 基准测试

```go
package split_string

import "testing"

/*
【基准测试】
基准测试就是在一定的工作负载之下检测程序性能的一种方法
基准测试以Benchmark为前缀，需要一个*testing.B类型的参数b
基准测试必须要执行b.N次，这样的测试才有对照性，b.N的值是系统根据实际情况去调整的，从而保证测试的稳定性

默认情况下，每个基准测试至少运行1秒。如果在Benchmark函数返回时没有到1秒，则b.N的值会按1,2,5,10,20,50，…增加，并且函数再次运行
*/

// BenchmarkSplit 基准测试
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c", ":")
	}
}

/*
在命令行执行: go test -bench=Split
结果:
goos: darwin
goarch: amd64
pkg: code.oldboyedu.com/gostudy/day09/04test_split/split_string
BenchmarkSplit-4         9997287               121 ns/op
PASS
ok      code.oldboyedu.com/gostudy/day09/04test_split/split_string      1.338s
结果注释:
BenchmarkSplit-4: 表示对Split函数进行基准测试,数字4表示GOMAXPROCS的值,这个对于并发基准测试很重要
9997287  121 ns/op 表示执行9997287次，平均每次耗时121 ns


添加 -benchmem 参数，来获得内存分配的统计数据
lichengguo@lichengguodeMacBook-Pro split_string % go test -bench=Split -benchmem
goos: darwin
goarch: amd64
pkg: code.oldboyedu.com/gostudy/day09/04test_split/split_string
BenchmarkSplit-4         4889278               240 ns/op             112 B/op          3 allocs/op
PASS
ok      code.oldboyedu.com/gostudy/day09/04test_split/split_string      1.431s
结果注释:
112 B/op 表示每次操作内存分配了112字节
3 allocs/op 则表示每次操作进行了3次内存分配


优化被测试函数以后
goos: darwinichengguodeMacBook-Pro split_string % go test -bench=Split -benchmem
goarch: amd64
pkg: code.oldboyedu.com/gostudy/day09/04test_split/split_string
BenchmarkSplit-4         9878463               122 ns/op              48 B/op          1 allocs/op
PASS
ok      code.oldboyedu.com/gostudy/day09/04test_split/split_string      1.339s
从结果可以看出:
只做了一次内存申请，并且占用的内存空间更小，执行时间更短，速度更快
*/
```



### 性能比较函数

```go
package split_string

import (
	"testing"
)

/*
[性能比较函数]
基准测试只能得到给定操作的绝对耗时
但是在很多性能问题是发生在两个不同操作之间的相对耗时，比如同一个函数处理1000个元素的耗时与处理1万甚至100万个元素的耗时的差别是多少?
再或者对于同一个任务究竟使用哪种算法性能最佳？我们通常需要对两个不同算法的实现使用相同的输入来进行基准比较测试
性能比较函数通常是一个带有参数的函数，被多个不同的Benchmark函数传入不同的值来调用


*/

// benchmarkFib 中间函数
func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}

}

// BenchmarkFib 性能比较测试
func BenchmarkFib1(b *testing.B)  { benchmarkFib(b, 1) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(b, 2) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(b, 3) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(b, 10) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(b, 20) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(b, 40) }

/*
split $ go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/Q1mi/studygo/code_demo/test_demo/fib
BenchmarkFib1-4         1000000000               2.03 ns/op
BenchmarkFib2-4         300000000                5.39 ns/op
BenchmarkFib3-4         200000000                9.71 ns/op
BenchmarkFib10-4         5000000               325 ns/op
BenchmarkFib20-4           30000             42460 ns/op
BenchmarkFib40-4               2         638524980 ns/op
PASS
ok      github.com/Q1mi/studygo/code_demo/test_demo/fib 12.944s
这里需要注意的是，默认情况下，每个基准测试至少运行1秒。
如果在Benchmark函数返回时没有到1秒，则b.N的值会按1,2,5,10,20,50，…增加，并且函数再次运行
最终的BenchmarkFib40只运行了两次，每次运行的平均值只有不到一秒。
像这种情况下我们应该可以使用-benchtime标志增加最小基准时间，以产生更准确的结果
例如：
split $ go test -bench=Fib40 -benchtime=20s
goos: darwin
goarch: amd64
pkg: github.com/Q1mi/studygo/code_demo/test_demo/fib
BenchmarkFib40-4              50         663205114 ns/op
PASS
ok      github.com/Q1mi/studygo/code_demo/test_demo/fib 33.849s
这一次BenchmarkFib40函数运行了50次，结果就会更准确一些了
*/
```



### 重置时间

```go
package split_string

import (
	"testing"
	"time"
)

/*
[重置时间]
b.ResetTimer之前的处理不会放到执行时间里，也不会输出到报告中，所以可以在之前做一些不计划作为测试报告的操作
*/

func BenchmarkSplit(b *testing.B) {
	time.Sleep(5 * time.Second) // 假设需要做一些耗时的无关操作
	b.ResetTimer()              // 重置计时器
	for i := 0; i < b.N; i++ {
		Split("沙河有沙又有河", "沙")
	}
}

/*
注释掉 b.ResetTimer() 的结果 刚看的时候可能以为是错的，总耗时居然比没注释还要少
但是看他的执行次数是1次，平均每次耗时 5005160299 ns
lichengguo@lichengguodeMacBook-Pro split_string % go test -bench=.
goos: darwin
goarch: amd64
pkg: code.oldboyedu.com/gostudy/day09/04test_split/split_string
BenchmarkSplit-4               1        5005160299 ns/op
PASS
ok      code.oldboyedu.com/gostudy/day09/04test_split/split_string      5.012s

没注释 b.ResetTimer() 的结果
lichengguo@lichengguodeMacBook-Pro split_string % go test -bench=.
goos: darwin
goarch: amd64
pkg: code.oldboyedu.com/gostudy/day09/04test_split/split_string
BenchmarkSplit-4         5357868               191 ns/op
PASS
ok      code.oldboyedu.com/gostudy/day09/04test_split/split_string      26.285s
*/
```



### 并行测试

```go
package split_string

import "testing"

/*
[并行测试]
b.RunParallel 会创建出多个goroutine，并将b.N分配给这些goroutine执行
其中goroutine数量的默认值为GOMAXPROCS
用户如果想要增加非CPU受限（non-CPU-bound）基准测试的并行性
那么可以在RunParallel之前调用SetParallelism

RunParallel通常会与-cpu标志一同使用
还可以通过在测试命令后添加-cpu参数如 go test -bench=. -cpu 1 来指定使用的CPU数量
*/

func BenchmarkSplitParallel(b *testing.B) {
	//b.SetParallelism(1) // 设置使用cpu的数
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Split("沙河有沙又有河", "沙")
		}
	})
}

/*
单核心
lichengguo@lichengguodeMacBook-Pro split_string % go test -bench=. -cpu 1
goos: darwin
goarch: amd64
pkg: code.oldboyedu.com/gostudy/day09/04test_split/split_string
BenchmarkSplitParallel   7113730               170 ns/op
PASS
ok      code.oldboyedu.com/gostudy/day09/04test_split/split_string      1.390s

4核心
lichengguo@lichengguodeMacBook-Pro split_string % go test -bench=. -cpu 4
goos: darwin
goarch: amd64
pkg: code.oldboyedu.com/gostudy/day09/04test_split/split_string
BenchmarkSplitParallel-4        13540974                87.3 ns/op
PASS
ok      code.oldboyedu.com/gostudy/day09/04test_split/split_string      1.279s
可以看出4核心平均每次执行的效果还是比单核心快了一倍

2核心
lichengguo@lichengguodeMacBook-Pro split_string % go test -bench=. -cpu 2
goos: darwin
goarch: amd64
pkg: code.oldboyedu.com/gostudy/day09/04test_split/split_string
BenchmarkSplitParallel-2        12745449                89.6 ns/op
PASS
ok      code.oldboyedu.com/gostudy/day09/04test_split/split_string      1.245s

*/
```



## pprof

```go
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
```



## 练习

目录结构

```shell
├── main.go
└── palindrome
    ├── palindrome.go
    └── palindrome_test.go
```



main.go文件

```go
package main

import (
	"code.oldboyedu.com/gostudy/day09/99homework/palindrome"
	"fmt"
)

/*
编写一个回文检测函数，并为其编写单元测试和基准测试，根据测试的结果逐步对其进行优化。
回文：一个字符串正序和逆序一样，如“Madam,I’mAdam”、“油灯少灯油”等
*/

func main() {
	if palindrome.Palindrome("abcba") {
		fmt.Println("是回文")
	}

}
```



palindrome.go文件

```go
package palindrome

/*
编写一个回文检测函数，并为其编写单元测试和基准测试，根据测试的结果逐步对其进行优化。
回文：一个字符串正序和逆序一样，如“Madam,I’mAdam”、“油灯少灯油”等
*/

func Palindrome(str string) bool {
	//1.将字符串转为[]rune类型的切片
	runeSlice := []rune(str)
	//2.拿切片的第一个元素和最后一个元素对比
	//再拿第二个元素和倒数第二个元素对比
	//对比的次数为切片的总长度 /2 取商
	for i := 0; i < len(runeSlice)/2; i++ {
		if runeSlice[i] != runeSlice[len(runeSlice)-1-i] {
			return false
		}
	}
	return true
}
```



palindrome_test.go文件

```go
package palindrome

import (
	"reflect"
	"testing"
)

//TestPalindrome 单元测试
//func TestPalindrome(t *testing.T) {
//	got := Palindrome("油灯少灯油")
//	want := true
//	if got != want {
//		t.Errorf("want:%v but got:%v\n", want, got)
//	}
//}

//TestPalindrome 子测试
func TestPalindrome(t *testing.T) {
	//定义一个测试用的类型
	type test struct {
		input string //输入
		want  bool   //输出
	}

	//测试用例使用map存储
	tests := map[string]test{
		"t1": {input: "abcdedcba", want: true},
		"t2": {input: "油灯少灯油", want: true},
		"t3": {input: "Madam,I’mAdam", want: false},
	}

	//遍历map，逐一测试
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Palindrome(tc.input)
			want := tc.want
			if !reflect.DeepEqual(got, want) {
				t.Errorf("name:%s wang:%v got:%v\n", name, want, got)
			}
		})
	}
}

/*
lichengguo@lichengguodeMacBook-Pro palindrome % go test -v
=== RUN   TestPalindrome
=== RUN   TestPalindrome/t1
=== RUN   TestPalindrome/t2
=== RUN   TestPalindrome/t3
--- PASS: TestPalindrome (0.00s)
    --- PASS: TestPalindrome/t1 (0.00s)
    --- PASS: TestPalindrome/t2 (0.00s)
    --- PASS: TestPalindrome/t3 (0.00s)
PASS
ok      code.oldboyedu.com/gostudy/day09/99homework/palindrome  0.005s
*/

//BenchmarkPalindrome 基准测试
func BenchmarkPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Palindrome("再或者对于同一个任务究竟使用哪种算法性能最佳?")
	}
}

/*
lichengguo@lichengguodeMacBook-Pro palindrome % go test -bench=Palindrome  -benchmem
goos: darwin
goarch: amd64
pkg: code.oldboyedu.com/gostudy/day09/99homework/palindrome
BenchmarkPalindrome-4            3765415               308 ns/op               0 B/op          0 allocs/op
PASS
ok      code.oldboyedu.com/gostudy/day09/99homework/palindrome  1.490s
*/
```
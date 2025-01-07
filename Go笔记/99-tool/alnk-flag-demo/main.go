package main

import (
	"flag"
	"fmt"
	"time"
)

/*
flag包实现了命令行参数的解析
flag包支持的命令行参数类型有bool、int、int64、uint、uint64、float float64、string、duration(时间)

字符串flag				合法字符串
整数flag				1234、0664、0x1234等类型，也可以是负数。
浮点数flag				合法浮点数
bool类型flag			1, 0, t, f, T, F, true, false, TRUE, FALSE, True, False。
时间段flag				任何合法的时间段字符串。如”300ms”、”-1.5h”、”2h45m”。
						合法的单位有”ns”、”us” /“µs”、”ms”、”s”、”m”、”h”。
*/

func f1() {
	// 定义命令行参数方式
	var name string
	var age int
	var married bool
	var delay time.Duration

	flag.StringVar(&name, "name", "张三", "姓名") //变量指针，参数名称，默认值，help提示语
	flag.IntVar(&age, "age", 0, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "delay", 0, "结婚多少时间了")

	// 解析命令行参数
	flag.Parse()

	//
	fmt.Println(name, age, married, delay)

	// 返回命令行参数后的其他参数
	fmt.Println(flag.Args()) //[]string 类型

	// 返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())

	// 返回使用的命令行参数的个数
	fmt.Println(flag.NFlag())
}

func main() {
	f1()
}

/*
使用方法
提示
% ./03flag_demo -help
Usage of ./02flag_demo:
  -age int
        年龄 (default 18)
  -delay duration
        延迟时间间隔
  -married
        婚否
  -name string
        姓名 (default "张三")

默认值
% ./03flag_demo
张三 18 false 0s
[]
0
0

正常使用命令行参数
./02flag_demo -name 沙河娜扎 --age 28 -married=false -delay=1h30m
沙河娜扎 28 false 1h30m0s
[]
0
4

% ./02flag_demo -name 沙河娜扎 --age 28 -married=false -delay=1h30m a b c
沙河娜扎 28 false 1h30m0s
[a b c]  flag.Args()
3 		 flag.NArg()
4		 flag.NFlag()
*/

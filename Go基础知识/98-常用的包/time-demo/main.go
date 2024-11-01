package main

import (
	"fmt"
	"time"
)

// time时间包

func f1() {
	now := time.Now()         //获取当前时间
	fmt.Println(now)          //2020-08-19 14:37:25.136076 +0800 CST m=+0.000102652
	fmt.Println(now.Year())   //2020
	fmt.Println(now.Month())  //August
	fmt.Println(now.Day())    //19
	fmt.Println(now.Hour())   //14
	fmt.Println(now.Minute()) //39
	fmt.Println(now.Second()) //30
	fmt.Println(now.Date())   //日期 2020 August 19

	// 时间戳
	fmt.Println(now.Unix())     //1597819214 获取当前秒
	fmt.Println(now.UnixNano()) //1597819255919494000  纳秒时间戳

	// time.Unix() 将时间戳转为时间格式
	ret := time.Unix(1597819214, 0)
	fmt.Println(ret)                         //2020-08-19 14:40:14 +0800 CST
	ret1 := time.Unix(1597819214, 484199000) //精确到纳秒
	fmt.Println(ret1)                        //2020-08-19 14:40:14.484199 +0800 CST
	fmt.Println(ret1.Year())                 //2020
	fmt.Println(ret1.Month())                //August
	fmt.Println(ret1.Day())                  //19
	fmt.Println(ret1.Date())                 //日期 2020 August 19

	// 时间间隔
	fmt.Println(time.Second)      //1s
	fmt.Println(time.Second * 10) //10s

	// 定时器
	//timer := time.Tick(5 * time.Second) //间隔5s
	////timer := time.Tick(1 * time.Second) //间隔1s
	//for t := range timer {
	//	fmt.Println("定时器")
	//	fmt.Println(t) //1秒钟执行一次
	//}

	// 格式化时间:把语言中时间对象转换成字符串类型的时间
	// 时间类型有一个自带的方法Format进行格式化，需要注意的是Go语言中格式化时间模板不是常见的Y-m-d H:M:S
	// 而是使用Go的诞生时间2006年1月2号15点04分05秒（记忆口诀为2006 1 2 3 4 5）
	fmt.Println(now.Format("2006-01-02"))                 //2020-08-19
	fmt.Println(now.Format("2006/01/02 15:04:05"))        //2020/08/19 07:41:40
	fmt.Println(now.Format("2006/01/02 03:04:05 PM"))     //2020/08/19 07:43:07 AM
	fmt.Println(now.Format("2006/01/02 15:04:05.000"))    //2020/08/19 07:45:52.486
	fmt.Println(time.Now().Format("2006-01-02 15:04:05")) //2020-08-19 15:51:08
	fmt.Println(time.Now().Format("2006-01-02"))          //2020-08-19

	// 按照对应的格式解析字符串类型的时间
	timeObj, _ := time.Parse("2006-01-02", "2019-08-03")
	fmt.Println(timeObj)        //2019-08-03 00:00:00 +0000 UTC
	fmt.Println(timeObj.Unix()) //1564790400

	// sleep
	n := 5 //int
	fmt.Println("开始sleep了")
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println("5秒钟过去了")
	time.Sleep(5 * time.Second)
	fmt.Println("5秒又钟过去了...")
}

// 根据时区计算时间差
func f2() {
	now := time.Now() //本地当前时间
	fmt.Println(now)  //2020-08-21 11:43:49.71391 +0800 CST m=+0.000093483

	// 按照东八区的时区和格式取解析一个字符串格式的时间
	// 根据字符串加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return
	}

	// 按照指定时区解析时间
	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", "2021-03-13 14:18:00", loc)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println(timeObj)

	//时间对象相减
	td := now.Sub(timeObj)
	fmt.Println(td) //24h0m38.279607s

}

// Add Sub 时间的加减
func f3() {
	// Add 时间相加
	now := time.Now() //获取当前时间
	fmt.Println(now)  //2020-08-21 10:31:43.955684 +0800 CST m=+0.000092436

	//十分钟以前
	m, _ := time.ParseDuration("-10m")
	m1 := now.Add(m)
	fmt.Println(m1)

	//8个小时以前
	h, _ := time.ParseDuration("-1h")
	h1 := now.Add(8 * h)
	fmt.Println(h1)

	//一天以前
	d, _ := time.ParseDuration("-24h")
	d1 := now.Add(d)
	fmt.Println(d1)

	//十分钟以后
	mm, _ := time.ParseDuration("10m")
	mm1 := now.Add(mm)
	fmt.Println(mm1)

	//8个小时以后
	hh, _ := time.ParseDuration("1h")
	hh1 := now.Add(8 * hh)
	fmt.Println(hh1)

	//一天以后
	dd, _ := time.ParseDuration("24h")
	dd1 := now.Add(dd)
	fmt.Println(dd1)

	// Sub 计算两个时间差
	subM := now.Sub(m1)
	fmt.Println(subM.Minutes()) //10分钟

	subH := now.Sub(h1)
	fmt.Println(subH) //8h0m0s

	subD := now.Sub(d1)
	fmt.Println(subD) //24h0m0s
}

// 计算两个string类型的时间差
func f4() {
	// 计算两个固定的string类型的时间差
	// 1.声明变量
	t1 := "2020-08-21 11:25:00"
	t2 := "2020-08-21 10:25:00"

	// 2.把string类型转化为time类型
	var baseTime = "2006-01-02 15:04:05"
	t1Time, _ := time.Parse(baseTime, t1)
	t2Time, _ := time.Parse(baseTime, t2)
	fmt.Println(t1Time) //2020-08-21 11:25:00 +0000 UTC
	fmt.Println(t2Time) //2020-08-21 10:25:00 +0000 UTC

	// 3.利用sub计算时间差
	sub1 := t1Time.Sub(t2Time)
	sub2 := t2Time.Sub(t1Time)
	fmt.Println(sub1)           //1h0m0s
	fmt.Println(sub2)           //-1h0m0s
	fmt.Println(sub1.Hours())   //1
	fmt.Println(sub1.Minutes()) //60
	fmt.Println(sub1.Seconds()) //3600

	// 计算某个时间(string类型)和当前时间的时间差
	// 注意要计算一个固定的时间串和本地时间间隔多少，解析时间字符串的时候需要把时区设置进去

	// 1.声明变量
	t3 := "2021-03-12 15:31:00"
	// 1.1设置时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return
	}
	// 1.2按照时区去解析时间
	t3Time, err := time.ParseInLocation("2006-01-02 15:04:05", t3, loc)
	if err != nil {
		return
	}
	fmt.Println(t3Time)

	fmt.Println(t3Time) //2020-09-08 18:03:00 +0000 UTC  //这里时区不对

	// 2.获取当前时间
	now := time.Now()
	fmt.Println(now) //2020-09-09 18:06:10.040535 +0800 CST m=+0.000185949

	// 3.相减
	sub3 := now.Sub(t3Time)

	// 4
	fmt.Println(sub3)           //24h0m24.746951s
	fmt.Println(sub3.Hours())   //24.006874153055556
	fmt.Println(sub3.Minutes()) //1440.4124491833334
	fmt.Println(sub3.Seconds()) //86424.746951
}

func main() {
	//f1()
	//f2()
	//f3()
	f4()
}

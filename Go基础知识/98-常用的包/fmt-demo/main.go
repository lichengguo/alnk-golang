package main

import "fmt"

/*
fmt包主要分为 向外输出内容 和 获取输入内容 两大部分
向外输出内容：
	终端输出：Print系列 Print、Printf、Println
	文件输出：Fprint系列 Fprint、Fprintf、Fprintln函数会将内容输出到一个io.Writer接口类型的变量中，我们通常用这个函数往文件中写入内容

	Sprint系列函数会把传入的数据生成并返回一个字符串，拼接字符串
	Errorf函数根据format参数生成格式化字符串并返回一个包含该字符串的错误

获取输入内容
	Go语言fmt包下有fmt.Scan、fmt.Scanf、fmt.Scanln三个函数，可以在程序运行过程中从标准输入获取用户的输入
	fmt.Scan
		Scan从标准输入扫描文本，读取由空白符分隔的值保存到传递给本函数的参数中，换行符视为空白符。
	fmt.Scanf

	fmt.Scanln
		fmt.Scanln遇到回车就结束扫描了，这个比较常用

有时候我们想完整获取输入的内容，而输入的内容可能包含空格，这种情况下可以使用bufio包来实现
func bufioDemo() {
	reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
	fmt.Print("请输入内容：")
	text, _ := reader.ReadString('\n') // 读到换行
	text = strings.TrimSpace(text)
	fmt.Printf("%#v\n", text)
}

*/

func main() {
	//fmt.Print("沙河")
	//fmt.Print("alex")
	//fmt.Println("------------")
	//fmt.Println("沙河")
	//fmt.Println("alex")
	//Printf("格式化字符串", 值)
	// %T :查看类型
	// %d :十进制数
	// %b ：二进制数
	// %o :八进制数
	// %x ：十六进制数
	// %c : 字符
	// %s ：字符串
	// %p： 指针
	// %v： 值
	// %f：浮点数
	// %t ：布尔值

	//var m1 = make(map[string]int, 1)
	//m1["lixiang"] = 100
	//fmt.Printf("%v\n", m1)  //map[lixiang:100]
	//fmt.Printf("%#v\n", m1) //map[string]int{"lixiang":100}

	//printBaifenbi(10)

	//fmt.Printf("%v\n", 100)
	////整数 -> 字符
	//fmt.Printf("%q\n", 65) // 'A'
	////浮点数 -> 复数
	//fmt.Printf("%b\n", 3.1415926)
	////字符串
	//fmt.Printf("%q\n", "理想有理想")
	//fmt.Printf("%7.3s\n", "abcdefghijk")

	//获取用户输入
	//var s string
	//fmt.Print("请输入内容：")
	//fmt.Scan(&s)
	//fmt.Println("用户输入的内容是：", s)

	//var (
	//	name  string
	//	age   int
	//	class string
	//)
	////fmt.Scanf("%s %d %s\n", &name, &age, &class) //获取用户输入，同一行
	////fmt.Println(name, age, class)
	////
	//fmt.Scanln(&name, &age, &class)
	//fmt.Println(name, age, class)

	//fmt.Printf("%b\n", 1024) //10000000000

}

func printBaifenbi(num int) {
	fmt.Printf("%d%%\n", num)
}

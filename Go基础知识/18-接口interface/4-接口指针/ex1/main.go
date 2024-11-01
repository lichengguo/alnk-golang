package main

import "fmt"

// 注意：这是一道你需要回答“能”或者“不能”的题！
// 首先请观察下面的这段代码，然后请回答这段代码能不能通过编译？

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

func main() {
	var peo People
	//peo = Student{}
	peo = &Student{} //注意这里要返回指针类型才行，因为方法中接收的是指针类型

	think := "bitch"
	fmt.Println(peo.Speak(think))
}

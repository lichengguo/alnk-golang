package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

func main() {
	// 执行shell命令
	// cmd := exec.Command("whoami")
	// str, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(string(str))

	// 在本地执行shell命令
	// cmd := exec.Command("tar", "-czf", "test.tar.gz", "a.txt", "b")
	// str, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(string(str))

	// 命令超时
	timeOut := 5
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeOut)*time.Second)
	defer cancel()

	cmdarray := []string{"-c", fmt.Sprintf("%s %s", "sleep", "10")}
	cmd := exec.CommandContext(ctx, "bash", cmdarray...)
	out, err := cmd.CombinedOutput()

	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("命令超时")
	}

	if err != nil {
		fmt.Println("命令错误")
	}

	fmt.Printf("ctx.Err : [%v]\n", ctx.Err())   // ctx.Err : [context deadline exceeded]
	fmt.Printf("error   : [%v]\n", err)         // error   : [signal: killed]
	fmt.Printf("out     : [%s]\n", string(out)) // out     : []

}

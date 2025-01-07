package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// 要执行的命令
	cmd := "ls -l"

	// 创建一个新的执行器
	execCmd := exec.Command("bash", "-c", cmd)

	// 执行命令并等待返回结果
	output, err := execCmd.Output()
	if err != nil {
		panic(err)
	}

	// 打印输出结果
	fmt.Println(string(output))
}

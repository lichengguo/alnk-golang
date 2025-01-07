package main

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/ssh"
)

func pwdConnect(sshHost, sshUser, sshPassword string, sshPort int) (*ssh.Client, error) {
	// 创建ssh登录配置
	config := &ssh.ClientConfig{
		Timeout:         5 * time.Second,                             // 超时时间
		User:            sshUser,                                     // 登录账号
		Auth:            []ssh.AuthMethod{ssh.Password(sshPassword)}, // 密码
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),                 // 忽略主机密钥验证，这个不够安全，生产环境不建议使用
		// HostKeyCallback: ssh.FixedHostKey(), // 这个更加安全，生产环境建议使用这种
	}

	// dial连接服务器
	addr := fmt.Sprintf("%s:%d", sshHost, sshPort)
	Client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		log.Fatal("连接到服务器失败", err)
		return nil, err
	}

	//defer sshClient.Close()
	return Client, nil
}

func main() {

	conn, err := pwdConnect("192.168.3.121", "root", "123456", 22)
	if err != nil {
		return
	}
	defer conn.Close()

	// 创建 ssh session 会话
	session, err := conn.NewSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// 执行远程命令
	cmd := "cd /tmp/;ls -l; tar -czf test.tar.gz hello.txt test.txt;ls -l"
	cmdInfo, err := session.CombinedOutput(cmd)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(cmdInfo))
}

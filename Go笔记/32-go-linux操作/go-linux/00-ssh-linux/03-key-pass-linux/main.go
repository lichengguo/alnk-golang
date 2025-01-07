package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
)

func keyOrPwdConnect(sshHost, sshUser, sshPassword, sshKey string, sshPort int) (*ssh.Client, error) {
	// 创建ssh登录配置
	config := &ssh.ClientConfig{
		Timeout:         5 * time.Second,             // 超时时间
		User:            sshUser,                     // 登录账号
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 忽略主机密钥检查，这个不够安全，生产环境不建议使用
		// HostKeyCallback: ssh.FixedHostKey(), // 建议使用这种
	}

	// 秘钥登录
	if sshKey != "" {
		// 读取秘钥
		key, err := os.ReadFile(sshKey)
		if err != nil {
			panic("秘钥读取失败")
		}
		// 创建秘钥签名
		signer, err := ssh.ParsePrivateKey(key)
		if err != nil {
			panic("秘钥签名失败")
		}
		// 配置秘钥登录
		config.Auth = []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		}
	} else if sshPassword != "" {
		// 密码登录
		config.Auth = []ssh.AuthMethod{ssh.Password(sshPassword)}
	} else {
		err := errors.New("秘钥或者密码登录")
		return nil, err
	}

	// dial连接服务器
	addr := fmt.Sprintf("%s:%d", sshHost, sshPort)
	Client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		log.Fatal("连接到服务器失败", err)
		return nil, err
	}

	return Client, nil
}

func main() {
	// 连接服务器
	// conn, err := keyOrPwdConnect("192.168.3.121", "root", "", "id_rsa", 22) // 秘钥登录
	conn, err := keyOrPwdConnect("192.168.3.121", "root", "123456", "", 22) // 密码登录
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 创建 ssh session 会话
	session, err := conn.NewSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// 执行远程命令
	cmd := "echo 'hello world!'"
	cmdInfo, err := session.CombinedOutput(cmd)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(cmdInfo))
}

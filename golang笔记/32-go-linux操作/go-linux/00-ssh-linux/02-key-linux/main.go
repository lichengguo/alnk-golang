package main

import (
	"fmt"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
)

// 使用秘钥连接Linux服务器

// 连接到linux服务器
func keyConnect(sshUser, sshHost, sshKey string, sshPort int) (*ssh.Client, error) {
	// 读取秘钥文件
	key, err := os.ReadFile(sshKey)
	if err != nil {
		panic("秘钥读取失败")
	}

	// 创建秘钥签名
	// 会拿着秘钥去登录服务器
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		panic("秘钥签名失败")
	}

	// ssh配置
	config := &ssh.ClientConfig{
		Timeout: 5 * time.Second, // 超时时间
		User:    sshUser,
		Auth: []ssh.AuthMethod{
			// 使用秘钥登录远程服务器
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 这个不够安全，生产环境不建议使用
		//var hostKey ssh.PublicKey
		// HostKeyCallback: ssh.FixedHostKey(hostKey),
	}

	// 连接远程服务器
	addr := fmt.Sprintf("%s:%d", sshHost, sshPort)
	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		fmt.Println("连接远程服务器失败", err)
		return nil, err
	}

	// defer client.Close()
	return client, nil
}

func main() {
	// 连接服务器
	conn, err := keyConnect("root", "192.168.3.121", "id_rsa", 22)
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
	cmd := "cd /tmp/; ls -l"
	cmdInfo, err := session.CombinedOutput(cmd)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(cmdInfo))
}

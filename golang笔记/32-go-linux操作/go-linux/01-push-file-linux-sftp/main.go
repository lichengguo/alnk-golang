package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// 通过SSH连接Linux服务器
func keyOrPwdConnectLinuxServer(sshHost, sshUser, sshPassword, sshKey string, sshPort int) (*ssh.Client, error) {
	// 创建ssh登录配置
	config := &ssh.ClientConfig{
		Timeout:         5 * time.Second,             // 超时时间
		User:            sshUser,                     // 登录账号
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 这个不够安全，生产环境不建议使用
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
		// 会拿着秘钥去登录服务器
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

// 创建sftp会话
func CreateSftp(sshHost, sshUser, sshPassword, sshKey string, sshPort int) (*sftp.Client, error) {
	// 连接Linux服务器
	conn, err := keyOrPwdConnectLinuxServer(sshHost, sshUser, sshPassword, sshKey, sshPort)
	if err != nil {
		fmt.Println("连接Linux服务器失败")
		panic(err)
	}
	// defer conn.Close()

	// 创建sftp会话
	client, err := sftp.NewClient(conn)
	if err != nil {
		return nil, err
	}
	//defer client.Close()
	return client, nil
}

func main() {
	// 连接sftp
	client, err := CreateSftp("192.168.3.121", "root", "123456", "", 22)
	if err != nil {
		return
	}
	defer client.Close()

	// 浏览服务器/home/devel目录
	//w := client.Walk("/home/devel")
	//for w.Step() {
	//	if w.Err() != nil {
	//		continue
	//	}
	//	fmt.Println(w.Path())
	//}

	// 在服务器创建文件
	//f, err := client.Create("/tmp/hello.txt")
	//if err != nil {
	//	panic(err)
	//}
	//_, err = f.Write([]byte("hello world!\n")) // 写入内容
	//if err != nil {
	//	panic(err)
	//}
	//f.Close()

	// 查看服务器的文件
	//fi, err := client.Lstat("/tmp/hello.txt")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%#v\n", fi)

	// 上传文件
	var localFilePath = "./test.txt" // 本地文件全路径
	var remoteDir = "/tmp"           // 服务器目录
	// 打开需要上传的本地文件
	srcFile, err := os.Open(localFilePath)
	if err != nil {
		panic(err)
	}
	defer srcFile.Close()
	// 获取需要上传的文件的名称
	var remoteFileName = path.Base(localFilePath)
	// 在服务器创建文件并打开
	dstFile, err := client.Create(path.Join(remoteDir, remoteFileName))
	if err != nil {
		panic(err)
	}
	defer dstFile.Close()
	// 写入内容
	buf := make([]byte, 1024)
	for {
		n, _ := srcFile.Read(buf) // 把文件中的内容读取到buf中，每次读取本地需要上传的文件1024字节内容
		// 当n=0时，证明本地需要上传的文件已经读取完毕了
		if n == 0 {
			break
		}
		dstFile.Write(buf)
	}
	fmt.Println("文件上传完毕")
}

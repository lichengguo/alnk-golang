# Go语言开发环境搭建

## 安装Go开发包

### 下载

[https://golang.google.cn/dl/](https://golang.google.cn/dl/)

![1561777435197](assets\1561777435197.png)

安装目录一定选一个好记的。

![1550236972659](https://www.liwenzhou.com/images/Go/install_go_dev/1550236972659.png)

安装完成后，输入 `go version`查看go版本号。

![1561777659521](assets\1561777659521.png)

## 配置GOPATH

![1561778001060](assets\1561778001060.png)

![1561778199604](assets\1561778199604.png)

详细步骤：

1. 在自己的电脑上新建一个目录 `D:\go`（存放我编写的Go语言代码）
2. 在环境变量里，新建一项：`GOPATH:D:\go`
3. 在 `D:\go`下新建三个文件夹，分别是：`bin`、`src`、`pkg`
4. 把 `D:\go\bin`这个目录添加到 `PATH`这个环境变量的后面
   1. Win7是英文的 `;`分隔
   2. Win10是单独一行
5. 你电脑上 `GOPATH`应该是有默认值的，通常是 `%USERPROFILE%/go`， 你把这一项删掉，自己按照上面的步骤新建一个就可以了。

![1561780163679](assets\1561780163679.png)

## GO语言项目结构

![1561780567275](assets\1561780567275.png)




## 下载并安装VS Code

### 下载VSCODE

[官方下载连接](https://code.visualstudio.com/Download)

### 安装

“下一步安装法”

### vscode进行go相关的基础设置

ctrl+b 打开左侧边栏
文件 --- 将文件夹添加到工作区... --- 可以同时打开多个不同的项目，对比查看代码很方便
ctrl+shift+x   打开扩展栏，搜索go插件并安装
ctrl+shift+p   选择 Go:Install/Update Tools  ， 选择列出的工具并确认安装
由于是从github下载，国内网络大概率会失败，这个时候在vscode打开一个终端，运行：
设置代理 go env -w GOPROXY=https://goproxy.cn
清空缓存 go clean --modcache

其他推荐的插件
vscode中文插件   Chinese
代码快捷运行插件  Code Runner

### 安装中文插件包

![1561780951477](assets\1561780951477.png)

### 安装Go扩展

![1561781081386](assets\1561781081386.png)




## 编译go build

使用 `go build`

1. 在项目目录下执行 `go build`
2. 在其他路径下执行 `go build`， 需要在后面加上项目的路径（项目路径从GOPATH/src后开始写起，编译之后的可执行文件就保存在当前目录下）
3. `go build -o hello.exe`

### go run

像执行脚本文件一样执行Go代码

### go install

`go install`分为两步：

    1. 先编译得到一个可执行文件
 	2. 将可执行文件拷贝到`GOPATH/bin`

### 交叉编译

Go支持跨平台编译

例如：在windows平台编译一个能在linux平台执行的可执行文件

```bash
SET CGO_ENABLED=0  // 禁用CGO
SET GOOS=linux  // 目标平台是linux
SET GOARCH=amd64  // 目标处理器架构是amd64
```

执行 `go build`

Mac平台交叉编译：

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
```

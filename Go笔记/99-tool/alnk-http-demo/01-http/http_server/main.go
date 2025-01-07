package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// http server

func f1(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close() //关闭连接

	b, err := ioutil.ReadFile("./xx.html")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err))) //如果找不到文件，直接把错误返回给浏览器
	}

	w.Write(b)
}

func f2(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close() //关闭连接

	// 获取请求的方式
	fmt.Println(r.Method) //GET POST等等

	// 请求头
	fmt.Println("Header: ", r.Header)
	//自己写的客户端访问的请求头
	//Header:  map[Accept-Encoding:[gzip] Content-Length:[33] Content-Type:[application/json] User-Agent:[Go-http-client/1.1]]
	//用浏览器访问的请求头
	//Header:  map[
	// Accept:[text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8]
	// Accept-Encoding:[gzip, deflate, br]
	// Accept-Language:[zh-CN,zh;q=0.9]
	// Connection:[keep-alive]
	// Cookie:[csrftoken=2tO2t7lIF0yq5bQhg7VQRORGmejfFtQhoJv4kSbzrDiExqLSuE5i198y31BzZAnp]
	// Upgrade-Insecure-Requests:[1]
	// User-Agent:[Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4)
	// AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36]]

	//GET请求
	if r.Method == "GET" {
		//Get请求，所有的参数都是放在URL上的，请求体中是没有数据的
		//发送请求的url http://127.0.0.1:9090/xx/?name=alnk&passwd=123456
		url := r.URL
		fmt.Println("url: ", url) //url:  /xx/?name=alnk&passwd=123456

		//获取请求中的参数
		//如果请求中一个参数有多个值怎么处理？
		queryParam := r.URL.Query()
		name := queryParam.Get("name")
		passwd := queryParam.Get("passwd")
		fmt.Printf("name:%s passwd:%s\n", name, passwd) //name:alnk passwd:123456

		//返回数据给客户端
		w.Write([]byte("get is ok!\n"))
	}

	//POST请求 这里也可以使用postman发送请求数据
	if r.Method == "POST" {
		//获取客户端传过来的表单数据，
		//注意如果打开获取表单数据，那么就把获取json数据注释掉，不然可能会出问题
		fmt.Printf("name:%s, %T\n", r.FormValue("name"), r.FormValue("name"))
		fmt.Printf("name:%s, %T\n", r.FormValue("passwd"), r.FormValue("passwd"))

		//获取客户端传过来的json数据
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("读取数据报错了哦", err)
			return
		}
		bToString := string(b) //可以通过内置的json包转换成对应的结构体
		fmt.Printf("%T\n", bToString)
		fmt.Printf("%v\n", bToString)

		//返回数据给客户端
		w.Write([]byte("post is ok!"))
	}
}

func main() {
	http.HandleFunc("/hello/", f1)
	http.HandleFunc("/xx/", f2)
	http.ListenAndServe("0.0.0.0:9090", nil)
}

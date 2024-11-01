package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

/*
客户端

[长连接和短连接]
https://www.zhihu.com/question/22677800
*/

func main() {
	//一.GET访问服务端
	// 1.参数直接写在Get方法中的URL里
	//resp, err := http.Get("http://127.0.0.1:9090/xx/?name=alnk&passwd=123456")
	//if err != nil {
	//	fmt.Printf("get url failed, err:%v\n", err)
	//	return
	//}
	//defer resp.Body.Close() //一定要记得关闭resp.Body
	//
	//// 获取服务端的返回
	//receByte, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Printf("read from resp.Body failed, err:%v\n", err)
	//	return
	//}
	//fmt.Println(string(receByte)) // get is ok!

	// 2.拼接参数发送Get请求 这个例子相比于上一个例子是更加灵活的，比如可以自定义请求头等
	//apiUrl := "http://127.0.0.1:9090/xx/" //URL
	//// URL参数
	//data := url.Values{} //可以看源码
	//data.Set("name", "alnk")
	//data.Set("passwd", "123")
	//
	//u, err := url.ParseRequestURI(apiUrl)
	//if err != nil {
	//	fmt.Printf("parse url requestUrl failed, err:%v\n", err)
	//	return
	//}
	//u.RawQuery = data.Encode() //url编码
	//
	//// 发送请求给服务端
	//resp, err := http.Get(u.String())
	//fmt.Printf("u.String(): %s\n", u.String()) //u.String():  http://127.0.0.1:9090/xx/?name=alnk&passwd=123
	//if err != nil {
	//	fmt.Printf("get failed , err:%v\n", err)
	//	return
	//}
	//defer resp.Body.Close() //关闭连接
	//
	//// 获取服务端返回值
	//b, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println("get resp failed ", err)
	//}
	//fmt.Println(string(b)) //get is ok!

	//
	//二.POST访问服务端
	url := "http://127.0.0.1:9090/xx/"
	// 表单数据
	//contentType := "application/x-www-form-urlencoded"
	//data := "name=alnk&passwd=123456"

	// json数据
	contentType := "application/json"
	data := `{"name":"alnk","passwd":"654321"}`

	// 发送post请求
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close() //一定要关闭连接

	// 获取服务端返回数据
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("get resp failed, ", err)
		return
	}
	fmt.Println(string(b)) //post is ok!

	//三.自定义client
	////URL
	//urlObj, err := url.Parse("http://127.0.0.1:9090/xx/")
	//if err != nil {
	//	fmt.Println("Parse URL failed, err:", err)
	//	return
	//}
	////data数据
	//data := url.Values{}
	//data.Set("name", "周林")
	//data.Set("passwd", "9000")
	//urlObj.RawQuery = data.Encode()
	////把data数据拼接到URL中
	//req, err := http.NewRequest("GET", urlObj.String(), nil)
	//
	////请求不是特别频繁，用完就关闭该链接。请求非常频繁，建议使用长连接
	//tr := &http.Transport{
	//	DisableKeepAlives: true, //请求头 禁用长连接，使用短连接
	//}
	//client := http.Client{
	//	Transport: tr,
	//}
	////发送请求
	//resp, err := client.Do(req)
	//if err != nil {
	//	fmt.Printf("get url failed, err:%v\n", err)
	//	return
	//}
	////一定要记得关闭resp.Body
	//defer resp.Body.Close()
	////读取服务端返回数据
	//b, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Printf("read resp.Body failed, err:%v\n", err)
	//	return
	//}
	//fmt.Println(string(b))
}

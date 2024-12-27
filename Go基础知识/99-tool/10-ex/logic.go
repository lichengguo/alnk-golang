package main

import (
	"10-ex/db"
	"10-ex/mylog"
	"fmt"
	"io/ioutil"
	"net/http"
)

// 终端输出日志
var conlog mylog.ConLog = mylog.NewConLog("info", false)

// 往文件写日志
var fileLog mylog.FileLog = *mylog.NewFileLog("info", "./logs", "log.log", 1*1024*1024, true)

// register 注册
func register(resp http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method == "GET" {
		b, err := ioutil.ReadFile("./statics/register.html")
		if err != nil {
			//如果找不到文件，直接把错误返回给客户端
			resp.Write([]byte(fmt.Sprintf("%v\n", err)))
			return
		}
		resp.Write(b)

	} else if r.Method == "POST" {
		//获取用户传递的账号密码
		userName := r.FormValue("username")
		passWord := r.FormValue("password")
		//检测合法性
		if userName != "" && passWord != "" {
			//1.到数据库查询看有没有该用户，有则提示已经注册
			//2.如果没有则提示注册成功，并且返回主页
			ret := db.QueryOne(userName)
			if ret {
				_, err := db.InsertOne(userName, passWord)
				if err != nil {
					fmt.Printf("database failed, err:%v\n", err)
					return
				}
				fileLog.Info("用户 %s 注册成功!", userName)
				b, err := ioutil.ReadFile("./statics/index.html")
				if err != nil {
					//如果找不到文件，直接把错误返回给客户端
					resp.Write([]byte(fmt.Sprintf("%v\n", err)))
					return
				}
				resp.Write(b)

			} else {
				fileLog.Error("用户 %s 已经注册过了! 注册失败!", userName)
				resp.Write([]byte("该用户已经注册过了"))
			}
		} else {
			resp.Write([]byte("用户名或密码不合法"))
		}
	}
}

// login 登录
func loginIndex(resp http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	//判断请求方式
	if r.Method == "GET" {
		b, err := ioutil.ReadFile("./statics/login.html")
		if err != nil {
			resp.Write([]byte(fmt.Sprintf("%v\n", err))) //如果找不到文件，直接把错误返回给客户端
			return
		}
		resp.Write(b)
	} else if r.Method == "POST" {
		//验证账号密码是否正确，正确就返回主页，否则返回账号密码有误
		//获取用户传递的账号密码
		userName := r.FormValue("username")
		passWord := r.FormValue("password")
		if userName == "" || passWord == "" {
			resp.Write([]byte("账号或者密码不能为空"))
			return
		}
		ret := db.QueryLogin(userName, passWord)
		if ret {
			b, err := ioutil.ReadFile("./statics/index.html")
			if err != nil {
				resp.Write([]byte(fmt.Sprintf("%v\n", err))) //如果找不到文件，直接把错误返回给客户端
				return
			}
			fileLog.Info("用户 %s 登录成功!", userName)
			resp.Write(b)
		} else {
			fileLog.Error("用户 %s 登录失败!", userName)
			fileLog.Info("用户 %s 登录失败!", userName)
			resp.Write([]byte("账号或者密码错误"))
		}

	} else {
		resp.Write([]byte("请求方式有问题"))
	}
}

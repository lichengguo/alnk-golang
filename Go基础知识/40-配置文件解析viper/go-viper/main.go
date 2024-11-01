package main

import (
	"fmt"
	"go-viper/conf"
	//_ "go-viper/conf"
)

func main() {
	fmt.Println(conf.Conf.Get("server.port"))
	fmt.Println(conf.Conf.Get("server.host"))
	fmt.Println(conf.Conf.Get("database.user"))
	fmt.Println(conf.Conf.Get("database.user"))
	fmt.Println(conf.Conf.Get("database.password"))
	fmt.Println(conf.Conf.Get("database.host"))
	fmt.Println(conf.Conf.Get("database.name"))
}

package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// 让map按照key排序打印

func main() {
	// rand.Seed(time.Now().UnixNano()) // 初始化随机数种子(1.20开始该方法已经被弃用)
	r := rand.New(rand.NewSource(time.Now().UnixNano())) // 初始化随机数种子

	var scoreMap = make(map[string]int, 100) // 定义一个map，key为string类型，值为int类型

	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("stu%02d", i) // 生成stu开头的字符串
		value := r.Intn(100)             // 生成0-99的随机整数
		scoreMap[key] = value
	}
	fmt.Println(scoreMap, len(scoreMap))

	// 取出map中所有的key存入切片keys
	var keys = make([]string, 0, 100)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	// fmt.Println(keys)
	// sort.Ints() // 按照整型排序
	sort.Strings(keys) // 按照字符串排序

	// 按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}

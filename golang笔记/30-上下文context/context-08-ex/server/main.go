package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// server端，随机出现慢响应

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// rand.Seed(time.Now().UnixNano())
	rand.New(rand.NewSource(time.Now().UnixNano()))
	number := rand.Intn(2) // 产生0和1的随机整数

	if number == 0 {
		time.Sleep(time.Second * 10) // 耗时10秒的慢响应
		fmt.Fprintf(w, "slow response")
		return
	}

	fmt.Fprintf(w, "quick response") // 正常响应
}

func main() {
	http.HandleFunc("/", indexHandler)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}

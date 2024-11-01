package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	f, err := os.OpenFile("./nginx.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 644)
	if err != nil {
		panic(err)
	}
	for {
		i := rand.Int63()
		msg := fmt.Sprintf("Nginx[%d]", i)
		fmt.Println(msg)
		fmt.Fprintln(f, msg)
		time.Sleep(time.Millisecond * 200)
	}
}

package main

import (
	"fmt"
	"sync"
)

/*
1、定义三个函数，分别可以打印cat、dog、fish
2、要求每个函数都起一个Goroutine
3、要求按照cat->dog->fish的顺序打印，每个50次
*/

var (
	wgp      sync.WaitGroup
	catChan  chan string
	dogChan  chan string
	fishChan chan string
)

func main() {
	catChan = make(chan string, 1)
	dogChan = make(chan string, 1)
	fishChan = make(chan string, 1)
	fishChan <- "fish"
	for i := 0; i < 50; i++ {
		print()
	}
}

func print() {
	wgp.Add(3)
	go cat()
	go dog()
	go fish()
	wgp.Wait()
}

func cat() {
	defer wgp.Done()
	for {
		select {
		case tag := <-fishChan:
			if tag == "fish" {
				fmt.Printf("cat ")
				catChan <- "cat"
				return
			}
		}
	}
}

func dog() {
	defer wgp.Done()
	for {
		select {
		case tag := <-catChan:
			if tag == "cat" {
				fmt.Print("dog ")
				dogChan <- "dog"
				return
			}
		}
	}

}

func fish() {
	defer wgp.Done()
	for {
		select {
		case tag := <-dogChan:
			if tag == "dog" {
				fmt.Print("fish ")
				fishChan <- "fish"
				return
			}
		}
	}
}

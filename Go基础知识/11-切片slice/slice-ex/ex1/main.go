package main

import "fmt"

// 切片练习

func main() {
	// ex1
	// s1 := make([]int, 4, 100)
	// s1[0] = 1
	// s1[1] = 2
	// fmt.Println(s1) //[1 2 0 0]
	// s2 := s1[:]     //再切片，切片的切片，底层数组还是同一个
	// //s2 := s1 //赋值 底层数组是同一个
	// s2[0] = 100
	// fmt.Println(s1)               //[100 2 0 0]
	// fmt.Println(s2)               //[100 2 0 0]
	// fmt.Println(cap(s1), cap(s2)) //100 100

	// ex2
	// s1 := make([]int, 4, 4)
	// s1[0] = 1
	// s1[1] = 2
	// fmt.Println(s1) //[1 2 0 0]
	// s2 := s1[:]     //切片的再切片 len:4 cap:4
	
	// //append()扩容以后会形成新的数组
	// s2 = append(s2, 100) //形成新的底层数组 [1 2 0 0 100] len:4 cap:8
	
	// fmt.Println(s1)               //[1 2 0 0]
	// fmt.Println(s2)               //[1 2 0 0 100]
	// fmt.Println(cap(s1), cap(s2)) //4 8

	// ex3
	// s1 := make([]int, 4, 4)
	// s1[0] = 1
	// s1[1] = 2
	// fmt.Println(s1)                         //[1 2 0 0]
	// s2 := s1[:2]                            //[1 2] len:2 cap:4
	// fmt.Println("==", s2, len(s2), cap(s2)) //[1 2] 2 4
	// s2 = append(s2, 100)                    //append()这里没有扩容，因为长度是2，容量是4，所以不会形成新的底层数组,此时底层数组是:[1 2 100]
	// fmt.Println(s1)                         //[1 2 100 0]
	// fmt.Println(s2)                         //[1 2 100] 这里s2的长度是3，不是4哦，所以是 [1 2 100]
	// fmt.Println(cap(s1), cap(s2))           //4 4
	// fmt.Println(len(s1), len(s2))           //4 3

	// ex4
	// var s1 = [...]int{1, 2, 3, 4, 5}
	// s2 := s1[2:] //[3 4 5] len:3 cap:3
	
	// s2 = append(s2, 100) //扩容了，产生新的底层数组 [3 4 5 100] len:4 cap:6
	
	// fmt.Println(s1)               //[1 2 3 4 5]
	// fmt.Println(s2)               //[3 4 5 100]
	// fmt.Println(cap(s1), cap(s2)) //5 6
	// fmt.Println(len(s1), len(s2)) //5 4

	// ex5
	// var arr = [...]int{1, 2, 3, 4, 5}
	// slice := arr[:2]                  //[1 2] len:2 cap:5
	// slice = append(slice, 6, 7)       //没扩容 切片是:[1 2 6 7] len:4 cap:5
	// slice[0] = 100                    //[100 2 6 7]
	// fmt.Println(arr)                  //[100 2 6 7 5]
	// fmt.Println(slice)                //[100 2 6 7]
	// fmt.Println(cap(arr), cap(slice)) //5 5
	// fmt.Println(len(arr), len(slice)) //5 4

	// ex6
	var arr = [...]int{1, 2, 3, 4, 5}
	slice := arr[:2] // [1 2] len:2 cap:5

	slice = append(slice, 6, 7, 8, 9, 10) //底层数组扩容了，切片是[1 2 6 7 8 9 10]
	slice[0] = 100                        //[100 2 6 7 8 9 10]

	fmt.Println(arr)                  //[1 2 3 4 5]
	fmt.Println(slice)                //[100 2 6 7 8 9 10]
	fmt.Println(cap(arr), cap(slice)) //5 10
	fmt.Println(len(arr), len(slice)) //5 7

}

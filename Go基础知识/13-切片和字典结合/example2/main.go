package main

import (
	"fmt"
)

//day03复习

func f1(a [3]int) {
	//Go语言中函数传递的都是值 （Ctrl+c  Ctrl+v）
	a[1] = 100 //此处修改的是副本
}

func main() {
	//var name string
	//name = "lixiang"
	//fmt.Println(name)

	//数组
	//var ages [30]int
	//ages = [30]int{1, 2, 3, 4, 5, 6}
	//fmt.Println(ages) //[1 2 3 4 5 6 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	//
	//var ages2 = [...]int{1, 2, 3, 4, 5, 7, 8, 90}
	//fmt.Println(ages2) //[1 2 3 4 5 7 8 90]
	//
	//var ages3 = [...]int{1: 100, 19: 200}
	//fmt.Println(ages3) //[0 100 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 200]

	//二维数组
	//var a1 = [...][2]int{
	//	[2]int{1, 2},
	//	[2]int{3, 4},
	//	[2]int{5, 6},
	//}
	//fmt.Println(a1) //[[1 2] [3 4] [5 6]]
	//多维数组只有最外层可以直接使用...
	//数组是值类型

	//x := [3]int{1, 2, 3}
	//y := x         //注意这里没有切片哦，相当于把x的值拷贝了一份给y
	//y[1] = 200     //修改的是副本，并不影响x
	//fmt.Println(x) //[1 2 3]
	//fmt.Println(y) //[1 200 3]
	//
	//fmt.Println(x) //[1 2 3]
	//f1(x)          //Go语言中函数传递的都是值 （Ctrl+c  Ctrl+v
	//fmt.Println(x) //[1 2 3]

	//切片slice
	//var s1 []int //没有分配内存 ==nil
	//fmt.Println(s1)
	//fmt.Println(s1 == nil)
	//s1 = []int{1, 2, 3}
	//fmt.Println(s1)
	//
	////make初始化分配内存
	//s2 := make([]bool, 2, 4)
	//fmt.Println(s2) //[false false]
	//s3 := make([]int, 0, 4)
	//fmt.Println(s3 == nil) //false

	//s1 := []int{1, 2, 3}
	//s2 := s1
	//var s3 = make([]int, 3, 3)
	//copy(s3, s1)
	//fmt.Println(s2) //[1 2 3]
	//s2[1] = 200
	//fmt.Println(s2) //[1 200 3]
	//fmt.Println(s1) //[1 200 3]
	//fmt.Println(s3) //[1 2 3]
	//var s1 []int //nil
	////s1 = make([]int, 1)
	////s1[0] = 100
	////fmt.Println(s1)
	//s1 = append(s1, 1) //append会自动初始化切片
	//fmt.Println(s1)

	//指针
	//Go里面的指针只能读不能修改，不能修改指针变量指向的地址
	//addr := "沙河"
	//addrP := &addr
	//fmt.Println(addrP)        //0xc0001041e0 内存地址
	//fmt.Printf("%T\n", addrP) //*string 一个string类型的指针
	//addrv := *addrP           //根据内存地址找值
	//fmt.Println(addrv)

	//map
	var m1 map[string]int
	fmt.Println(m1 == nil) //true
	m1 = make(map[string]int, 10)
	fmt.Println(m1 == nil) //false
	m1["lixiang"] = 100
	fmt.Println(m1)       //map[lixiang:100]
	fmt.Println(m1["ji"]) //0 如果key不存在返回的是value对应类型的零值
	//如果返回是bool值，我们通常用ok去接收他
	score, ok := m1["ji"]
	if !ok {
		fmt.Println("没有此key")
	} else {
		fmt.Println("分数是：", score)
	}
	delete(m1, "ji")     //删除的key不存在，什么都不做
	fmt.Println(len(m1)) //1
	delete(m1, "lixiang")
	fmt.Println(m1)        //map[]
	fmt.Println(m1 == nil) //false 虽然map里面没有值了，但是已经开辟了内存空间
	fmt.Println(len(m1))   //0
}

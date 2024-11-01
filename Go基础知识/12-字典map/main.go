package main

import "fmt"

// map

// make()函数和new()函数的区别
// make和new都是用来申请内存的
// new很少用，一般用来给基本数据类型申请内存，`string`、`int`,返回的是对应类型的指针(*string、*int)
// make是用来给`slice`、`map`、`chan`申请内存的，make函数返回的的是对应的这三个类型本身

func main() {
	var m1 map[string]int
	fmt.Println(m1 == nil)        // true 还没初始化，没有在内存中开辟空间
	m1 = make(map[string]int, 10) // 要估算好改map的容量，避免在程序运行期间再动态扩容，影响性能
	fmt.Println(m1 == nil)        // false 已经在内存中开辟空间了

	m1["age"] = 18
	m1["salary"] = 2000
	fmt.Println(m1)
	fmt.Println(m1["salary"])

	// 不存在的键
	// 约定成俗用ok接收返回的布尔值
	fmt.Println(m1["tom"]) // 如果不存在这个key，则拿到对应值类型的零值 0
	value, ok := m1["tom"]
	if !ok {
		fmt.Println("没有此key")
	} else {
		fmt.Println(value)
	}

	// map的遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	// 只遍历key
	for k := range m1 {
		fmt.Println(k)
	}

	// 只遍历value
	for _, v := range m1 {
		fmt.Println(v)
	}

	// 删除
	delete(m1, "salary")
	fmt.Println(m1)
	delete(m1, "tom") // 删除不存在的key，程序也不会报错
}

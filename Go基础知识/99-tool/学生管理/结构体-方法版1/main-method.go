package main

import (
	"fmt"
	"os"
)

// 学生信息结构体
type student struct {
	id   int
	name string
}

// 存储所有学生结构体
type class struct {
	Map map[int]*student
}

// 增加
func (c *class) addStudent() {
	var (
		id   int
		name string
	)
	fmt.Printf("input ID:")
	fmt.Scanln(&id)
	fmt.Printf("input name:")
	fmt.Scanln(&name)
	stu := &student{
		id:   id,
		name: name,
	}

	c.Map[id] = stu
}

// 查看
func (c *class) showStudent() {
	//fmt.Println(c.Map)
	for index, value := range c.Map {
		fmt.Printf("ID:%d Name:%s\n", index, value.name)
	}
}

// 删除
func (c *class) delStudent() {
	var id int
	fmt.Printf("del ID:")
	fmt.Scanln(&id)
	delete(c.Map, id)
	c.showStudent()
}

func main() {
	// 初始化
	c := &class{}
	c.Map = make(map[int]*student, 60)

	for {
		fmt.Printf("\t\t--------system admin --------")
		fmt.Println(`
			1.add student
			2.show students
			3.del student
			4.ex
 		`)

		var choice int
		fmt.Print("input ID: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			c.addStudent()
		case 2:
			c.showStudent()
		case 3:
			c.delStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("ID Error!")
		}
	}
}

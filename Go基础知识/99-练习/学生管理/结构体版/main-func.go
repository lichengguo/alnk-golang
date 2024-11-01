package main

import (
	"fmt"
	"os"
)

//结构体，存储单个学生信息
type student struct {
	id   int
	name string
}

//map类型，存储所有学生信息
var (
	allStudents map[int]*student
)

//构造函数
func newStudent(id int, name string) (stu *student) {
	stu = &student{
		id:   id,
		name: name,
	}
	return
}

func addStudent() {
	var (
		id   int
		name string
	)
	fmt.Print("请输入学生ID: ")
	fmt.Scanln(&id)
	fmt.Print("请输入学生名字: ")
	fmt.Scanln(&name)
	stu := newStudent(id, name)
	allStudents[id] = stu
}

func showStudents() {
	for k, v := range allStudents {
		fmt.Printf("ID:%d 姓名:%s\n", k, v.name)
	}
}

func deleteStudent() {
	var id int
	fmt.Print("请输入删除学生的ID: ")
	fmt.Scanln(&id)
	delete(allStudents, id)
}

func signOut() {
	os.Exit(1)
}

func main() {
	allStudents = make(map[int]*student, 50)
	for {
		fmt.Printf("\t--------学生管理系统--------")
		fmt.Println(`
		1.添加学生
		2.查看学生
		3.删除学生
		4.退出
 		`)
		var choice int
		fmt.Print("请输入选项: ")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			addStudent()
		case 2:
			showStudents()
		case 3:
			deleteStudent()
		case 4:
			signOut()
		default:
			fmt.Println("输入有误!")
		}
	}
}

package main

import (
	"fmt"
	"os"
)

/*
1、函数版学生管理系统
2、实现学生的查看、增加、删除
*/

//定义结构体，用来描述一个学生
type student struct {
	id int
	name string
}

//newStudent是student类型的构建函数
func newStudent(id int,name string) *student{
	return  &student{id:id, name:name}
}

//定义一个全局map用于保存学生信息
var (
	allStudent map[int] *student
)

//定义一个查看用户的函数
func showAllStudent(){
	//把所有的学生打印出来
	for k,v:=range allStudent{
		fmt.Printf("学号%d 姓名:%s\n",k,v.name)
	}
}

//定义一个增加用户的函数
func addStudent(){
	//向allStudent添加一个学生
	//1、创建一个学生
	//1.1获取用户输入
	var(
		id int
		name string
	)
	fmt.Print("请输入学生学号:")
	fmt.Scan(&id)
	fmt.Print("请输入学生姓名:")
	fmt.Scan(&name)
	//1.2造学生(调用newStudent的构建函数)
	newStu := newStudent(id,name)
	//2、追加到allStudent这个map中
	allStudent[id] = newStu
}

//定义一个删除用户的函数
func deleteStudent() {
	//1、请用户输入要删除学生的学号


	var(
		deleteID int
	)
	fmt.Printf("请输入学生学号:")
	fmt.Scan(&deleteID)
	//2、去allStudent删除学生
	delete(allStudent,deleteID)

}

func main() {
	allStudent = make(map[int]*student,48)  //开辟内存空间
	for {
		//输出功能列表
		fmt.Println("========欢迎使用学生管理系统======")
		fmt.Println("1、查看学生")
		fmt.Println("2、增加学生")
		fmt.Println("3、删除学生")
		fmt.Println("4、退出系统")

		//获取用户的输入
		fmt.Print("请输入你要使用的功能:")
		//定义一个变量用于保存用户的输入
		var choice int
		fmt.Scan(&choice)
		//fmt.Println(choice)

		//根据用户的选择执行相应的函数
		switch choice {
		case 1:
			showAllStudent()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("该功能不存在，请重新选择")
		}
	}
}
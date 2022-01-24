package main

import (
	"fmt"
	"os"
)

type student2 struct {
	id int
	name string
}

var allStudent2=make(map[int] *student2)

func addStudent2() {
	var (
		id int
		name string
	)
	fmt.Print("请输入学生ID:")
	fmt.Scan(&id)
	fmt.Print("请输入学生名称:")
	fmt.Scan(&name)
	allStudent2[id]=&student2{id,name}
}

func delStudent(){
	var delKey int
	fmt.Print("请输入你要删除的学生ID号:")
	fmt.Scan(&delKey)
	delete(allStudent2,delKey)

}

func selectStudent(){
	for k,v:=range allStudent2{
		fmt.Printf("id是:%d名称是:%s\n",k,v.name)
	}
}

func main(){
	for {
		fmt.Println("======欢迎使用学生管理系统======")
		fmt.Println("1、增加学生")
		fmt.Println("2、查看学生")
		fmt.Println("3、删除学生")
		fmt.Println("4、退出系统")

		var choise int
		fmt.Print("请输入你要使用的功能选项:")
		fmt.Scan(&choise)

		switch choise {
		case 1:
			addStudent2()
		case 2:
			selectStudent()
		case 3:
			delStudent()
		case 4:
			os.Exit(1)
		}
	}

}
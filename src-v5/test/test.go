package main

import "fmt"

func main() {
	//创建一个map，值是string，key是interface{}
	m1 := make(map[string]interface{})
	//空接口可以储存任何类型的值
	m1["name"] = "王林"
	m1["age"] = 1800
	m1["hobby"] = [...]string{"吃", "喝", "玩", "乐"}
	fmt.Println(m1)
}

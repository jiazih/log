package main

import (
	"fmt"
	"runtime"
)

func Test(){
	pc,file,line,ok := runtime.Caller(1)
	if !ok{
		fmt.Println("runtime.Caller failed")
	}
	fmt.Println(runtime.FuncForPC(pc).Name())
	fmt.Println(file)
	fmt.Println(line)
}



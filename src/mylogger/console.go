package mylogger

import (
	"fmt"
	"time"
)
type LogLevel uint8

//定义日志级别
const (
	DEBUG  LogLevel = iota
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

// Logger 结构体具有多个日志方法
type Logger struct {
	/*
	1、需要给日志定义一个开关级别，比如测试从debug向上都输入，但是生产只输出error以上的
	2、定义一个Level级别类型是LogLevel,LogLevel实际就是uint8
	 */
	Level LogLevel
}

/*
	1、NewLog return一个Logger结构体
	2、传入一个level字符串类型，用于对比输出的日志级别
*/
func NewLog(level string)Logger{
	return Logger{}
}

//Debug 日志方法,在调用方法的时候传进来一个字符串类型，然后使用日志方法打印出来
func (l Logger)Debug(msg string){
	now := time.Now()
	fmt.Printf("[%s] [DEBUG] %s\n",now.Format("2006-01-02 15:04:05"),msg)
}

//Trace 日志方法,在调用方法的时候传进来一个字符串类型，然后使用日志方法打印出来
func (l Logger)Trace(msg string){
	now := time.Now()
	fmt.Printf("[%s] [TRACE] %s\n",now.Format("2006-01-02 15:04:05"),msg)
}

//Info 日志方法,在调用方法的时候传进来一个字符串类型，然后使用日志方法打印出来
func (l Logger)Info(msg string){
	now := time.Now()
	fmt.Printf("[%s] [INFO] %s\n",now.Format("2006-01-02 15:04:05"),msg)
}

//Warning 日志方法,在调用方法的时候传进来一个字符串类型，然后使用日志方法打印出来
func (l Logger)Warning(msg string){
	now := time.Now()
	fmt.Printf("[%s] [WARNING] %s\n",now.Format("2006-01-02 15:04:05"),msg)
}

//Error 日志方法,在调用方法的时候传进来一个字符串类型，然后使用日志方法打印出来
func (l Logger)Error(msg string){
	now := time.Now()
	fmt.Printf("[%s] [ERROR] %s\n",now.Format("2006-01-02 15:04:05"),msg)
}

//Fatal 日志方法,在调用方法的时候传进来一个字符串类型，然后使用日志方法打印出来
func (l Logger)Fatal(msg string){
	now := time.Now()
	fmt.Printf("[%s] [FATAL] %s\n",now.Format("2006-01-02 15:04:05"),msg)
}
package mylogger

import (
	"errors"
	"fmt"
	"strings"
	"time"
)
type LogLevel uint8

/*
	1、定义日志级别，因为要设置日志开关，需要做一个比较，string本身无法比较，所以需要转换一下
 */
const (
	UNKNOWN LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

//转换函数，转换成上面自定义的日志级别，进行比较
func parseLogLevel(s string)(LogLevel,error){
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG,nil
	case "trace":
		return TRACE,nil
	case "info":
		return INFO,nil
	case "warning":
		return WARNING,nil
	case "error":
		return ERROR,nil
	case "fatal":
		return FATAL,nil
	default:
		//如果用户传递的值不再范围内，那么返回一个默认值
		err := errors.New("无效的日志级别")
		return UNKNOWN,err
	}
}


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
	3、这里是字符串没办法直接比较，需要转换类型才可以，类型转换函数为parseLogLevel
*/
func NewLog(levelStr string)Logger{
	/*
		1、将用户输入的字符串类型的日志级别，转换为我们自定义的LogLevel级别，这样就可以进行比较了
		2、返回两个值，一个是自定义的日志级别，一个是错误
	 */
	level,err := parseLogLevel(levelStr)
	if err != nil{
		panic(err)
	}
	//将结构体return回去，就是多了一个Level参数，类型是level，level是我们的自定义的日志级别，方便下面进行判断
	return Logger{
		Level: level,
	}
}

/*
	1、判断方法，接收一个LogLevel类型的值，就是我们自定义的值
	2、使用用户传入的值和自定义的值做对比(用户传入的值已经被转变为我们自定义的Loglevel值了)
 */
func (l Logger) enable(logLevel LogLevel)bool{
	/*
		1、进行判断，返回bool值
		2、loglevel是我们自定义的值，而l.Level是用户传入的值
		3、进行判断，返回true或者false
		4、这里的logLevel是Logger方法传进来的值
	 */
	return logLevel >= l.Level
}

//Debug 日志方法,在调用方法的时候传进来一个字符串类型，然后使用日志方法打印出来
func (l Logger)Debug(msg string){
	//1、进行判断，如果用户传入的值大于等于自定义的值，那么执行if里面的东西，否则不执行
	if l.enable(DEBUG){
		now := time.Now()
		fmt.Printf("[%s] [DEBUG] %s\n",now.Format("2006-01-02 15:04:05"),msg)
	}
}

//Trace 日志方法,在调用方法的时候传进来一个字符串类型，然后使用日志方法打印出来
func (l Logger)Trace(msg string){
	if l.enable(TRACE){
		now := time.Now()
		fmt.Printf("[%s] [TRACE] %s\n",now.Format("2006-01-02 15:04:05"),msg)
	}
}

//Info 日志方法,在调用方法的时候传进来一个字符串类型，然后使用日志方法打印出来
func (l Logger)Info(msg string){
	if l.enable(INFO){
		now := time.Now()
		fmt.Printf("[%s] [INFO] %s\n",now.Format("2006-01-02 15:04:05"),msg)
	}
}

//Warning 日志方法,在调用方法的时候传进来一个字符串类型，然后使用日志方法打印出来
func (l Logger)Warning(msg string){
	if l.enable(WARNING){
		now := time.Now()
		fmt.Printf("[%s] [WARNING] %s\n",now.Format("2006-01-02 15:04:05"),msg)
	}
}

//Error 日志方法,在调用方法的时候传进来一个字符串类型，然后使用日志方法打印出来
func (l Logger)Error(msg string){
	if l.enable(ERROR){
		now := time.Now()
		fmt.Printf("[%s] [ERROR] %s\n",now.Format("2006-01-02 15:04:05"),msg)
	}
}

//Fatal 日志方法,在调用方法的时候传进来一个字符串类型，然后使用日志方法打印出来
func (l Logger)Fatal(msg string){
	if l.enable(FATAL){
		now := time.Now()
		fmt.Printf("[%s] [FATAL] %s\n",now.Format("2006-01-02 15:04:05"),msg)
	}
}
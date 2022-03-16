package mylogger

import (
	"fmt"
	"time"
)
type LogLevel uint8




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

/*
	1、因为要打印时间每一个都需要修改很麻烦，所以就定义了一个log函数，每个方法直接调用log函数即可
	2、log函数的参数接收一个LogLevel类型的变量，msg是main函数调用的时候传入的变量信息，
 */
func log(lv LogLevel,formant string,a...interface{}){
	//接收一个字符串， a...是不定参数，传不传值都可以
	msg := fmt.Sprintf(formant,a...)
	now := time.Now()
	funcName, fileName ,lineNo := getInfo(3)
	//这里有一个getLogstring参数，因为LogLevel是我们自定义的变量类型，底层是uint8类型的，不是字符串，所以这里需要再转换为字符串类型
	fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n",now.Format("2006-01-02 15:04:05"),getLogstring(lv),fileName,funcName,lineNo,msg)
}

//Debug 日志方法,在调用方法的时候传进来一个字符串类型，然后使用日志方法打印出来
func (l Logger)Debug(format string,a...interface{}){
	//1、进行判断，如果用户传入的值大于等于自定义的值，那么执行if里面的东西，否则不执行
	if l.enable(DEBUG){
		//接收调用者传进来的格式化的值
		log(DEBUG,format,a...)
	}
}

//Trace 日志方法,在调用方法的时候传进来一个字符串类型，然后使用日志方法打印出来
func (l Logger)Trace(msg string){
	if l.enable(TRACE){
		log(TRACE,msg)
	}
}

//Info 日志方法,在调用方法的时候传进来一个字符串类型，然后使用日志方法打印出来
func (l Logger)Info(msg string){
	if l.enable(INFO){
		log(INFO,msg)
	}
}

//Warning 日志方法,在调用方法的时候传进来一个字符串类型，然后使用日志方法打印出来
func (l Logger)Warning(msg string){
	if l.enable(WARNING){
		log(WARNING,msg)
	}
}

//Error 日志方法,在调用方法的时候传进来一个字符串类型，然后使用日志方法打印出来
func (l Logger)Error(msg string){
	if l.enable(ERROR){
		log(ERROR,msg)
	}
}

//Fatal 日志方法,在调用方法的时候传进来一个字符串类型，然后使用日志方法打印出来
func (l Logger)Fatal(msg string){
	if l.enable(FATAL){
		log(FATAL,msg)
	}
}
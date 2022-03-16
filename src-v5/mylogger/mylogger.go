package mylogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
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

func getLogstring(lv LogLevel)string{
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	}
	return "DEBUG"
}


//用于获取代码的文件名、函数名以及行数
func getInfo(skip int)(funcName,fileName string,lineNo int){
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok{
		fmt.Printf("runtime.Caller() failed\n")
		return
	}
	//return三个值，分别是文件名称，调用runtime.Caller的行号以及函数名称
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	funcName = strings.Split(funcName,".")[1]
	return
}


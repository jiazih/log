package main

import (
	"mylogger"
	"time"
)

func main(){
	/*
	1、调用mylogger.NewLog()生成一个对象，这个对象具有多个日志方法
	2、给mylogger传入一个字符串，告诉他我要打印哪个级别的日志
	3、mylogger包中的Newlog是一个构造函数
	 */
	log := mylogger.NewLog("fatal")


	//调用日志的主函数,log中有如下几个方法
	for {
		log.Debug("这是一条Debug日志")
		log.Trace("这是一条Trace日志")
		log.Info("这是一条Info日志")
		log.Warning("这是一条Warning日志")
		log.Error("这是一条Error日志")
		log.Fatal("这是一条Fatal日志")
		time.Sleep(time.Second * 2)
	}
}

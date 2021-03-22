/***************************
@File        : mylogger.go
@Time        : 2020/12/03 15:51:22
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : go实现logger日志库
****************************/

package mylogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

type Level uint16


type Log interface {
	Debug(msg string, a ...interface{})
	Info(msg string, a ...interface{})
	Waring(msg string, a ...interface{})
	Error(msg string, a ...interface{})
	Fatal(msg string, a ...interface{})
}

const (
	UNKNOW Level = iota
	DEBUG
	INFO
	WARING
	ERROR
	FATAL
)

// 解析日志级别
func ParseLevel(s string) (Level, error) {
	
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "info":
		return INFO, nil
	case "waring":
		return WARING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKNOW, err
	}
}

// NewLog 构造函数
func NewLog(levelStr string) Logger {
	level, err := ParseLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return Logger{
		level: level,
	}
}

// 构造函数 文件输入
func NewFileLogger(level, fp, fn string, maxSize int64) *FileLogger {
	logLevel, err := ParseLevel(level)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		level:       logLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}
	err = fl.initFile() // 按照文件路径文件名打开文件
	if err != nil {
		panic(err)
	}
	return fl
}

func getInfo(n int) (funcName, fileName string, lineNo int) {
	pc, file, line, ok := runtime.Caller(n)
	if !ok {
		fmt.Println("runtime.caller() failed \n")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	return funcName, fileName, line
}

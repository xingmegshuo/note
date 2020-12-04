/***************************
@File        : test.go
@Time        : 2020/12/03 16:05:28
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 测试logerr库
****************************/

package main

import (
	"note/go/basis/demo/mylogger"
	"time"
)

var log mylogger.Log

func main() {
	// log := mylogger.NewFileLogger("debug", "./go/basis/demo/mylogger_test/", "today", 1*1024)
	log = mylogger.NewLog("DEBUG")
	for {
		log.Debug("hhhh")
		log.Info("ok")
		time.Sleep(time.Second * 2)

		log.Waring("ah")
		time.Sleep(time.Second * 2)

		id := 10010
		name := "李小龙"
		log.Error("error,id:%d,name:%s", id, name)
		log.Fatal("what fuck")
		time.Sleep(time.Second * 5)
	}
}

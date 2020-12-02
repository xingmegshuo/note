/***************************
@File        : goroutines.go
@Time        : 2020/12/02 16:38:22
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : go 程
****************************/

package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// 启动一个新的go程运行
	go say("world")
	say("hello")
}

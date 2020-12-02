/***************************
@File        : default-selection.go
@Time        : 2020/12/02 16:54:08
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : select 分支的default分支
****************************/

package main

import (
	"fmt"
	"time"
)

// 当 select 中的其它分支都没有准备好时，default 分支就会执行。

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.Tick(500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("boom!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

/***************************
@File        : range.go
@Time        : 2020/11/27 13:17:13
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        :   go 语言 range 方法学习
****************************/

package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d=%d\n", i, v)
	}
}

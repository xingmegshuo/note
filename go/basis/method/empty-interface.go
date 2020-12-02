/***************************
@File        : empty-interface.go
@Time        : 2020/12/02 15:09:03
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 空接口
****************************/

package main

import "fmt"

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v,%T)\n", i, i)
}

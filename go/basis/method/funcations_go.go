/***************************
@File        : funcations_go.go
@Time        : 2020/12/02 14:54:18
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : go 函数
****************************/
package main

import "fmt"

// 求和函数
func add(x, y int) int {
	return x + y
}

// 多返回值
func swap(x, y string) (string, string) {
	return y, x
}
func main() {
	fmt.Println(add(22, 12))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

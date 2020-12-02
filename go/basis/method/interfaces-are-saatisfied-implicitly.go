/***************************
@File        : interfaces-are-saatisfied-implicitly.go
@Time        : 2020/12/02 14:53:04
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 接口与隐式实现
****************************/

package main

import "fmt"

type I interface {
	M()
}
type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}

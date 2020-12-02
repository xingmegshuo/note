/***************************
@File        : interface-values.go
@Time        : 2020/12/02 14:56:47
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 接口值
****************************/

package main

import "fmt"

type I interface {
	M()
}
type T struct {
	S string
}

func (t *T) M() {
	// 底层值为nil
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I
	// 这是接口值的两个类型
	// i = &T{"Hello"}
	// describe(i)
	// i.M()

	// i = F(math.Pi)
	// describe(i)
	// i.M()

	// nil 值
	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"Hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v,%T)\n", i, i)
}

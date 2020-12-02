/***************************
@File        : indirection.go
@Time        : 2020/12/02 13:55:43
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 方法与指针重定向
****************************/

// 带指针参数的函数必须接收一个指针而指针为接收者的方法被调用时，接收者技能为值又能为指针

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y float64
}

func (v *Vertex) Scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

// 接收一个值作为参数的函数，必须接收一个指定类型的值
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)

	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}

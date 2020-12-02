/***************************
@File        : methods-pointers.go
@Time        : 2020/12/02 13:46:14
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 指针接收者
****************************/

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func (v *Vertex) Scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.x, v.y)
	fmt.Println(v.Abs())
	fmt.Println(v.x, v.y)
}

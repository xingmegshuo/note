/***************************
@File        : methods-with-pointer-receivers.go
@Time        : 2020/12/02 14:52:19
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 选择值或指针作为接收者
****************************/

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

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("before scaling: %+v, Abs:%v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("before scaling: %+v, Abs:%v\n", v, v.Abs())
}

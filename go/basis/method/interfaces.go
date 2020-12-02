/***************************
@File        : interfaces.go
@Time        : 2020/12/02 14:52:36
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 接口
****************************/
package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := Myfloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat 实现了Abser
	a = &v // a *Vertex 实现了Abser
	// v 是一个Vertex 所以不能实现Abser
	// a = v
	fmt.Println(a.Abs())
}

type Myfloat float64

func (f Myfloat) Abs() float64 {
	if f < 0 {
		return float64(f)
	}
	return float64(f)
}

type Vertex struct {
	x, y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

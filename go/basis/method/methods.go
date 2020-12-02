/***************************
@File        : methods.go
@Time        : 2020/12/02 13:35:45
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : go 方法
****************************/

// go没有类，不过可以为结构体类型定义方法

// abs 方法拥有一个名为v，类型为Vertex的接收者

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

// 方法即函数
func Abs(v Vertex) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

// 接收者的类型定义和方法声明必须在同一包内；不能为内建类型声明方法

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(Abs(v))
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

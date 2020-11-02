package main

import (
	"fmt"
	"math"
)

// 牛顿法求平方根

func Sqrt(F float64) float64 {
	x := 1.0
	for math.Abs(x*x-F) > 1e-10 {
		x -= (x*x - F) / (2 * x)
		fmt.Println("每次循环的x=", x)

	}
	return x
}

func main() {
	fmt.Println("牛顿法求平方根:Sqrt(10)=", Sqrt(10))
	fmt.Println("库函数求平方根:math.Sqrt(10)=", math.Sqrt(10))
}

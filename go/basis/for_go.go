package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) {
	z := float64(1)
	for {
		if z*z < x {
			z -= (z*z - x) / (2 * z)
		}
		fmt.Println(z, z*z)
	}
}

func main() {
	// sum := 0
	// for i := 0; i < 15; i++ {
	// 	sum += i
	// 	fmt.Println(sum, i)
	// }
	sqrt(20)
	fmt.Println(math.Sqrt(20))

}

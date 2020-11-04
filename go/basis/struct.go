// 定义变量一组字段,相当于定义几个相同类型的变量

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}

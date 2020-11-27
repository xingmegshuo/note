// 定义变量一组字段,相当于定义几个相同类型的变量

package main

import "fmt"

// 结构体字段使用.来访问

type Vertex struct {
	X int
	Y int
}

var (
	v1    = Vertex{1, 2}  // 创建一个Vertex 类型的结构体
	v2    = Vertex{X: 1}  // Y:0 被隐式的赋予
	v3    = Vertex{}      // XY都为0
	point = &Vertex{1, 2} // 指针
)

func main() {
	fmt.Println(Vertex{1, 2})
	v := Vertex{1, 2}
	fmt.Println(v.X)
	// 结构体指针
	p := &v
	p.X = 1e9
	fmt.Println(v)
	fmt.Println(v1, v2, v3, point)
}

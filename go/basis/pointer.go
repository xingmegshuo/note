// go 指针
package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i // p 是 i 的指针

	fmt.Println(*p) // 通过指针读取i的值

	*p = 22
	fmt.Println(i) //指针修改后i的值  通过指针i的值被修改

	p = &j
	*p = *p / 37 // 指针已经等于j 通过指针进行除法运算
	fmt.Println(j)
}

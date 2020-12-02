package main

import "fmt"

// [n]T 表示拥有n个类型值的数组

func main() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	// 数组切片
	var s []int = primes[1:4]
	fmt.Println(s)
	// 切片文法
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)
	z := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(z)
	// 切片的默认行为
	s = primes[1:4]
	fmt.Println(s)
	s = primes[:2]
	fmt.Println(s)
	s = primes[1:]
	fmt.Println(s)
	// 截取切片使长度为 0
	s = primes[:0]
	printSlice(s)
	// 扩展长度
	s = s[:4]
	printSlice(s)
	// 舍弃前两个值
	s = s[:2]
	printSlice(s)
	// nil 切片
	var h []int
	fmt.Println(h, len(h), cap(h))
	if h == nil {
		fmt.Println("nil!")
	}
}

// 切片的长度与容量 len 长度 cap容量
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

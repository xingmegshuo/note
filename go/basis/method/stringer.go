/***************************
@File        : stringer.go
@Time        : 2020/12/02 15:33:25
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : stringer
****************************/

// stringer 是一个可以用字符串描述自己的类型. fmt 包 都通过此接口来输出值

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

// 一个stringer 练习

type IPAddr [4]byte

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)

	host := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range host {
		fmt.Printf("%v :%v\n", name, ip)
	}
}

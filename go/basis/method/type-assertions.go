/***************************
@File        : type-assertions.go
@Time        : 2020/12/02 15:12:01
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 类型断言
****************************/

package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string) // 判断一个接口值是否保存了一个特定类型可通过此方法返回值ok来判断
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	do(21)
	do("hello")
	do(true)

}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v \n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long \n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

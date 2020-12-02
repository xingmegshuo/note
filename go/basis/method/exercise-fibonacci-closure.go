/***************************
@File        : exercise-fibonacci-closure.go
@Time        : 2020/12/02 12:18:08
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 斐波纳契闭包
****************************/

// 斐波那契数，也称之为斐波那契数列,又称黄金分割数列、费波那西数列、费波拿契数、费氏数列，指的是这样一个数列：
// 0、1、1、2、3、5、8、13、21、……在数学上，斐波纳契数列以如下被以递归的方法定义：
// F0=0，F1=1，Fn=Fn-1+Fn-2（n>=2，n∈N*），用文字来说，就是斐波那契数列列由
//  0 和 1 开始，之后的斐波那契数列系数就由之前的两数相加。

package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int {
	back1, back2 := 0, 1
	return func() int {
		// 记录back1
		temp := back1
		// 重新赋值
		back1, back2 = back2, (back1 + back2)
		return temp
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

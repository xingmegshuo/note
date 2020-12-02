/***************************
@File        : range-and-close.go
@Time        : 2020/12/02 16:46:37
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : range 和 close
****************************/

package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y

	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	fmt.Println(cap(c))
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

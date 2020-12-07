/***************************
@File        : exerise_channel.go
@Time        : 2020/12/07 12:42:28
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : channel练习
****************************/

package main

import (
	"fmt"
	"sync"
)

// 一个线程生成数字，发送到通道, 另一个线程从通道a取出数字计算平方，放到通道2

var wg sync.WaitGroup

// 生成数字
func make_num(a chan int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		a <- i // 把i放进通道a
		fmt.Println("通道a:", a)
	}
	close(a)
}

// 取出数字
func get_num(a, b chan int) {
	defer wg.Done()
	for { // 从通道a取值
		x, ok := <-a
		if !ok {
			break
		}
		b <- x * x // 向通道b赋值
		fmt.Println("通道b:", b)
	}
	close(b)
}

func main() {
	c := make(chan int, 10) // 创建两个通道
	d := make(chan int, 10)
	wg.Add(2)
	go make_num(c)
	go get_num(c, d)
	wg.Wait()
	for ret := range d {
		fmt.Println(ret)
	}

}

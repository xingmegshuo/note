/***************************
@File        : channels.go
@Time        : 2020/12/02 16:41:26
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 信道
****************************/

package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把和送入c
}

// 写入chanl
func send(mes string, ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- mes
}

// 读出数据
func read(ch chan string) {
	x, _ := <-ch
	switch x {
	case "hhh":
		fmt.Println("ok")
	}
}

func main() {
	// s := []int{7, 2, 8, -9, 4, 0}

	// c := make(chan int)
	// go sum(s[:len(s)/2], c)
	// go sum(s[len(s)/2:], c)
	// x, y := <-c, <-c
	// fmt.Println(x, y, x+y)

	// 带缓冲的信道

	// ch := make(chan int, 2)
	// ch <- 1
	// ch <- 2
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	ch := make(chan string)
	sock := 0
	if sock == 0 {
		go send("hhh", ch)
	}
	for {

		if sock == 0 {
			fmt.Println("first")
			go read(ch)

		}
		fmt.Println(sock)
		sock = sock + 1
		time.Sleep(time.Second * 3)
	}
}

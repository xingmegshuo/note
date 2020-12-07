/***************************
@File        : goroutines_proces.go
@Time        : 2020/12/04 17:56:53
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : go 程 的最大管理
****************************/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func a() {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		fmt.Printf("A:%d\n", i)
	}
}
func b() {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		fmt.Printf("B:%d\n", i)
	}
}

func main() {
	runtime.GOMAXPROCS(2)         //设置cpu逻辑核心数
	fmt.Println(runtime.NumCPU()) //输出cpu逻辑核心数
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}

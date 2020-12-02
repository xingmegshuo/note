/***************************
@File        : mutex-counter.go
@Time        : 2020/12/02 16:58:29
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 互斥锁
****************************/

package main

import (
	"fmt"
	"sync"
	"time"
)

// safecounter 的并发使用是安全的

type safecounter struct {
	v   map[string]int
	mux sync.Mutex
}

// inc 增加给定key的计数器的值
func (c *safecounter) Inc(key string) {
	c.mux.Lock()
	// Lock 之后同一时刻只有一个go程访问
	c.v[key]++
	c.mux.Unlock()
}

func (c *safecounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := safecounter{v: make(map[string]int)}
	for i := 0; i < 100; i++ {
		go c.Inc("somekey")
	}
	time.Sleep(time.Second)
	fmt.Println(c)
	fmt.Println(c.Value("somekey"))

}

package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Go runs on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
	// switch 求值顺序
	fmt.Println("When's Sunday?")
	today := time.Now().Weekday()
	switch time.Sunday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
		fmt.Println(time.Sunday)
	}
	// 没有条件的switch
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("good morring!")
	case t.Hour() <= 17:
		fmt.Println("good afternoon!")
	default:
		fmt.Println(t.Hour())
		fmt.Println("good evening.")
	}
	// what is defer 它会将函数推迟到外层函数返回后执行
	defer fmt.Println("world")
	fmt.Println("hello")
	// defer example
	fmt.Println("start cycle")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

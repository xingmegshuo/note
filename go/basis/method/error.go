/***************************
@File        : error.go
@Time        : 2020/12/02 15:53:36
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 错误
****************************/

package main

import (
	"fmt"
	"time"
)

type MyError struct {
	when time.Time
	what string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v,%s", e.when, e.what)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

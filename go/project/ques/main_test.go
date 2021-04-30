/***************************
@File        : server_test.go
@Time        : 2021/04/08 15:19:05
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : server test
****************************/

package main

import (
	// _ "ques/server"
	"testing"
	"time"
)

// api 测试
func TestRun(t *testing.T) {
	time.Sleep(time.Second * 1)
	t.Log("test start.")
	d1 := []string{"a", "b", "a", "d", "b"}
	t.Log(d1)
	r, e := Run(d1)
	t.Log(r, e)
	if r[2] == true && r[4] == true && e == nil {
		t.Log("test case success.")
	} else {
		t.Error("test case Error")
	}
	d2 := []string{"早", "早上", "早上好", "早上", "b"}
	t.Log(d2)
	r, e = Run(d2)
	t.Log(r, e)
	if r[3] == true && r[4] == true && e == nil {
		t.Log("test case success.")
	} else {
		t.Error("test case Error")
	}
}

/***************************
@File        : reader.go
@Time        : 2020/12/02 15:57:22
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : reader
****************************/

package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n=%v err=%v b=%v", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

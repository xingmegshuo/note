/***************************
@File        : _reader.go
@Time        : 2020/12/02 16:07:27
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : reader 练习
****************************/

package main

import "golang.org/x/tour/reader"

type MyReader struct {
}

func (r MyReader) Read(b []byte) (int, error) {
	b[0] = 'A'
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}

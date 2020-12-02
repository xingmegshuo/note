/***************************
@File        : images.go
@Time        : 2020/12/02 16:27:21
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 图像
****************************/

package main

import (
	"fmt"
	"image"
)

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}

/***************************
@File        : ex.go
@Time        : 2021/04/25 14:22:47
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 测试
****************************/

package main

import (
	"fmt"
	"log"
	"regexp"
)

func main() {
	pattern := regexp.MustCompile(fmt.Sprintf(`(%s)\.(\d+)\.key`, "search"))
	matched := pattern.FindAllStringSubmatch("search.2.key", -1)
	log.Println(matched)
}

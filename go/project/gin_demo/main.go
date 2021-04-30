/***************************
@File        : main.go
@Time        : 2020/12/02 17:57:46
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : gin 使用
****************************/

package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "hello, geek")
	})
	r.POST("/login", func(c *gin.Context) {
		c.String(200, "what fuck")
	})
	r.Run(":8080")
}

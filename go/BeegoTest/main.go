package main

import (
	_ "BeegoTest/models"
	_ "BeegoTest/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}


package routers

import (
	"BeegoTest/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
}

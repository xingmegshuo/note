package controllers

import (
	// "BeegoTest/models"

	// "github.com/astaxie/beego/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

type LoginController struct {
	beego.Controller
}

func (c *MainController) Get() {
	// o := orm.NewOrm()
	// user := models.User{}
	// user.Name = "111"
	// user.Pwd = "222"
	// _, err := o.Insert(&user)
	// if err != nil {
	// 	logs.Info("插入失败", err)
	// }
	// user.Id = 1
	// err := o.Read(&user)
	// if err != nil {
	// 	logs.Info("查询失败", err)
	// 	return
	// }
	// logs.Info("查询成功", user)
	// c.Data["data"] = "SmallAnt"
	c.TplName = "main.html"
}

func (c *MainController) Post() {
	name := c.GetString("userName")
	pws := c.GetString("password")
	logs.Info("数据", name, pws)
	c.TplName = "main.html"
}

func (c *LoginController) Get() {
	c.TplName = "index.tpl"
}

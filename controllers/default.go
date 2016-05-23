package controllers

import "github.com/astaxie/beego"

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.tpl"
}

func (c *MainController) Create() {
	c.TplName = "create.tpl"

	if c.Ctx.Input.Method() == "POST" {
		sessName := c.GetString("name")
		sessPass := c.GetString("password")
		c.SetSession(sessName, sessPass)
		c.Redirect("/room/"+sessPass, 302)
	}
}

func (c *MainController) Join() {
	c.TplName = "join.tpl"
}

func (c *MainController) Room() {
	c.TplName = "room.tpl"
	c.Data["Pass"] = c.Ctx.Input.Param(":id")
}

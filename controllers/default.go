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
		c.Redirect("/room/"+sessName, 302)
	}
}

func (c *MainController) Join() {
	c.TplName = "join.tpl"

	if c.Ctx.Input.Method() == "POST" {
		sessName := c.GetString("name")
		sessPass := c.GetString("password")
		c.SetSession(sessName, sessPass)
		c.Redirect("/room/"+sessName, 302)
	}
}

func (c *MainController) Room() {
	c.TplName = "room.tpl"

	roomName := c.Ctx.Input.Param(":id")
	sess := c.GetSession(roomName)
	if sess == nil {
		c.Redirect("/", 302)
	}
	c.Data["Pass"] = roomName
}

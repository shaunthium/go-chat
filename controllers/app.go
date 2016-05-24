package controllers

import "github.com/astaxie/beego"

var (
	messages = make([]string, 20)
)

type MainController struct {
	beego.Controller
}

func (controller *MainController) Get() {
	controller.TplName = "index.tpl"
}

func (controller *MainController) Create() {
	controller.TplName = "create.tpl"

	if controller.Ctx.Input.Method() == "POST" {
		sessName := controller.GetString("name")
		sessPass := controller.GetString("password")
		controller.SetSession(sessName, sessPass)
		controller.Redirect("/room/"+sessName, 302)
	}
}

func (controller *MainController) Join() {
	controller.TplName = "join.tpl"

	if controller.Ctx.Input.Method() == "POST" {
		sessName := controller.GetString("name")
		sessPass := controller.GetString("password")
		controller.SetSession(sessName, sessPass)
		controller.Redirect("/room/"+sessName, 302)
	}
}

func (controller *MainController) Room() {
	controller.TplName = "room.tpl"

	roomName := controller.Ctx.Input.Param(":id")
	sess := controller.GetSession(roomName)
	if sess == nil {
		controller.Redirect("/", 302)
	}
	controller.Data["Pass"] = roomName
}

func (controller *MainController) Messages() {
	controller.TplName = "room.tpl"
	if controller.Ctx.Input.Method() == "POST" {
		messages = append(messages, controller.GetString("message"))
	}
	if controller.Ctx.Input.Method() == "GET" {
		controller.Data["json"] = messages
		controller.ServeJSON()
	}
}
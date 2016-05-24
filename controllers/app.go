package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

var (
	sendChannel    = make(chan string)
	receiveChannel = make(chan string)
	messages       = make([]string, 20)
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
	// go chat()
}

func (controller *MainController) Messages() {
	controller.TplName = "room.tpl"
	if controller.Ctx.Input.Method() == "POST" {
		// sendChannel <- controller.GetString("input")
		messages = append(messages, controller.GetString("input"))
	}
	if controller.Ctx.Input.Method() == "GET" {
		fmt.Printf("messages is:" + messages[0])
		controller.Data["json"] = messages
		controller.ServeJSON()
	}
}

func chat() {
	for {
		select {
		case receivedMessage := <-sendChannel:
			fmt.Println("received message:" + receivedMessage)
			receiveChannel <- receivedMessage
		}
	}
}

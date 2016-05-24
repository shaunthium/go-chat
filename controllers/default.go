package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

var (
	sendChannel    = make(chan string)
	receiveChannel = make(chan string)
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

	if controller.Ctx.Input.Method() == "GET" {
		roomName := controller.Ctx.Input.Param(":id")
		sess := controller.GetSession(roomName)
		if sess == nil {
			controller.Redirect("/", 302)
		}
		controller.Data["Pass"] = roomName
		go chat()
	}

	if controller.Ctx.Input.Method() == "POST" {
		sendChannel <- controller.GetString("input")
		return
	}
}

func chat() {
	for {
		select {
		case receivedMessage := <-sendChannel:
			fmt.Println("received message:" + receivedMessage)
		case sentMessage := <-receiveChannel:
			fmt.Println("sent message:" + sentMessage)
		}
	}
}

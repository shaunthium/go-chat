package controllers

import "github.com/astaxie/beego"

const (
	httpMethodPOST = "POST"
)

var (
	messages = make([]string, 20)
	rooms    = make([]string, 20)
)

// MainController is the main controller for the app
type MainController struct {
	beego.Controller
}

func contains(rooms []string, roomName string) bool {
	for _, name := range rooms {
		if name == roomName {
			return true
		}
	}
	return false
}

// Get method for index page
func (controller *MainController) Get() {
	controller.TplName = "index.html"
	beego.ReadFromRequest(&controller.Controller)
}

// Create method for creating rooms
func (controller *MainController) Create() {
	controller.TplName = "create.html"

	if controller.Ctx.Input.Method() == httpMethodPOST {
		sessionName := controller.GetString("name")
		controller.SetSession(sessionName, "")
		controller.Redirect("/room/"+sessionName, 302)
	}
}

// Join method for joining rooms
func (controller *MainController) Join() {
	controller.TplName = "join.html"

	if controller.Ctx.Input.Method() == httpMethodPOST {
		roomName := controller.GetString("name")
		if contains(rooms, roomName) {
			controller.SetSession(roomName, "")
			controller.Redirect("/room/"+roomName, 302)
		} else {
			flash := beego.NewFlash()
			flash.Error("No such room found!")
			flash.Store(&controller.Controller)
			controller.Redirect("/", 302)
		}
	}
}

// Room method representing a chat room
func (controller *MainController) Room() {
	controller.TplName = "room.html"

	roomName := controller.Ctx.Input.Param(":id")
	sess := controller.GetSession(roomName)
	if sess == nil {
		controller.Redirect("/", 302)
	}
	controller.Data["Pass"] = roomName
}

// Messages method, used as an API
func (controller *MainController) Messages() {
	controller.TplName = "room.html"
	if controller.Ctx.Input.Method() == httpMethodPOST {
		messages = append(messages, controller.GetString("message"))
	}
	if controller.Ctx.Input.Method() == "GET" {
		controller.Data["json"] = messages
		controller.ServeJSON()
	}
}

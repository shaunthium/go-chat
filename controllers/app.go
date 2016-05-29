package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

const (
	METHOD_POST = "POST"
	METHOD_GET  = "GET"
)

var (
	_     = fmt.Printf // To prevent compiler complaining about unused imports
	data  = make(map[string]RoomData)
	rooms = make([]string, 20) // Contains name of every currently active room
)

// Message struct represents a data type containing data about each message
type Message struct {
	Sender  string `json:"sender"`
	Content string `json:"content"`
}

// RoomData represents data associated with a room
type RoomData struct {
	Messages        []Message
	RemainingPeople int
}

// MainController is the main controller for the app
type MainController struct {
	beego.Controller
}

// Get method for index page
func (controller *MainController) Get() {
	controller.TplName = "index.html"
	beego.ReadFromRequest(&controller.Controller)
}

// Create method for creating rooms
func (controller *MainController) Create() {
	controller.TplName = "create.html"

	if controller.Ctx.Input.Method() == METHOD_POST {
		roomName := controller.GetString("room-name")
		username := controller.GetString("username")
		temp := make(map[string]interface{})
		temp["roomName"] = roomName
		temp["username"] = username
		controller.SetSession(roomName, temp)
		rooms = append(rooms, roomName)
		data[roomName] = RoomData{make([]Message, 0, 0), 1}
		controller.Redirect("/room/"+roomName, 302)
	}
}

// Join method for joining rooms
func (controller *MainController) Join() {
	controller.TplName = "join.html"

	if controller.Ctx.Input.Method() == METHOD_POST {
		roomName := controller.GetString("name")
		if contains(rooms, roomName) {
			temp := make(map[string]interface{})
			username := controller.GetString("username")
			temp["roomName"] = roomName
			temp["username"] = username

			// Set session with necessary data
			controller.SetSession(roomName, temp)

			// Increment number of people in the room
			tempRoomData := data[roomName]
			tempRoomData.RemainingPeople = tempRoomData.RemainingPeople + 1
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
	session := controller.GetSession(roomName)
	temp := make(map[string]interface{})
	sessionRoomName := ""
	if session != nil {
		temp = session.(map[string]interface{})
		sessionRoomName = temp["roomName"].(string)
	}
	if session == nil || sessionRoomName != roomName {
		// Check that session exists
		flash := beego.NewFlash()
		flash.Error("If you know this room exists, please click the 'Join Room' button.")
		flash.Store(&controller.Controller)
		controller.Redirect("/", 302)
	} else {
		// Leave room when user navigates away from page
		defer leaveRoom(roomName)
		controller.Data["username"] = temp["username"].(string)
		controller.Data["roomName"] = temp["roomName"].(string)
	}
}

// Messages method, used as an API
func (controller *MainController) Messages() {
	controller.TplName = "room.html"
	roomName := controller.GetString("roomName")
	if controller.Ctx.Input.Method() == METHOD_POST {
		// Add message to history
		sender := controller.GetString("sender")
		content := controller.GetString("content")
		message := Message{sender, content}
		temp := data[roomName]
		temp.Messages = append(temp.Messages, message)
		data[roomName] = temp
	}
	if controller.Ctx.Input.Method() == METHOD_GET {
		controller.Data["json"] = data[roomName].Messages
		controller.ServeJSON()
	}
}

package controllers

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
)

var (
	_ = fmt.Printf
	_ = strconv.Atoi
)

// Checks if the supplied room name exists in the 'room' array
func contains(rooms []string, roomName string) bool {
	for _, name := range rooms {
		if name == roomName {
			return true
		}
	}
	return false
}

// Creates a new room array without the specified room name
// and returns the new array
func deleteRoom(rooms []string, roomName string) []string {
	newRooms := make([]string, len(rooms))
	counter := 0
	for _, val := range rooms {
		if val != roomName {
			newRooms[counter] = val
			counter++
		}
	}
	return newRooms
}

// Redirects to index page and shows flash error
func redirectWithError(controller *MainController,
	errorMessage string,
	path string) {

	flash := beego.NewFlash()
	flash.Error(errorMessage)
	flash.Store(&controller.Controller)
	controller.Redirect(path, 302)
}

// Leaves room specified by room name
func leaveRoom(roomName string) {
	temp := data[roomName]
	if temp == nil {
		return
	}
	temp.RemainingPeople = temp.RemainingPeople - 1
	// Delete room data and room if no people are left in the room
	if temp.RemainingPeople == 0 {
		data[roomName] = nil
		rooms = deleteRoom(rooms, roomName)
	}
}

package controllers

// Checks if the supplied room name exists in the 'room' array
func contains(rooms []string, roomName string) bool {
	for _, name := range rooms {
		if name == roomName {
			return true
		}
	}
	return false
}

func leaveRoom(roomName string) {
	temp := data[roomName]
	temp.RemainingPeople = temp.RemainingPeople - 1
	// Delete room data if no people are left in the room
	if temp.RemainingPeople == 0 {
		delete(data, roomName)
	}
}

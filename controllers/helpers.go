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

func leaveRoom(roomName string) {
	temp := data[roomName]
	temp.RemainingPeople = temp.RemainingPeople - 1
	// Delete room data and room if no people are left in the room
	if temp.RemainingPeople == 0 {
		delete(data, roomName)
		rooms = deleteRoom(rooms, roomName)
	}
}

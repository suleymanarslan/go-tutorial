package system

import "fmt"

type Room struct {
	clients []*client
	Id      string
}

var Rooms map[string]Room

func InitiateRooms() {
	fmt.Println("InitiateRooms")
	Rooms = make(map[string]Room)
}

func getClientsFromRoom(room string) []*client {
	var retval Room
	retval = Rooms[room]
	return retval.clients
}

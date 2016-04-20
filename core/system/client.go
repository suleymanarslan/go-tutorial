package system

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/gorilla/websocket"
)

var rooms map[string]Room

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 1024 * 1024
)

type client struct {
	ws   *websocket.Conn
	send chan []byte
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  maxMessageSize,
	WriteBufferSize: maxMessageSize,
}

func InitiateRooms() {
	fmt.Println("InitiateRooms")
	rooms = make(map[string]Room)
}

func ServeWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside web sock")
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	c := &client{
		send: make(chan []byte, maxMessageSize),
		ws:   ws,
	}

	Hub.register <- c

	go c.writePump()
	c.readPump()
}

func (c *client) readPump() {
	fmt.Println("inside read")

	defer func() {
		Hub.unregister <- c
		c.ws.Close()
	}()

	var message Message

	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error {
		c.ws.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {

		err := c.ws.ReadJSON(&message)

		if message.MessageType == "createjoin" {
			room := rooms[message.Room]
			fmt.Println("room count", len(room.Clients))

			var returnMessage Message
			if reflect.DeepEqual((Room{}), room) {
				var clients []string

				room = Room{Id: message.Room, Clients: clients}
				room.Clients = append(room.Clients, "suleyman")
				returnMessage.MessageType = "created"
				returnMessage.Room = message.Room
				rooms[message.Room] = room
				fmt.Println("room count2", len(room.Clients))
			} else {
				if len(room.Clients) < 2 {
					Hub.broadcast <- "join"
					fmt.Println("room count3", len(room.Clients))
					room.Clients = append(room.Clients, "test")
					returnMessage.MessageType = "joined"
					returnMessage.Room = message.Room
					rooms[message.Room] = room
				}
			}

			c.send <- []byte(returnMessage.MessageType)
		} else if message.MessageType == "gotusermedia" {

			Hub.broadcast <- message.MessageType

		}

		if err != nil {
			fmt.Println(err)
		}

	}
}

func (c *client) writePump() {
	ticker := time.NewTicker(pingPeriod)

	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.write(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func (c *client) write(mt int, message []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	fmt.Println("sending....", string(message))
	return c.ws.WriteMessage(mt, message)
}

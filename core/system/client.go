package system

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 1024 * 1024
)

type client struct {
	ws   *websocket.Conn
	send chan []byte
	id   string
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  maxMessageSize,
	WriteBufferSize: maxMessageSize,
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
    
	go c.writePump()
	c.readPump()
}

func (c *client) readPump() {
	fmt.Println("inside read")

	defer func() {
		///Hub.unregister <- c
		c.ws.Close()
	}()

	var message Message
	var baseMessage RawMessage

	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error {
		c.ws.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {

		err := c.ws.ReadJSON(&baseMessage)

		if err != nil {
			fmt.Println("Panic in getting base message.")
			panic(err)
		}

		fmt.Println(baseMessage.BaseMessageType)

		if baseMessage.BaseMessageType == "common" {
			data := baseMessage.Message
			Source := (*json.RawMessage)(&data)
			err := json.Unmarshal(*Source, &message)

			if err != nil {
				panic(err)
			}
			
			if message.MessageType == "createjoin" {
				reg := register{ c, message.Room  }
				Hub.register <- reg
			} else if message.MessageType == "gotusermedia" {
				brTo := broadcastTo{"gotusermedia", message.Room, *c}
				Hub.broadcastto <- brTo
			} else if message.MessageType == "bye" {
				unreg := unregister{c, message.Room }
				Hub.unregister <- unreg
			}
		} else if baseMessage.BaseMessageType == "rtc" {

			if err != nil {
				fmt.Println("Panic in getting rtc message.")
				panic(err)
			}

			brTo := broadcastTo{string(baseMessage.Message), baseMessage.Room, *c}
			Hub.broadcastto <- brTo
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
	fmt.Println("sending....", string(message)+c.id)
	return c.ws.WriteMessage(mt, message)
}

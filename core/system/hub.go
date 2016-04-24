package system

import "fmt"

type hub struct {
	// Registered clients
	clients map[*client]bool

	// Inbound messages
	broadcast chan string

	broadcastto chan broadcastTo

	broadcastin chan broadcastIn

	// Register requests
	register chan *client

	// Unregister requests
	unregister chan *client

	join chan string

	content string

	room string
}

type broadcastIn struct {
	message string
	room    string
}

type broadcastTo struct {
	message string
	room    string
	client  client
}

var Hub = hub{
	broadcast:   make(chan string),
	broadcastin: make(chan broadcastIn),
	broadcastto: make(chan broadcastTo),
	register:    make(chan *client),
	unregister:  make(chan *client),
	clients:     make(map[*client]bool),
	content:     "",
	room:        "",
}

func (h *hub) Run() {
	for {
		select {
		case c := <-h.register:
			h.clients[c] = true
			c.send <- []byte(h.content)
			break

		case c := <-h.unregister:
			_, ok := h.clients[c]
			if ok {
				delete(h.clients, c)
				close(c.send)
			}
			break
		case m := <-h.broadcast:
			h.content = m
			h.broadcastMessage()
			break
		case t := <-h.broadcastto:
			h.content = t.message
			h.broadcastMessageTo(t.room, t.client)
			break
		case s := <-h.broadcastin:
			h.content = s.message
			h.broadcastMessageIn(s.room)
			break
		}
	}
}

func (h *hub) broadcastMessageTo(room string, cli client) {
	var clients []*client
	clients = getClientsFromRoom(room)
	fmt.Println("broadcastMessageTo", len(clients))
	for _, c := range clients {
		fmt.Println("client id>>>>", cli.id)
		fmt.Println("c id>>>>", c.id)
		fmt.Println(cli.ws == c.ws)
		if cli.ws != c.ws {
			fmt.Println("broadcastMessageTo", c.id)
			select {
			case c.send <- []byte(h.content):
				break
			default:
				close(c.send)
				delete(h.clients, c)
			}
		}
	}
}

func (h *hub) broadcastMessageIn(room string) {
	var clients []*client
	clients = getClientsFromRoom(room)
	fmt.Println("broadcastMessagein", len(clients))

	for _, c := range clients {
		fmt.Println("broadcastMessageIn", c.id)

		select {

		case c.send <- []byte(h.content):
			break

		// We can't reach the client
		default:
			close(c.send)
			delete(h.clients, c)
		}
	}
}

func (h *hub) broadcastMessage() {
	fmt.Println("broadcastingggg")
	for c := range h.clients {
		select {
		case c.send <- []byte(h.content):
			break

		// We can't reach the client
		default:
			close(c.send)
			delete(h.clients, c)
		}
	}
}

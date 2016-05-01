package system

import ("fmt")

type hub struct {
	roomClients map[string][]*client

	broadcast chan string

	broadcastto chan broadcastTo

	broadcastin chan broadcastIn

	register chan register

	unregister chan unregister

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

type unregister struct{
	leavingClient *client
	clientRoom string
}

type register struct{
	newClient *client
	clientRoom string
	}

var Hub = hub{
	broadcast:   make(chan string),
	broadcastin: make(chan broadcastIn),
	broadcastto: make(chan broadcastTo),
	register:    make(chan register),
	unregister:  make(chan unregister),
	roomClients: make(map[string][]*client),
	content:     "",
	room:        "",
}

func (h *hub) Run() {
	for {
		select {
		case c := <-h.register:
		if h.roomClients[c.clientRoom] == nil {
			h.content = "created"
			h.roomClients[c.clientRoom] =append(h.roomClients[c.clientRoom], c.newClient)
			c.newClient.send <- []byte(h.content)
		} else{
			h.content = "join"
			h.broadcastMessageIn(c.clientRoom)
			h.content = "joined"
			h.roomClients[c.clientRoom] =append(h.roomClients[c.clientRoom], c.newClient)
			c.newClient.send <- []byte(h.content)
		}
			break

		case c := <-h.unregister:
			_, ok := h.roomClients[c.clientRoom]
			if ok {
				for index := 0; index < len(h.roomClients[c.clientRoom]); index++ {
					if h.roomClients[c.clientRoom][index] == c.leavingClient {
						h.roomClients[c.clientRoom] = append(h.roomClients[c.clientRoom][:index], h.roomClients[c.clientRoom][index + 1:]...)
					}
				}
				close(c.leavingClient.send)
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
	clients = h.roomClients[room]
	fmt.Println("broadcastMessageTo", len(clients))
	for _, c := range clients {
		if cli.ws != c.ws {
			fmt.Println("broadcastMessageTo", c.id)
			select {
			case c.send <- []byte(h.content):
				break
			default:
				close(c.send)
			}
		}
	}
}

func (h *hub) broadcastMessageIn(room string) {
	var clients []*client
	clients =  h.roomClients[room]
	for _, c := range clients {
		select {

		case c.send <- []byte(h.content):
			break

		default:
			close(c.send)
		}
	}
}

func (h *hub) broadcastMessage() {
	for _, value := range h.roomClients {
		for _, c:= range value{
		select {
		case c.send <- []byte(h.content):
			break
		default:
			close(c.send)
		}			
		}

	}
}

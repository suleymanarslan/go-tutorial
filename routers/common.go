package routers

import (
	"hoditgo/controllers"
	"hoditgo/core/authentication"
	"hoditgo/core/system"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetHelloRoutes(router *mux.Router) *mux.Router {
	router.Handle("/test/hello",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.HelloController),
		)).Methods("GET")

	return router
}

func SetAppCommonRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/create-user", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/signal", system.ServeWs)
	return router
}

/*func (wsh wsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        //log.Println(err)
        return
    }


    var message system.Message

    for {

        err := conn.ReadJSON(&message)

        if err != nil {
            fmt.Printf("err", err)
        }

        if message.MessageType == "createjoin" {

            fmt.Printf("Room: ", message.Room)

            room := rooms[message.Room]
            var returnMessage system.Message
            if reflect.DeepEqual((system.Room{}), room){
                fmt.Printf("creating new room: ", message.Room)
                var clients []string

                room = system.Room{Id: message.Room, Clients: clients}
                room.Clients = append(room.Clients,"suleyman")
                returnMessage.MessageType = "created"
                returnMessage.Room = message.Room
                rooms[message.Room] = room
            } else {
                fmt.Printf("room count", len(room.Clients))
                if len(room.Clients) < 2 {
                    fmt.Printf("joining room : ", message.Room)
                    room.Clients = append(room.Clients,"test")
                    returnMessage.MessageType = "joined"
                    returnMessage.Room = message.Room
                    rooms[message.Room] = room

                }
            }

            conn.WriteJSON(returnMessage)
        }

        else if message.MessageType == "gotusermedia" {

        }

        if  err != nil {
            return
        }
    }

}*/

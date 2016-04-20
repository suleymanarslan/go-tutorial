package main

import (
	"hoditgo/core/system"
	"hoditgo/routers"
	"hoditgo/settings"
	"net/http"

	"github.com/codegangsta/negroni"
)

func main() {
	go system.Hub.Run()
	settings.Init()
	router := routers.InitRoutes()

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	n := negroni.Classic()
	n.UseHandler(router)
	system.InitiateRooms()
	http.ListenAndServe(":5000", n)
}

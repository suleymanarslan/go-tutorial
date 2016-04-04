package main

import (
	"hoditgo/routers"
	"hoditgo/settings"
	"github.com/codegangsta/negroni"
	"net/http"
)

func main() {
	settings.Init()
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe(":5000", n)
}

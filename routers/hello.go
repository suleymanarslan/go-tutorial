package routers

import (
	"hoditgo/controllers"
	"hoditgo/core/authentication"
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

func SetAppCommonRoutes(router *mux.Router)  *mux.Router {
		router.HandleFunc("/create-user", controllers.CreateUser).Methods("POST")
			return router
}

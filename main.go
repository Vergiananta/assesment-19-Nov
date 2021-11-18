package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"superapp/connect"
)

type superApp struct {
	connect connect.Connect
	router  *mux.Router
}

func (app *superApp) run(){
	h := app.connect.ApiServer([]string{})
	log.Println("Listening on", h)
	NewAppRouter(app).InitMainRoutes()
	err := http.ListenAndServe(h, app.router)
	if err != nil {
		log.Fatalln(err)
	}
}

func NewSuperApp() *superApp {
	r := mux.NewRouter()
	var appConnect = connect.NewConnect()

	return &superApp{
		connect: appConnect,
		router:  r,
	}
}

func main() {
	NewSuperApp().run()
}
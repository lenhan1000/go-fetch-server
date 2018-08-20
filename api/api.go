package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
}

//func (a *Server) Initialize()

func (a *Server) Run(host string) {
	fmt.Printf("Server started on %v\n", host)
	log.Fatal(http.ListenAndServe(host, a.Router))
}

func (a *Server) Init() {
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *Server) setRouters() {}

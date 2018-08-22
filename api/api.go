package api

import (
	"fmt"
	"go-fetch-server/api/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
)

type Server struct {
	Router *mux.Router
	DB     *mgo.Database
}

func (s *Server) Run(host string) {
	fmt.Printf("Server started on %v\n", host)
	log.Fatal(http.ListenAndServe(host, s.Router))
}

func (s *Server) Init(server string, db string) {
	s.Router = mux.NewRouter()
	session, err := mgo.Dial(server)
	if err != nil {
		log.Fatal(err)
	}
	s.DB = session.DB(db)
	s.setRouters()
}

func (s *Server) setRouters() {
	s.Get("/users", s.GetAllUsers)
	s.Post("/users", s.CreateUser)
}

// Get wraps the router for GET method
func (s *Server) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	s.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (s *Server) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	s.Router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (s *Server) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	s.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (s *Server) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	s.Router.HandleFunc(path, f).Methods("DELETE")
}

func (s *Server) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	controller.FindAllUsers(s.DB, w, r)
}

func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	controller.CreateUser(s.DB, w, r)
}

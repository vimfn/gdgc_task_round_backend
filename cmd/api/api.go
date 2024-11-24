package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) APIServer {
	return APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()

	// technically I should use `/api/v1` as subroute path prefix (best practices fyi.), but cuz the task didn't mention, i'm skipping this.

	// register user routes

	// register product routes

	// TODO: do a normal healthcheck
	router.PathPrefix("/").Handler()

	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, router)
}

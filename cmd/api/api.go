package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"vitshop.vimfn.in/services/product"
	"vitshop.vimfn.in/utils"
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

	// register product routes
	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore)
	productHandler.RegisterRoutes(router)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		utils.WriteJSON(w, http.StatusOK,
			map[string]interface{}{
				"data": "hi seniors 👻",
			})
	}).Methods("GET")

	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, router)
}

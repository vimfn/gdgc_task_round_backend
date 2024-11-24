package user

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"vitshop.vimfn.in/types"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRouter(router *mux.Router) {
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	fmt.Println("69")
}

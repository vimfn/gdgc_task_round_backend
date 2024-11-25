package user

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"vitshop.vimfn.in/configs"
	"vitshop.vimfn.in/services/auth"
	"vitshop.vimfn.in/types"
	"vitshop.vimfn.in/utils"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRouter(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var user types.NewUserPayload
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	// note: the default behaviour is to yeild error for same email (which is fine)
	if _, err := h.store.GetUserByEmail(user.Email); err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", user.Email))
		return
	}

	// bcrypt hashpass
	hashedPass, err := auth.HashPassword(user.Password)
	if err != nil {
		// not really a good idea to send in err stack to user but ok for now
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// make sure to do migrations before testing
	// else it will fail not being able to find the vitshop.users table
	err = h.store.CreateNewUser(types.User{
		Email:    user.Email,
		Password: hashedPass,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	// ik i should make a seperate payload but they are necessarily the same
	var user types.NewUserPayload
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	u, err := h.store.GetUserByEmail(user.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid email or password"))
		return
	}

	if !auth.ComparePasswords(u.Password, []byte(user.Password)) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}

	secret := []byte(configs.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, u.ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}

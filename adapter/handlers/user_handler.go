package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/chahar4/aura/core/services"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

type registerRequest struct {
	Username string
	Email    string
	Password string
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	request := registerRequest{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors.New("Invalid request"))
		return
	}

	if err := h.userService.Register(context.Background(), request.Username, request.Password, request.Email); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("registered")
}

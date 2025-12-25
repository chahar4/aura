package handlers

import (
	"context"
	"encoding/json"
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
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	request := registerRequest{}
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Request invalid", http.StatusBadRequest)
		return
	}

	if err := h.userService.Register(context.Background(), request.Username, request.Password, request.Email); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	request := loginRequest{}
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Request invalid", http.StatusBadRequest)
	}

	token, err := h.userService.Login(context.Background(), request.Email, request.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})

}

type forgotPasswordRequest struct {
	Email string `json:"email"`
}

func (h *UserHandler) ForgotPasswordSend(w http.ResponseWriter, r *http.Request) {
	request := forgotPasswordRequest{}
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Request invalid", http.StatusBadRequest)
	}

	if err := h.userService.ForgotPasswordSendCode(context.Background(), request.Email); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "email sended"})
}

type forgotPasswordRecoveryRequest struct {
	Email       string `json:"email"`
	NewPassword string `json:"new_password"`
	Token       string `json:"token"`
}

func (h *UserHandler) ForgotPasswordRecovery(w http.ResponseWriter, r *http.Request) {
	request := forgotPasswordRecoveryRequest{}
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Request invalid", http.StatusBadRequest)
	}
	if err := h.userService.ForgotPasswordRecovery(context.Background(), request.Email, request.NewPassword, request.Token); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message" : "password changed"})

}

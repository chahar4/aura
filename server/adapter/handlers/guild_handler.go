package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/chahar4/aura/core/services"
)

type GuildHandler struct {
	guildService *services.GuildService
}

func NewGuildHandler(guildService *services.GuildService) *GuildHandler {
	return &GuildHandler{
		guildService: guildService,
	}
}

type addGuildRequest struct {
	Name    string  `json:"name"`
	Profile *string `json:"profile"`
}

func (h *GuildHandler) AddGuild(w http.ResponseWriter, r *http.Request) {
	request := addGuildRequest{}
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Request invalid", http.StatusBadRequest)
		return
	}
	userIDString := r.Context().Value("userID").(string)
	userID ,err := strconv.Atoi(userIDString)
	if err != nil{
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if err := h.guildService.AddGuild(context.Background(),uint(userID), request.Name, request.Profile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}


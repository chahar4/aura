package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/chahar4/aura/core/services"

	"github.com/go-chi/chi/v5"
)

type ChannelHandler struct {
	channelService *services.ChannelService
}

func NewChannelHandler(channelService *services.ChannelService) *ChannelHandler {
	return &ChannelHandler{
		channelService: channelService,
	}
}

type addChannelRequest struct {
	Name string `json:"name"`
}

func (h *ChannelHandler) AddChannel(w http.ResponseWriter, r *http.Request) {
	request := addChannelRequest{}
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	groupChannelIDStr := chi.URLParam(r, "id")
	if groupChannelIDStr == "" {
		http.Error(w, "Request invalid2", http.StatusBadRequest)
		return
	}
	groupChannelID, err := strconv.Atoi(groupChannelIDStr)
	if err != nil {
		http.Error(w, "Request invalid3", http.StatusBadRequest)
		return
	}

	if err := h.channelService.AddChannel(r.Context(), request.Name, uint(groupChannelID)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *ChannelHandler) GetAllChannelsByGroupChannelID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	groupChannelIDStr := chi.URLParam(r, "id")
	if groupChannelIDStr == "" {
		http.Error(w, "Request invalid", http.StatusBadRequest)
		return
	}
	groupChannelID, err := strconv.Atoi(groupChannelIDStr)
	if err != nil {
		http.Error(w, "Request invalid", http.StatusBadRequest)
		return
	}

	channels, err := h.channelService.GetAllChannelsByGroupChannelID(r.Context(), uint(groupChannelID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println(channels)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(channels)
}

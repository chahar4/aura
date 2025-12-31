package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/chahar4/aura/core/services"
)

type GuildMemberHandler struct {
	guildMemberService *services.GuildMemberService
}

func NewGuildMemberHandler(guildMemberService *services.GuildMemberService) *GuildMemberHandler {
	return &GuildMemberHandler{
		guildMemberService: guildMemberService,
	}
}

func (h *GuildMemberHandler) GetAllGuildsByUserID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	
	userIDString := r.Context().Value("userID").(string)
	userID ,err := strconv.Atoi(userIDString)
	if err != nil{
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	guilds, err := h.guildMemberService.GetAllGuildsByUserID(context.Background(), uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(guilds)

}
func (h *GuildMemberHandler) GetAllMembersByGuildID(w http.ResponseWriter, r *http.Request) {

}

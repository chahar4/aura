package services

import (
	"context"
	"time"

	"github.com/chahar4/aura/core/domains"
)

type GuildMemberService struct {
	guildMemberRepo domains.GuildMemberRepository
	timeout         time.Duration
}

func NewGuildMemberService(guildMemberRepo domains.GuildMemberRepository) *GuildMemberService {
	return &GuildMemberService{
		guildMemberRepo: guildMemberRepo,
		timeout:         time.Duration(2) * time.Second,
	}
}

func (s *GuildMemberService) GetAllMemberByGuildID(ctx context.Context, guildID uint) ([]*domains.User, error) {
	ctx, cansel := context.WithTimeout(ctx, s.timeout)
	defer cansel()

	return s.GetAllMemberByGuildID(ctx, guildID)

}

func (s *GuildMemberService) GetAllGuildsByUserID(ctx context.Context, userID uint) ([]*domains.Guild, error) {
	ctx, cansel := context.WithTimeout(ctx, s.timeout)
	defer cansel()

	return s.GetAllGuildsByUserID(ctx, userID)
}

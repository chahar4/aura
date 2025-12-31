package services

import (
	"context"
	"time"

	"github.com/chahar4/aura/core/domains"
)

type GuildService struct {
	guildRepo       domains.GuildRepository
	guildMemberRepo domains.GuildMemberRepository
	timeout         time.Duration
}

func NewGuildService(guildRepo domains.GuildRepository, guildMemberRepository domains.GuildMemberRepository) *GuildService {
	return &GuildService{
		guildRepo:       guildRepo,
		guildMemberRepo: guildMemberRepository,
		timeout:         time.Duration(2) * time.Second,
	}
}

func (s *GuildService) AddGuild(ctx context.Context, userID uint, name string, profile *string) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	guild := domains.Guild{
		Name:    name,
		Profile: profile,
	}
    err := s.guildRepo.AddGuild(ctx, guild ,userID)
	if err != nil {
		return err
	}
	return nil
}

func (s *GuildService) RemoveGuild(ctx context.Context, id uint) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	if err := s.guildRepo.DeleteGulid(ctx, id); err != nil {
		return err
	}
	return nil
}

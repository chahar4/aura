package services

import (
	"context"

	"github.com/chahar4/aura/core/domains"
)

type GuildService struct {
	gulidRepo domains.GuildRepository
}

func NewGuildService(gulidRepo domains.GuildRepository) *GuildService {
	return &GuildService{
		gulidRepo: gulidRepo,
	}
}

func (s *GuildService) AddGuild(ctx context.Context, name string, profile *string) error {
	guild := domains.Guild{
		Name:    name,
		Profile: profile,
	}

	if err := s.gulidRepo.AddGuild(ctx, guild); err != nil {
		return err
	}
	return nil
}

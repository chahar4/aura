package services

import (
	"context"
	"time"

	"github.com/chahar4/aura/core/domains"
)

type RoleService struct {
	repo    domains.RoleRepository
	timeout time.Duration
}

func NewRoleService(roleRepository domains.RoleRepository) *RoleService {
	return &RoleService{
		repo:    roleRepository,
		timeout: time.Duration(2) * time.Second,
	}
}

func (s *RoleService) AddRole(ctx context.Context, guildID uint, name string) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	role := domains.Role{
		GuildID: guildID,
		Name:    name,
	}

	if err := s.repo.AddRole(ctx, role); err != nil {
		return err
	}
	return nil
}

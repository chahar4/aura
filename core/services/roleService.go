package services

import (
	"context"

	"github.com/chahar4/aura/core/domains"
)

type RoleService struct {
	repo domains.RoleRepository
}

func NewRoleService(roleRepository domains.RoleRepository) *RoleService {
	return &RoleService{
		repo: roleRepository,
	}
}

func (s *RoleService) AddRole(ctx context.Context, guildID uint, name string) error {
	role := domains.Role{
		GuildID: guildID,
		Name:    name,
	}

	if err := s.repo.AddRole(ctx, role); err != nil {
		return err
	}
	return nil
}

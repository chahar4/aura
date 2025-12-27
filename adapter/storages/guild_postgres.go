package storages

import (
	"context"
	"errors"

	"github.com/chahar4/aura/core/domains"
	"github.com/chahar4/aura/core/tools"
	"gorm.io/gorm"
)

type GuildPostgresRepo struct {
	db *gorm.DB
}

func NewGuildPostgreRepo(db *gorm.DB) *GuildPostgresRepo {
	return &GuildPostgresRepo{
		db: db,
	}
}

func (p *GuildPostgresRepo) AddGuild(ctx context.Context, guild domains.Guild) error {
	err := gorm.G[domains.Guild](p.db).Create(ctx, &guild)
	if err != nil {
		return tools.ProblemErrDb
	}
	return nil
}

func (p *GuildPostgresRepo) DeleteGulid(ctx context.Context, id uint) error {
	_, err := gorm.G[domains.Guild](p.db).Where("id = ?", id).Delete(ctx)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return tools.NotFoundErrDb
	}
	if err != nil {
		return tools.ProblemErrDb
	}
	return nil
}

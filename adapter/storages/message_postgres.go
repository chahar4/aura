package storages

import (
	"context"

	"github.com/chahar4/aura/core/domains"
	"github.com/chahar4/aura/core/tools"
	"gorm.io/gorm"
)

type MessagePostgresRepo struct {
	db *gorm.DB
}

func NewMessagePostgresRepo(db *gorm.DB) *MessagePostgresRepo {
	return &MessagePostgresRepo{
		db: db,
	}
}

func (p *MessagePostgresRepo) AddMessage(ctx context.Context, message *domains.Message) error {
	err := gorm.G[domains.Message](p.db).Create(ctx, message)
	if err != nil {
		return tools.ProblemErrDb
	}
	return nil
}

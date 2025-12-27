package storages

import (
	"context"
	"errors"

	"github.com/chahar4/aura/core/domains"
	"github.com/chahar4/aura/core/tools"
	"gorm.io/gorm"
)

type GroupChannelPostgresRepo struct {
	db *gorm.DB
}

func NewGroupChannelPostgresRepo(db *gorm.DB) *GroupChannelPostgresRepo {
	return &GroupChannelPostgresRepo{
		db: db,
	}
}

func (p *GroupChannelPostgresRepo) AddGroupChannel(ctx context.Context, groupChannel domains.GroupChannel) error {
	err := p.db.WithContext(ctx).Create(&groupChannel).Error
	if err != nil {
		return tools.ProblemErrDb
	}
	return nil
}

func (p *GroupChannelPostgresRepo) DeleteGroupChannel(ctx context.Context, id uint) error {
	err := p.db.WithContext(ctx).Where("id = ?", id).Delete(&domains.GroupChannel{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return tools.NotFoundErrDb
	}
	if err != nil {
		return tools.ProblemErrDb
	}
	return nil
}
package storages

import (
	"context"

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

func (p *GroupChannelPostgresRepo) GetGuildIDByChannelID(ctx context.Context, channelID uint) (uint, error) {

	channel, err := gorm.G[domains.Channel](p.db).Where("id = ?", channelID).First(ctx)
	if err != nil {
		return 0, tools.ProblemErrDb
	}
	groupChannel, err := gorm.G[domains.GroupChannel](p.db).Where("id = ?", channel.GroupChannelID).First(ctx)

	if err != nil {
		return 0, tools.ProblemErrDb
	}
	return groupChannel.GuildID, nil
}

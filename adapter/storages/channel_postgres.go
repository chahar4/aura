package storages

import (
	"context"

	"github.com/chahar4/aura/core/domains"
	"github.com/chahar4/aura/core/tools"
	"gorm.io/gorm"
)

type ChannelPostgresRepo struct {
	db *gorm.DB
}

func NewChannelPostgresRepo(db *gorm.DB) *ChannelPostgresRepo {
	return &ChannelPostgresRepo{
		db: db,
	}
}

func (p *ChannelPostgresRepo) AddChannel(ctx context.Context, channel *domains.Channel) error {
	err := gorm.G[domains.Channel](p.db).Create(ctx, channel)
	if err != nil {
		return tools.ProblemErrDb
	}

	return nil
}

func (p *ChannelPostgresRepo) GetAllChannelByGroupChannelID(ctx context.Context, groupChannelID uint) ([]*domains.Channel, error) {
	channels, err := gorm.G[*domains.Channel](p.db).Where("group_channel_id = ?", groupChannelID).Find(ctx)
	if err != nil {
		return nil, tools.ProblemErrDb
	}
	return channels, nil
}

func (p *GroupChannelPostgresRepo) IsUserInChannel(ctx context.Context, channelID, userID uint) bool {

	channel, err := gorm.G[domains.Channel](p.db).Where("id = ?", channelID).First(ctx)
	if err != nil {
		return false
	}
	groupChannel, err := gorm.G[domains.GroupChannel](p.db).Where("id = ?", channel.GroupChannelID).First(ctx)

	if err != nil {
		return false
	}

	guildMember, err := gorm.G[domains.GuildMember](p.db).Where("guild_id = ?", groupChannel.GuildID).Find(ctx)
	if err != nil {
		return false
	}

	for _, gm := range guildMember {
		if gm.UserID == userID {
			return true
		}
	}

	return false
}

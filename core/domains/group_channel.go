package domains

import (
	"context"

	"gorm.io/gorm"
)

type GroupChannel struct {
	gorm.Model
	Name     string
	GuildID  uint
	Channels []*Channel
}

type GroupChannelRepository interface {
	AddGroupChannel(ctx context.Context, groupChannel GroupChannel) error
	DeleteGroupChannel(ctx context.Context, id uint) error
}

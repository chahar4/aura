package domains

import (
	"context"

	"gorm.io/gorm"
)

type Channel struct {
	gorm.Model
	Name           string
	GroupChannelID uint
	Messages       []Message
}

type ChannelRepository interface {
	AddChannel(ctx context.Context, channel *Channel) error
	GetAllChannelByGroupChannelID(ctx context.Context, groupChannelID uint) ([]*Channel, error)
	IsUserInChannel(ctx context.Context, channelID, userID uint) bool
}

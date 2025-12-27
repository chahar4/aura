package domains

import (
	"context"

	"gorm.io/gorm"
)

type Guild struct {
	gorm.Model
	Name          string
	Profile       *string
	Members       []GuildMember `gorm:"foreignKey:GuildID;"`
	GroupChannels []GroupChannel
	Role          []Role `gorm:"foreignKey:GuildID;"`
}

type GuildRepository interface {
	AddGuild(ctx context.Context, guild Guild, userID uint) error 
	DeleteGulid(ctx context.Context, id uint) error
}

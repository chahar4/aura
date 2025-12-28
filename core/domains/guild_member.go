package domains

import (
	"context"
	"time"
)

type GuildMember struct {
	GuildID uint `gorm:"primaryKey;autoIncrement:false"`
	UserID  uint `gorm:"primaryKey;autoIncrement:false"`

	Nickname string
	JoinedAt time.Time

	Roles []Role `gorm:"many2many:member_roles;"`

	User  User  `gorm:"foreignKey:UserID"`
	Guild Guild `gorm:"foreignKey:GuildID"`
}

type GuildMemberRepository interface {
	GetAllMemberByGuildID(ctx context.Context, guildID uint) ([]*User, error)
}

package domains

import "time"

type GuildMember struct {
	GuildID uint `gorm:"primaryKey;autoIncrement:false"`
	UserID  uint `gorm:"primaryKey;autoIncrement:false"`

	Nickname string
	JoinedAt time.Time

	Roles []Role `gorm:"many2many:member_roles;"`

	User  User  `gorm:"foreignKey:UserID"`
	Guild Guild `gorm:"foreignKey:GuildID"`
}

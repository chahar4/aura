package domains

import "gorm.io/gorm"

type Guild struct {
	gorm.Model
	Name          string
	Profile       *string
	Users         []*User `gorm:"many2many:guild_user;"`
	GroupChannels []*GroupChannel
}

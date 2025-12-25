package domains

import (
	"gorm.io/gorm"
)

type GroupChannel struct {
	gorm.Model
	Name     string
	GuildID  uint
	Channels []*Channel
}

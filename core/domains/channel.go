package domains

import "gorm.io/gorm"

type Channel struct {
	gorm.Model
	Name           string
	GroupChannelID uint
	Messages       []*Message
}

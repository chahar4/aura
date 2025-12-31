package domains

import (
	"context"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	ChannelID uint
	UserID    uint
	Content   string
}

type MessageRepository interface {
	AddMessage(ctx context.Context, message *Message) error
}

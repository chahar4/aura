package services

import (
	"context"
	"time"

	"github.com/chahar4/aura/core/domains"
	"github.com/chahar4/aura/core/tools"
)

type MessageService struct {
	messageRepository domains.MessageRepository
	userRepository    domains.UserRepository
	channelRepository domains.ChannelRepository
	timeout           time.Duration
}

func NewMessageService(messageRepository domains.MessageRepository, userRepository domains.UserRepository, channelRepository domains.ChannelRepository) *MessageService {
	return &MessageService{
		messageRepository:      messageRepository,
		userRepository:         userRepository,
		channelRepository: channelRepository,
		timeout:                time.Duration(2) * time.Second,
	}
}

func (s *MessageService) SaveMessage(ctx context.Context, channelID, userID uint, content string) error {
	ctx, cansel := context.WithTimeout(ctx, s.timeout)
	defer cansel()

	if ok := s.channelRepository.IsUserInChannel(ctx, channelID,userID); !ok{
		return tools.ProblemErrDb
	}


}

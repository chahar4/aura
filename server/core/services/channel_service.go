package services

import (
	"context"
	"time"

	"github.com/chahar4/aura/core/domains"
)

type ChannelService struct {
	channelRepo domains.ChannelRepository
	timeout     time.Duration
}

func NewChannelService(channelRepo domains.ChannelRepository) *ChannelService {
	return &ChannelService{
		channelRepo: channelRepo,
		timeout:     time.Duration(2) * time.Second,
	}
}

func (s *ChannelService) AddChannel(ctx context.Context, name string, groupChannelID uint) error {
	ctx, cansel := context.WithTimeout(ctx, s.timeout)
	defer cansel()

	channel := domains.Channel{
		Name:           name,
		GroupChannelID: groupChannelID,
	}

	return s.channelRepo.AddChannel(ctx, &channel)
}

func (s *ChannelService) GetAllChannelsByGroupChannelID(ctx context.Context, groupChannelID uint) ([]*domains.Channel, error) {
	ctx, cansel := context.WithTimeout(ctx, s.timeout)
	defer cansel()

	return s.channelRepo.GetAllChannelByGroupChannelID(ctx, groupChannelID)
}

func (s *ChannelService) IsUserInChannel(ctx context.Context, channelID, userID uint) bool {
	ctx, cansel := context.WithTimeout(ctx, s.timeout)
	defer cansel()
	return s.channelRepo.IsUserInChannel(ctx, channelID, userID)
}

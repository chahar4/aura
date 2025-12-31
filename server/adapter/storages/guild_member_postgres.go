package storages

import (
	"context"

	"github.com/chahar4/aura/core/domains"
	"github.com/chahar4/aura/core/tools"
	"gorm.io/gorm"
)

type GuildMemberPostgresRepo struct {
	db *gorm.DB
}

func NewGuildMemberPostgresRepo(db *gorm.DB) *GuildMemberPostgresRepo {
	return &GuildMemberPostgresRepo{
		db: db,
	}
}

func (p *GuildMemberPostgresRepo) GetAllMemberByGuildID(ctx context.Context, guildID uint) ([]*domains.User, error) {
	guildMember, err := gorm.G[domains.GuildMember](p.db).Preload("users", nil).Where("id = ?", guildID).Find(ctx)
	if err != nil {
		return nil, tools.ProblemErrDb
	}
	var users []*domains.User
	for _, u := range guildMember {
		users = append(users, &u.User)
	}
	return users, nil
}

func (p *GuildMemberPostgresRepo) GetAllGuildsByUserID(ctx context.Context, userID uint) ([]*domains.Guild, error) {
	guildMember, err := gorm.G[domains.GuildMember](p.db).Preload("Guilds", nil).Where("id = ?", userID).Find(ctx)
	if err != nil {
		return nil, tools.ProblemErrDb
	}
	var guilds []*domains.Guild
	for _, u := range guildMember {
		guilds = append(guilds, &u.Guild)
	}
	return guilds, nil
}

package storages

import (
	"context"

	"github.com/chahar4/aura/core/domains"
	"github.com/chahar4/aura/core/tools"
	"gorm.io/gorm"
)



type GuildMemberRepo struct{
	db *gorm.DB
}

func NewGuildMemberRepo(db *gorm.DB) *GuildMemberRepo{
	return &GuildMemberRepo{
		db: db,
	}
}

func(p *GuildMemberRepo) AddGuildMember(ctx context.Context , guildMember domains.GuildMember) error{

	err := gorm.G[domains.GuildMember](p.db).Create(ctx, &guildMember)
	if err != nil{
		return tools.ProblemErrDb
	}
	return nil

}

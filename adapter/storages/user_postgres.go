package storages

import (
	"context"
	"errors"
	"time"

	"github.com/chahar4/aura/core/domains"
	"github.com/chahar4/aura/core/tools"
	"gorm.io/gorm"
)

type UserPostgresRepo struct {
	db *gorm.DB
}

func NewUserPostgresRepo(db *gorm.DB) *UserPostgresRepo {
	return &UserPostgresRepo{
		db: db,
	}
}

func (p *UserPostgresRepo) AddUser(ctx context.Context, user domains.User) error {
	err := gorm.G[domains.User](p.db).Create(ctx, &user)
	if err != nil {
		return tools.ProblemErrDb
	}
	return nil
}

func (p *UserPostgresRepo) ChangeForgotCodeUser(ctx context.Context, id uint, code string, expireAt time.Time) error {
	_, err := gorm.G[domains.User](p.db).Where("id = ?", id).Updates(ctx, domains.User{ForgotToken: &code, ExpireForgotToken: &expireAt})
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return tools.NotFoundErrDb
	}
	if err != nil {
		return tools.ProblemErrDb
	}
	return nil
}

func (p *UserPostgresRepo) ChangePasswordUser(ctx context.Context, id uint, password string) error {
	_, err := gorm.G[domains.User](p.db).Where("id = ?", id).Update(ctx, "password_hash", password)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return tools.NotFoundErrDb
	}
	if err != nil {
		return tools.ProblemErrDb
	}
	return nil
}

func (p *UserPostgresRepo) GetUserByID(ctx context.Context, ID uint) (*domains.User, error) {
	user, err := gorm.G[domains.User](p.db).Where("id = ?", ID).First(ctx)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, tools.NotFoundErrDb
	}
	if err != nil {
		return nil, tools.ProblemErrDb
	}
	return &user, nil
}

func (p *UserPostgresRepo) GetUserByEmail(ctx context.Context, email string) (*domains.User, error) {
	user, err := gorm.G[domains.User](p.db).Preload("guild_member", nil).Where("email = ?", email).First(ctx)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, tools.NotFoundErrDb
	}
	if err != nil {
		return nil, tools.ProblemErrDb
	}
	return &user, nil
}

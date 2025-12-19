package domains

import (
	"context"
	"time"

	"gorm.io/gorm"
)

// user model
type User struct {
	gorm.Model
	Username          string
	PasswordHash      string
	Email             string
	Status            string
	OnlineStatus      OnlineStatus
	Role              Role
	ForgotToken       *string
	ExpireForgotToken *time.Time
}

type OnlineStatus int

const (
	StateIdle OnlineStatus = iota
	StateOnline
	StateOffline
	StateDoNotDisturb
)

type UserRepository interface {
	AddUser(ctx context.Context, user User) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	ChangeForgotCodeUser(ctx context.Context, id uint, code string, expireAt time.Time) error
	ChangePasswordUser(ctx context.Context, id uint, password string) error
}

package domains

import (
	"context"

	"gorm.io/gorm"
)

// role model
type Role struct {
	gorm.Model
	GuildID uint
	Name    string
}

type RoleRepository interface {
	AddRole(ctx context.Context, role Role) error
}

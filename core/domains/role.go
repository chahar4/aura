package domains


import "gorm.io/gorm"


//role model
type Role struct {
	gorm.Model
	UserID uint
	Name   string
}

type RoleRepository interface{
	AddRole(role Role) error
}

type RoleService struct {
	repo RoleRepository
}

func NewRoleService(roleRepository RoleRepository) *RoleService{
	return &RoleService{
		repo: roleRepository,
	}
}

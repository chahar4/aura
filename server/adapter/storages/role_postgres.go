package storages

import (
	"github.com/chahar4/aura/core/domains"
	"gorm.io/gorm"
)

type RolePostgresRepo struct{
	db *gorm.DB
}



func (p *RolePostgresRepo) AddRole(role domains.Role) error{
	result :=p.db.Create(role)
	if result.Error != nil || result.RowsAffected == 0{
		return  result.Error
	}
	return nil
}

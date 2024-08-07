package po

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID       int    `gorm:"column:uuid; type:int; not null; index:idx_uuid; primaryKey; autoIncrement; comment:primary Key is ID"`
	RoleName string `gorm:"column:role_name"`
	RoleNote string `gorm:"column:role_note; type:text;"`
}

func (u *Role) TableName() string {
	return "go_db_role"
}

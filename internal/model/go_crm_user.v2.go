// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameGoCrmUserV2 = "go_crm_user_v2"

// GoCrmUser mapped from table <go_crm_user>
type GoCrmUserV2 struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	UUID      int32          `gorm:"column:uuid;not null" json:"uuid"`
	UserName  string         `gorm:"column:user_name" json:"user_name"`
	IsActive  bool           `gorm:"column:is_active" json:"is_active"`
}

// TableName GoCrmUser's table name
func (*GoCrmUserV2) TableName() string {
	return TableNameGoCrmUserV2
}

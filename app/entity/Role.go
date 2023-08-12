package entity

import "gorm.io/gorm"

type Role struct {
	ID        int64          `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`
	Name      string         `gorm:"column:name;type:varchar(100);NOT NULL" json:"name"`
	Hotile    string         `gorm:"column:hotline;type:varchar(100);NOT NULL" json:"hotline"`
	IsService int            `gorm:"column:is_service;type:tinyint(1)" json:"is_service"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (m *Role) TableName() string {
	return "ms_roles"
}

package entity

import "gorm.io/gorm"

type Status struct {
	ID        int            `gorm:"column:id;type:int(11);primary_key" json:"id"`
	Name      string         `gorm:"column:name;type:varchar(255)" json:"name"`
	Color     string         `gorm:"column:color;type:varchar(255)" json:"color"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (m *Status) TableName() string {
	return "ms_status"
}

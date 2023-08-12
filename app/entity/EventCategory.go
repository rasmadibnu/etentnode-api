package entity

import "gorm.io/gorm"

type EventCategory struct {
	ID        int                  `gorm:"column:id;type:int(11);primary_key" json:"id"`
	Name      string               `gorm:"column:name;type:varchar(255)" json:"name"`
	Types     []EventCategoryType  `json:"types"`
	Fields    []EventCategoryField `json:"fields"`
	Roles     []EventCategoryRole  `json:"roles"`
	DeletedAt gorm.DeletedAt       `json:"-"`
}

func (m *EventCategory) TableName() string {
	return "ms_event_categories"
}

type EventCategoryType struct {
	ID              int            `gorm:"column:id;type:int(11);primary_key" json:"id"`
	EventCategoryID int            `gorm:"column:event_category_id;type:int(11)" json:"event_category_id"`
	Name            string         `gorm:"column:name;type:varchar(255)" json:"name"`
	DeletedAt       gorm.DeletedAt `json:"-"`
}

func (m *EventCategoryType) TableName() string {
	return "ms_event_category_types"
}

type EventCategoryField struct {
	ID              int            `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	EventCategoryID int            `gorm:"column:event_category_id;type:int(11)" json:"event_category_id"`
	Name            string         `gorm:"column:name;type:varchar(255)" json:"name"`
	DeletedAt       gorm.DeletedAt `json:"-"`
}

func (m *EventCategoryField) TableName() string {
	return "ms_event_category_fields"
}

type EventCategoryRole struct {
	ID              int  `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	EventCategoryID int  `gorm:"column:event_category_id;type:int(11)" json:"event_category_id"`
	RoleID          int  `gorm:"column:role_id;type:int(11)" json:"role_id"`
	Role            Role `json:"role" `
}

func (m *EventCategoryRole) TableName() string {
	return "tr_event_category_role"
}

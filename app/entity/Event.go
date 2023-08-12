package entity

import (
	"os"
	"time"

	"gorm.io/gorm"
)

type Event struct {
	ID                  int                 `gorm:"column:id;type:int(11);primary_key" json:"id"`
	Description         string              `gorm:"column:description;type:text" json:"description"`
	TypeDescription     string              `gorm:"column:type_description;type:text" json:"type_description"`
	OtherDescription    string              `gorm:"column:other_description;type:text" json:"other_description"`
	TypeVehicleInvolved string              `gorm:"column:type_vehicle_involved;type:text" json:"type_vehicle_involved"`
	VictimInvolved      string              `gorm:"column:victim_involved;type:text" json:"victim_involved"`
	Responsible         string              `gorm:"column:responsible;type:text" json:"responsible"`
	Image               string              `gorm:"column:image;type:text" json:"image"`
	Lat                 string              `gorm:"column:lat;type:varchar(255)" json:"lat"`
	Lng                 string              `gorm:"column:lng;type:varchar(255)" json:"lng"`
	Location            string              `gorm:"column:location;type:text" json:"location"`
	EventCategoryID     int                 `gorm:"column:event_category_id;type:int(11)" json:"event_category_id"`
	EventCategory       EventCategory       `gorm:"foreignKey:EventCategoryID" json:"event_category"`
	EventCategoryTypeID int                 `gorm:"column:event_category_type_id;type:int(11)" json:"event_category_type_id"`
	EventCategoryType   EventCategoryType   `gorm:"foreignKey:EventCategoryTypeID" json:"event_category_type"`
	EventHandling       []EventHandling     `json:"event_handling"`
	EventUserHandling   []EventUserHandling `json:"event_user_handling"`
	StatusID            int                 `gorm:"column:status_id;type:int(11)" json:"status_id"`
	Status              Status              `gorm:"foreignKey:StatusID" json:"status"`
	AssginedBy          int                 `gorm:"column:assigned_by;type:int(11)" json:"assigned_by"`
	UserAssign          User                `gorm:"foreignKey:AssginedBy" json:"user_assign"`
	AssignedAt          time.Time           `gorm:"column:assigned_at;type:timestamp" json:"assigned_at"`
	CreatedAt           time.Time           `gorm:"column:created_at;type:timestamp" json:"created_at"`
	CreatedBy           int                 `gorm:"column:created_by;type:int(11)" json:"created_by"`
	UserCreate          User                `gorm:"foreignKey:CreatedBy" json:"user_create"`
	UpdatedAt           time.Time           `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
	UpdatedBy           int                 `gorm:"column:updated_by;type:int(11)" json:"updated_by"`
	UserUpdate          User                `gorm:"foreignKey:UpdatedBy" json:"user_update"`
	DeletedAt           gorm.DeletedAt      `json:"-"`
}

func (m *Event) TableName() string {
	return "tr_events"
}

func (m *Event) AfterFind(tx *gorm.DB) (err error) {
	m.Image = os.Getenv("APP_HOST") + "/uploads/" + m.Image
	return
}

type EventFields struct {
	ID                       int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	EventID                  int    `gorm:"column:event_id;type:int(11)" json:"event_id"`
	EventCategoryTypeFieldID int    `gorm:"column:event_category_type_field_id;type:int(11)" json:"event_category_type_field_id"`
	Value                    string `gorm:"column:value;type:text" json:"value"`
}

func (m *EventFields) TableName() string {
	return "tr_event_fields"
}

type EventUserHandling struct {
	ID      int  `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	EventID int  `gorm:"column:event_id;type:int(11)" json:"event_id"`
	UserID  int  `gorm:"column:user_id;type:int(11)" json:"user_id"`
	User    User `gorm:"foreignKey:UserID" json:"user"`
}

func (m *EventUserHandling) TableName() string {
	return "tr_event_user_handling"
}

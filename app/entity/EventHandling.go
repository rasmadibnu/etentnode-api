package entity

import (
	"os"
	"time"

	"gorm.io/gorm"
)

type EventHandling struct {
	ID                  int               `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	EventID             int               `gorm:"column:event_id;type:int(11)" json:"event_id"`
	Event               *Event            `gorm:"foreignKey:EventID" json:"event,omitempty"`
	Image               string            `gorm:"column:image;type:text" json:"image"`
	Description         string            `gorm:"column:description;type:text" json:"description"`
	Identification      string            `gorm:"column:identification;type:text" json:"identification"`
	MinorInjuries       int               `gorm:"column:minor_injuries;type:int(11)" json:"minor_injuries"`
	SeriouslyInjuries   int               `gorm:"column:seriously_injuries;type:int(11)" json:"seriously_injuries"`
	Die                 int               `gorm:"column:die;type:int(11)" json:"die"`
	VictimInvolved      string            `gorm:"column:victim_involved;type:text" json:"victim_involved"`
	Location            string            `gorm:"column:location;type:text" json:"location"`
	Lat                 string            `gorm:"column:lat;type:varchar(255)" json:"lat"`
	Lng                 string            `gorm:"column:lng;type:varchar(255)" json:"lng"`
	EventCategoryTypeID int               `gorm:"column:event_category_type_id;type:int(11)" json:"event_category_type_id"`
	EventCategoryType   EventCategoryType `gorm:"foreignKey:EventCategoryTypeID" json:"event_category_type"`
	CreatedAt           time.Time         `gorm:"column:created_at;type:timestamp" json:"created_at"`
	CreatedBy           int               `gorm:"column:created_by;type:int(11)" json:"created_by"`
	UserCreate          User              `gorm:"foreignKey:CreatedBy" json:"user_create"`
}

func (m *EventHandling) TableName() string {
	return "tr_event_handling"
}

func (m *EventHandling) AfterFind(tx *gorm.DB) (err error) {
	m.Image = os.Getenv("APP_HOST") + "/uploads/" + m.Image
	return
}

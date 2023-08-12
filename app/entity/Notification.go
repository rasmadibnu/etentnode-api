package entity

import (
	"time"
)

type Notification struct {
	ID         int       `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Title      string    `gorm:"column:title;type:varchar(255)" json:"title"`
	Message    string    `gorm:"column:message;type:varchar(255)" json:"message"`
	EventID    int       `gorm:"column:event_id;type:int(11)" json:"event_id"`
	Type       int       `gorm:"column:type;type:int(11)" json:"type"`
	ExternalID string    `gorm:"column:external_id;type:text" json:"external_id"`
	Segment    string    `gorm:"column:segment;type:text" json:"segment"`
	SentBy     int       `gorm:"column:sent_by;type:int(11)" json:"sent_by"`
	IsRead     int       `gorm:"column:is_read;type:int(11)" json:"is_read"`
	CreatedAt  time.Time `gorm:"column:created_at;type:timestamp" json:"created_at"`
}

func (m *Notification) TableName() string {
	return "tr_notification"
}

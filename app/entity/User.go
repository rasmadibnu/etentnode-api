package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          int64          `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`
	NIK         string         `gorm:"column:nik;type:varchar(50)" json:"nik"`
	Username    string         `gorm:"column:username;type:varchar(100);NOT NULL" json:"username"`
	Name        string         `gorm:"column:name;type:varchar(100);NOT NULL" json:"name"`
	PhoneNumber string         `gorm:"column:phone_number;type:varchar(100);NOT NULL" json:"phone_number"`
	Email       string         `gorm:"column:email;type:varchar(255);NOT NULL" json:"email"`
	Password    string         `gorm:"column:password;type:varchar(100);NOT NULL" json:"-"`
	RoleID      int            `gorm:"column:role_id;type:int(11);NOT NULL" json:"role_id"`
	Lat         string         `gorm:"column:lat;type:varchar(255)" json:"lat"`
	Lng         string         `gorm:"column:lng;type:varchar(255)" json:"lng"`
	Location    string         `gorm:"column:location;type:text" json:"location"`
	IsActive    *int           `gorm:"column:is_active;type:tinyint(1);NOT NULL" json:"is_active"`
	Role        Role           `gorm:"foreignKey:RoleID" json:"role"`
	CreatedAt   time.Time      `gorm:"column:created_at;type:timestamp" json:"created_at"`
	DeletedAt   gorm.DeletedAt `json:"-"`
}

func (m *User) TableName() string {
	return "ms_users"
}

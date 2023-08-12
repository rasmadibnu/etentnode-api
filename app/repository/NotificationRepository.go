package repository

import (
	"etentnode-api/app/entity"
	"etentnode-api/config"
)

type NotificitaionRepository struct {
	config config.Database
}

func NewNotificitaionRepository(database config.Database) NotificitaionRepository {
	return NotificitaionRepository{
		config: database,
	}
}

// @Summary : Insert Notification
// @Description : Insert Notification to database
// @Author : rasmadibbnu
func (r *NotificitaionRepository) Insert(Notification []entity.Notification) ([]entity.Notification, error) {
	err := r.config.DB.Create(&Notification).Error

	if err != nil {
		return Notification, err
	}

	return Notification, nil
}

// @Summary : Get Notification and update read
// @Description : -
// @Author : rasmadibbnu
func (r *NotificitaionRepository) FindAll(param map[string]interface{}, id int, role string) ([]entity.Notification, error) {
	var Notification []entity.Notification

	err := r.config.DB.Where(param).Where("segment", role).Order("created_at desc").Limit(10).Find(&Notification).Error

	if err != nil {
		return Notification, err
	}

	return Notification, nil
}

// @Summary : Update Notification
// @Description : Update Notification by EventID
// @Author : rasmadibbnu
func (r *NotificitaionRepository) UpdateByEvent(Notification entity.Notification, ID int, segment string) (entity.Notification, error) {
	err := r.config.DB.Debug().Where("event_id = ? and segment = ?", ID, segment).Updates(&Notification).Error

	if err != nil {
		return Notification, err
	}

	return Notification, nil
}

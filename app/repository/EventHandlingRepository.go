package repository

import (
	"etentnode-api/app/entity"
	"etentnode-api/config"
)

type EventHandlingRepository struct {
	config config.Database
}

func NewEventHandlingRepository(database config.Database) EventHandlingRepository {
	return EventHandlingRepository{
		config: database,
	}
}

// @Summary : Insert Event Handling
// @Description : Insert Event Handling to database
// @Author : rasmadibbnu
func (r *EventHandlingRepository) Insert(EventHandling entity.EventHandling) (entity.EventHandling, error) {
	err := r.config.DB.Create(&EventHandling).Error

	if err != nil {
		return EventHandling, err
	}

	return EventHandling, nil
}

// @Summary : Get Event Handling
// @Description : -
// @Author : rasmadibbnu
func (r *EventHandlingRepository) FindAll(param map[string]interface{}) ([]entity.EventHandling, error) {
	var EventHandling []entity.EventHandling

	err := r.config.DB.Where(param).Preload("EventCategoryType").Preload("Event").Find(&EventHandling).Error

	if err != nil {
		return EventHandling, err
	}

	return EventHandling, nil
}

// @Summary : Get Event Handling
// @Description : Find Event Handling by ID
// @Author : rasmadibbnu
func (r *EventHandlingRepository) FindById(ID int) (entity.EventHandling, error) {
	var EventHandling entity.EventHandling

	err := r.config.DB.Where("id = ?", ID).Preload("EventCategoryType").Preload("Event").First(&EventHandling).Error

	if err != nil {
		return EventHandling, err
	}

	return EventHandling, nil
}

// @Summary : Update Event Handling
// @Description : Update Event Handling by ID
// @Author : rasmadibbnu
func (r *EventHandlingRepository) Update(EventHandling entity.EventHandling, ID int) (entity.EventHandling, error) {
	err := r.config.DB.Where("id = ?", ID).Updates(&EventHandling).Error

	if err != nil {
		return EventHandling, err
	}

	return EventHandling, nil
}

// @Summary : Delete Event Handling
// @Description : Delete Event Handling temporary
// @Author : rasmadibbnu
func (r *EventHandlingRepository) Delete(ID int) (bool, error) {
	var EventHandling entity.EventHandling

	err := r.config.DB.Where("id = ?", ID).Delete(&EventHandling).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

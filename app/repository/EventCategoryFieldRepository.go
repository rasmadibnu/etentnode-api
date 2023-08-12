package repository

import (
	"etentnode-api/app/entity"
	"etentnode-api/config"
)

type EventCategoryFieldRepository struct {
	config config.Database
}

func NewEventCategoryFieldRepository(database config.Database) EventCategoryFieldRepository {
	return EventCategoryFieldRepository{
		config: database,
	}
}

// @Summary : Insert Event Category Field
// @Description : Insert Event Category Field to database
// @Author : rasmadibbnu
func (r *EventCategoryFieldRepository) Insert(EventCategoryField entity.EventCategoryField) (entity.EventCategoryField, error) {
	err := r.config.DB.Create(&EventCategoryField).Error

	if err != nil {
		return EventCategoryField, err
	}

	return EventCategoryField, nil
}

// @Summary : Get Event Category Field
// @Description : -
// @Author : rasmadibbnu
func (r *EventCategoryFieldRepository) FindAll(param map[string]interface{}) ([]entity.EventCategoryField, error) {
	var EventCategoryField []entity.EventCategoryField

	err := r.config.DB.Where(param).Find(&EventCategoryField).Error

	if err != nil {
		return EventCategoryField, err
	}

	return EventCategoryField, nil
}

// @Summary : Get EventCategoryField
// @Description : Find EventCategoryField by ID
// @Author : rasmadibbnu
func (r *EventCategoryFieldRepository) FindById(ID int) (entity.EventCategoryField, error) {
	var EventCategoryField entity.EventCategoryField

	err := r.config.DB.First(&EventCategoryField).Error

	if err != nil {
		return EventCategoryField, err
	}

	return EventCategoryField, nil
}

// @Summary : Update Event Category Field
// @Description : Update Event Category Field by ID
// @Author : rasmadibbnu
func (r *EventCategoryFieldRepository) Update(EventCategoryField entity.EventCategoryField, ID int) (entity.EventCategoryField, error) {
	err := r.config.DB.Where("id = ?", ID).Updates(&EventCategoryField).Error

	if err != nil {
		return EventCategoryField, err
	}

	return EventCategoryField, nil
}

// @Summary : Delete Event Category Field
// @Description : Delete Event Category Field temporary
// @Author : rasmadibbnu
func (r *EventCategoryFieldRepository) Delete(ID int) (bool, error) {
	var EventCategoryField entity.EventCategoryField

	err := r.config.DB.Where("id = ?", ID).Delete(&EventCategoryField).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

package repository

import (
	"etentnode-api/app/entity"
	"etentnode-api/config"
)

type EventCategoryTypeRepository struct {
	config config.Database
}

func NewEventCategoryTypeRepository(database config.Database) EventCategoryTypeRepository {
	return EventCategoryTypeRepository{
		config: database,
	}
}

// @Summary : Insert Event Category Type
// @Description : Insert Event Category Type to database
// @Author : rasmadibbnu
func (r *EventCategoryTypeRepository) Insert(EventCategoryType entity.EventCategoryType) (entity.EventCategoryType, error) {
	err := r.config.DB.Create(&EventCategoryType).Error

	if err != nil {
		return EventCategoryType, err
	}

	return EventCategoryType, nil
}

// @Summary : Get Event Category Type
// @Description : -
// @Author : rasmadibbnu
func (r *EventCategoryTypeRepository) FindAll(param map[string]interface{}) ([]entity.EventCategoryType, error) {
	var EventCategoryType []entity.EventCategoryType

	err := r.config.DB.Where(param).Find(&EventCategoryType).Error

	if err != nil {
		return EventCategoryType, err
	}

	return EventCategoryType, nil
}

// @Summary : Get EventCategoryType
// @Description : Find EventCategoryType by ID
// @Author : rasmadibbnu
func (r *EventCategoryTypeRepository) FindById(ID int) (entity.EventCategoryType, error) {
	var EventCategoryType entity.EventCategoryType

	err := r.config.DB.First(&EventCategoryType).Error

	if err != nil {
		return EventCategoryType, err
	}

	return EventCategoryType, nil
}

// @Summary : Update Event Category Type
// @Description : Update Event Category Type by ID
// @Author : rasmadibbnu
func (r *EventCategoryTypeRepository) Update(EventCategoryType entity.EventCategoryType, ID int) (entity.EventCategoryType, error) {
	err := r.config.DB.Where("id = ?", ID).Updates(&EventCategoryType).Error

	if err != nil {
		return EventCategoryType, err
	}

	return EventCategoryType, nil
}

// @Summary : Delete Event Category Type
// @Description : Delete Event Category Type temporary
// @Author : rasmadibbnu
func (r *EventCategoryTypeRepository) Delete(ID int) (bool, error) {
	var EventCategoryType entity.EventCategoryType

	err := r.config.DB.Where("id = ?", ID).Delete(&EventCategoryType).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

package repository

import (
	"etentnode-api/app/entity"
	"etentnode-api/config"
)

type EventCategoryRepository struct {
	config config.Database
}

func NewEventCategoryRepository(database config.Database) EventCategoryRepository {
	return EventCategoryRepository{
		config: database,
	}
}

// @Summary : Insert Event Category
// @Description : Insert Event Category to database
// @Author : rasmadibbnu
func (r *EventCategoryRepository) Insert(EventCategory entity.EventCategory) (entity.EventCategory, error) {
	err := r.config.DB.Create(&EventCategory).Error

	if err != nil {
		return EventCategory, err
	}

	return EventCategory, nil
}

// @Summary : Get Event Category
// @Description : -
// @Author : rasmadibbnu
func (r *EventCategoryRepository) FindAll(param map[string]interface{}) ([]entity.EventCategory, error) {
	var EventCategory []entity.EventCategory

	err := r.config.DB.Where(param).Preload("Fields").Preload("Roles.Role").Preload("Types").Find(&EventCategory).Error

	if err != nil {
		return EventCategory, err
	}

	return EventCategory, nil
}

// @Summary : Get Event Category
// @Description : Find Event Category by ID
// @Author : rasmadibbnu
func (r *EventCategoryRepository) FindById(ID int) (entity.EventCategory, error) {
	var EventCategory entity.EventCategory

	err := r.config.DB.Where("id = ?", ID).Preload("Fields").Preload("Roles.Role").Preload("Types").First(&EventCategory).Error

	if err != nil {
		return EventCategory, err
	}

	return EventCategory, nil
}

// @Summary : Update Event Category
// @Description : Update Event Category by ID
// @Author : rasmadibbnu
func (r *EventCategoryRepository) Update(EventCategory entity.EventCategory, ID int) (entity.EventCategory, error) {
	err := r.config.DB.Where("id = ?", ID).Updates(&EventCategory).Error

	if err != nil {
		return EventCategory, err
	}

	return EventCategory, nil
}

// @Summary : Delete Event Category
// @Description : Delete Event Category temporary
// @Author : rasmadibbnu
func (r *EventCategoryRepository) Delete(ID int) (bool, error) {
	var EventCategory entity.EventCategory

	err := r.config.DB.Where("id = ?", ID).Delete(&EventCategory).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

// @Summary : Assign Event Category Role
// @Description : Assign Event Category Role to database
// @Author : rasmadibbnu
func (r *EventCategoryRepository) AssignRole(EventCategory []entity.EventCategoryRole) ([]entity.EventCategoryRole, error) {

	_, err := r.DeleteRoleByCategory(EventCategory[0].EventCategoryID)

	if err != nil {
		return EventCategory, err
	}

	err = r.config.DB.Create(&EventCategory).Error

	if err != nil {
		return EventCategory, err
	}

	return EventCategory, nil
}

// @Summary : Delete Role by Event Category
// @Description : Delete Role by Event Category temporary
// @Author : rasmadibbnu
func (r *EventCategoryRepository) DeleteRoleByCategory(ID int) (bool, error) {
	var EventCategoryRole entity.EventCategoryRole

	err := r.config.DB.Where("event_category_id = ?", ID).Delete(&EventCategoryRole).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

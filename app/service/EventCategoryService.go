package service

import (
	"etentnode-api/app/entity"
	"etentnode-api/app/repository"
)

type EventCategoryService struct {
	repository repository.EventCategoryRepository
}

func NewEventCategoryService(r repository.EventCategoryRepository) EventCategoryService {
	return EventCategoryService{
		repository: r,
	}
}

// @Summary : List Event Category
// @Description : Get Event Category from repository
// @Author : rasmadibnu
func (s *EventCategoryService) List(param map[string]interface{}) ([]entity.EventCategory, error) {
	EventCategory, err := s.repository.FindAll(param)

	if err != nil {
		return EventCategory, err
	}

	return EventCategory, nil
}

// @Summary : Insert Event Category
// @Description : insert Event Category to repository
// @Author : rasmadibnu
func (s *EventCategoryService) Insert(EventCategory entity.EventCategory) (entity.EventCategory, error) {
	EventCategory, err := s.repository.Insert(EventCategory)

	if err != nil {
		return EventCategory, err
	}

	return EventCategory, nil
}

// @Summary : Find Event Category
// @Description : Find Event Category by id from repository
// @Author : rasmadibnu
func (s *EventCategoryService) FindById(ID int) (entity.EventCategory, error) {
	EventCategory, err := s.repository.FindById(ID)

	if err != nil {
		return EventCategory, err
	}

	return EventCategory, nil
}

// @Summary : Update Event Category
// @Description : Update Event Category by id from repository
// @Author : rasmadibnu
func (s *EventCategoryService) Update(EventCategory entity.EventCategory, ID int) (entity.EventCategory, error) {

	EventCategory, err := s.repository.Update(EventCategory, ID)

	if err != nil {
		return EventCategory, err
	}

	return EventCategory, nil
}

// @Summary : Delete Event Category
// @Description : Delete Event Category from repository
// @Author : rasmadibnu
func (s *EventCategoryService) Delete(ID int) (bool, error) {
	EventCategory, err := s.repository.Delete(ID)

	if err != nil {
		return false, err
	}

	return EventCategory, nil
}

// @Summary : Assign Role to Event Category
// @Description : Assign Role to Event Category
// @Author : rasmadibnu
func (s *EventCategoryService) AssignRole(EventCategoryRole []entity.EventCategoryRole) ([]entity.EventCategoryRole, error) {
	EventCategoryRole, err := s.repository.AssignRole(EventCategoryRole)

	if err != nil {
		return EventCategoryRole, err
	}

	return EventCategoryRole, nil
}

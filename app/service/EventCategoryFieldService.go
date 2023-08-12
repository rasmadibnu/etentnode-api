package service

import (
	"etentnode-api/app/entity"
	"etentnode-api/app/repository"
)

type EventCategoryFieldService struct {
	repository repository.EventCategoryFieldRepository
}

func NewEventCategoryFieldService(r repository.EventCategoryFieldRepository) EventCategoryFieldService {
	return EventCategoryFieldService{
		repository: r,
	}
}

// @Summary : List Event Category Field
// @Description : Get Event Category Field from repository
// @Author : rasmadibnu
func (s *EventCategoryFieldService) List(param map[string]interface{}) ([]entity.EventCategoryField, error) {
	EventCategoryField, err := s.repository.FindAll(param)

	if err != nil {
		return EventCategoryField, err
	}

	return EventCategoryField, nil
}

// @Summary : Insert Event Category Field
// @Description : insert Event Category Field to repository
// @Author : rasmadibnu
func (s *EventCategoryFieldService) Insert(EventCategoryField entity.EventCategoryField) (entity.EventCategoryField, error) {
	EventCategoryField, err := s.repository.Insert(EventCategoryField)

	if err != nil {
		return EventCategoryField, err
	}

	return EventCategoryField, nil
}

// @Summary : Find Event Category Field
// @Description : Find Event Category Field by id from repository
// @Author : rasmadibnu
func (s *EventCategoryFieldService) FindById(ID int) (entity.EventCategoryField, error) {
	EventCategoryField, err := s.repository.FindById(ID)

	if err != nil {
		return EventCategoryField, err
	}

	return EventCategoryField, nil
}

// @Summary : Update Event Category Field
// @Description : Update Event Category Field by id from repository
// @Author : rasmadibnu
func (s *EventCategoryFieldService) Update(EventCategoryField entity.EventCategoryField, ID int) (entity.EventCategoryField, error) {

	EventCategoryField, err := s.repository.Update(EventCategoryField, ID)

	if err != nil {
		return EventCategoryField, err
	}

	return EventCategoryField, nil
}

// @Summary : Delete Event Category Field
// @Description : Delete Event Category Field from repository
// @Author : rasmadibnu
func (s *EventCategoryFieldService) Delete(ID int) (bool, error) {
	EventCategoryField, err := s.repository.Delete(ID)

	if err != nil {
		return false, err
	}

	return EventCategoryField, nil
}

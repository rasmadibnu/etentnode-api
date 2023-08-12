package service

import (
	"etentnode-api/app/entity"
	"etentnode-api/app/repository"
)

type EventCategoryTypeService struct {
	repository repository.EventCategoryTypeRepository
}

func NewEventCategoryTypeService(r repository.EventCategoryTypeRepository) EventCategoryTypeService {
	return EventCategoryTypeService{
		repository: r,
	}
}

// @Summary : List Event Category Type
// @Description : Get Event Category Type from repository
// @Author : rasmadibnu
func (s *EventCategoryTypeService) List(param map[string]interface{}) ([]entity.EventCategoryType, error) {
	EventCategoryType, err := s.repository.FindAll(param)

	if err != nil {
		return EventCategoryType, err
	}

	return EventCategoryType, nil
}

// @Summary : Insert Event Category Type
// @Description : insert Event Category Type to repository
// @Author : rasmadibnu
func (s *EventCategoryTypeService) Insert(EventCategoryType entity.EventCategoryType) (entity.EventCategoryType, error) {
	EventCategoryType, err := s.repository.Insert(EventCategoryType)

	if err != nil {
		return EventCategoryType, err
	}

	return EventCategoryType, nil
}

// @Summary : Find Event Category Type
// @Description : Find Event Category Type by id from repository
// @Author : rasmadibnu
func (s *EventCategoryTypeService) FindById(ID int) (entity.EventCategoryType, error) {
	EventCategoryType, err := s.repository.FindById(ID)

	if err != nil {
		return EventCategoryType, err
	}

	return EventCategoryType, nil
}

// @Summary : Update Event Category Type
// @Description : Update Event Category Type by id from repository
// @Author : rasmadibnu
func (s *EventCategoryTypeService) Update(EventCategoryType entity.EventCategoryType, ID int) (entity.EventCategoryType, error) {

	EventCategoryType, err := s.repository.Update(EventCategoryType, ID)

	if err != nil {
		return EventCategoryType, err
	}

	return EventCategoryType, nil
}

// @Summary : Delete Event Category Type
// @Description : Delete Event Category Type from repository
// @Author : rasmadibnu
func (s *EventCategoryTypeService) Delete(ID int) (bool, error) {
	EventCategoryType, err := s.repository.Delete(ID)

	if err != nil {
		return false, err
	}

	return EventCategoryType, nil
}

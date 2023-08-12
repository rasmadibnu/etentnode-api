package service

import (
	"etentnode-api/app/entity"
	"etentnode-api/app/repository"
)

type EventHandlingService struct {
	repository repository.EventHandlingRepository
}

func NewEventHandlingService(r repository.EventHandlingRepository) EventHandlingService {
	return EventHandlingService{
		repository: r,
	}
}

// @Summary : List Event Handling
// @Description : Get Event Handling from repository
// @Author : rasmadibnu
func (s *EventHandlingService) List(param map[string]interface{}) ([]entity.EventHandling, error) {
	EventHandling, err := s.repository.FindAll(param)

	if err != nil {
		return EventHandling, err
	}

	return EventHandling, nil
}

// @Summary : Insert Event Handling
// @Description : insert Event Handling to repository
// @Author : rasmadibnu
func (s *EventHandlingService) Insert(EventHandling entity.EventHandling) (entity.EventHandling, error) {
	EventHandling, err := s.repository.Insert(EventHandling)

	if err != nil {
		return EventHandling, err
	}

	return EventHandling, nil
}

// @Summary : Find Event Handling
// @Description : Find Event Handling by id from repository
// @Author : rasmadibnu
func (s *EventHandlingService) FindById(ID int) (entity.EventHandling, error) {
	EventHandling, err := s.repository.FindById(ID)

	if err != nil {
		return EventHandling, err
	}

	return EventHandling, nil
}

// @Summary : Update Event Handling
// @Description : Update Event Handling by id from repository
// @Author : rasmadibnu
func (s *EventHandlingService) Update(EventHandling entity.EventHandling, ID int) (entity.EventHandling, error) {

	EventHandling, err := s.repository.Update(EventHandling, ID)

	if err != nil {
		return EventHandling, err
	}

	return EventHandling, nil
}

// @Summary : Delete Event Handling
// @Description : Delete Event Handling from repository
// @Author : rasmadibnu
func (s *EventHandlingService) Delete(ID int) (bool, error) {
	EventHandling, err := s.repository.Delete(ID)

	if err != nil {
		return false, err
	}

	return EventHandling, nil
}

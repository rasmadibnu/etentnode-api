package service

import (
	"etentnode-api/app/entity"
	"etentnode-api/app/repository"
	"time"
)

type EventService struct {
	repository repository.EventRepository
}

func NewEventService(r repository.EventRepository) EventService {
	return EventService{
		repository: r,
	}
}

// @Summary : List Event by user assign
// @Description : Get Event from repository
// @Author : rasmadibnu
func (s *EventService) ListByUserAssign(param map[string]interface{}, id string) ([]entity.Event, error) {
	Event, err := s.repository.ListByUserAssign(param, id)

	if err != nil {
		return Event, err
	}

	return Event, nil
}

// @Summary : Get count
// @Description : Get count
// @Author : rasmadibnu
func (s *EventService) GetCountEvent(param map[string]interface{}, id int) (int64, error) {
	Event, err := s.repository.GetCountEvent(param, id)

	if err != nil {
		return Event, err
	}

	return Event, nil
}

// @Summary : List Event
// @Description : Get Event from repository
// @Author : rasmadibnu
func (s *EventService) List(param map[string]interface{}) ([]entity.Event, error) {
	Event, err := s.repository.FindAll(param)

	if err != nil {
		return Event, err
	}

	return Event, nil
}

// @Summary : Insert Event
// @Description : insert Event to repository
// @Author : rasmadibnu
func (s *EventService) Insert(Event entity.Event) (entity.Event, error) {
	Event, err := s.repository.Insert(Event)

	if err != nil {
		return Event, err
	}

	return Event, nil
}

// @Summary : Find Event
// @Description : Find Event by id from repository
// @Author : rasmadibnu
func (s *EventService) FindById(ID int) (entity.Event, error) {
	Event, err := s.repository.FindById(ID)

	if err != nil {
		return Event, err
	}

	return Event, nil
}

// @Summary : Update Event
// @Description : Update Event by id from repository
// @Author : rasmadibnu
func (s *EventService) Update(Event entity.Event, ID int) (entity.Event, error) {

	Event, err := s.repository.Update(Event, ID)

	if err != nil {
		return Event, err
	}

	return Event, nil
}

// @Summary : Delete Event
// @Description : Delete Event from repository
// @Author : rasmadibnu
func (s *EventService) Delete(ID int) (bool, error) {
	Event, err := s.repository.Delete(ID)

	if err != nil {
		return false, err
	}

	return Event, nil
}

// @Summary : Assign User to Event Category
// @Description : Assign User to Event Category
// @Author : rasmadibnu
func (s *EventService) AssignUser(EventUserHandling []entity.EventUserHandling, user_id int) ([]entity.EventUserHandling, error) {
	EventUserHandling, err := s.repository.AssignUser(EventUserHandling)

	if err != nil {
		return EventUserHandling, err
	}

	_, err = s.repository.Update(entity.Event{StatusID: 2, AssginedBy: user_id, AssignedAt: time.Now()}, EventUserHandling[0].EventID)

	if err != nil {
		return EventUserHandling, err
	}

	return EventUserHandling, nil
}

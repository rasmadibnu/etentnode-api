package service

import (
	"etentnode-api/app/entity"
	"etentnode-api/app/repository"
)

type NotificationService struct {
	repository repository.NotificitaionRepository
}

func NewNotificationService(r repository.NotificitaionRepository) NotificationService {
	return NotificationService{
		repository: r,
	}
}

// @Summary : List Notification
// @Description : Get Notification from repository
// @Author : rasmadibnu
func (s *NotificationService) List(param map[string]interface{}, id int, role string) ([]entity.Notification, error) {
	Notification, err := s.repository.FindAll(param, id, role)

	if err != nil {
		return Notification, err
	}

	return Notification, nil
}

// @Summary : Insert Notification
// @Description : insert Notification to repository
// @Author : rasmadibnu
func (s *NotificationService) Insert(Notification []entity.Notification) ([]entity.Notification, error) {
	Notification, err := s.repository.Insert(Notification)

	if err != nil {
		return Notification, err
	}

	return Notification, nil
}

// @Summary : Update Notification
// @Description : Update Notification by id from repository
// @Author : rasmadibnu
func (s *NotificationService) UpdateByEvent(Notification entity.Notification, ID int, segment string) (entity.Notification, error) {

	Notification, err := s.repository.UpdateByEvent(Notification, ID, segment)

	if err != nil {
		return Notification, err
	}

	return Notification, nil
}

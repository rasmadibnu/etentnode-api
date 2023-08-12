package service

import (
	"etentnode-api/app/entity"
	"etentnode-api/app/repository"
)

type StatusService struct {
	repository repository.StatusRepository
}

func NewStatusService(r repository.StatusRepository) StatusService {
	return StatusService{
		repository: r,
	}
}

// @Summary : List Status
// @Description : Get Status from repository
// @Author : rasmadibnu
func (s *StatusService) List(param map[string]interface{}) ([]entity.Status, error) {
	Status, err := s.repository.FindAll(param)

	if err != nil {
		return Status, err
	}

	return Status, nil
}

// @Summary : Insert Status
// @Description : insert Status to repository
// @Author : rasmadibnu
func (s *StatusService) Insert(Status entity.Status) (entity.Status, error) {
	Status, err := s.repository.Insert(Status)

	if err != nil {
		return Status, err
	}

	return Status, nil
}

// @Summary : Find Status
// @Description : Find Status by id from repository
// @Author : rasmadibnu
func (s *StatusService) FindById(ID int) (entity.Status, error) {
	Status, err := s.repository.FindById(ID)

	if err != nil {
		return Status, err
	}

	return Status, nil
}

// @Summary : Update Status
// @Description : Update Status by id from repository
// @Author : rasmadibnu
func (s *StatusService) Update(Status entity.Status, ID int) (entity.Status, error) {

	Status, err := s.repository.Update(Status, ID)

	if err != nil {
		return Status, err
	}

	return Status, nil
}

// @Summary : Delete Status
// @Description : Delete Status from repository
// @Author : rasmadibnu
func (s *StatusService) Delete(ID int) (bool, error) {
	Status, err := s.repository.Delete(ID)

	if err != nil {
		return false, err
	}

	return Status, nil
}

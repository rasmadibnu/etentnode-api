package service

import (
	"etentnode-api/app/entity"
	"etentnode-api/app/repository"
)

type RoleService struct {
	repository repository.RoleRepository
}

func NewRoleService(r repository.RoleRepository) RoleService {
	return RoleService{
		repository: r,
	}
}

// @Summary : List Role
// @Description : Get Role from repository
// @Author : rasmadibnu
func (s *RoleService) List(param map[string]interface{}) ([]entity.Role, error) {
	Role, err := s.repository.FindAll(param)

	if err != nil {
		return Role, err
	}

	return Role, nil
}

// @Summary : Insert Role
// @Description : insert Role to repository
// @Author : rasmadibnu
func (s *RoleService) Insert(Role entity.Role) (entity.Role, error) {
	Role, err := s.repository.Insert(Role)

	if err != nil {
		return Role, err
	}

	return Role, nil
}

// @Summary : Find Role
// @Description : Find Role by id from repository
// @Author : rasmadibnu
func (s *RoleService) FindById(ID int) (entity.Role, error) {
	Role, err := s.repository.FindById(ID)

	if err != nil {
		return Role, err
	}

	return Role, nil
}

// @Summary : Update Role
// @Description : Update Role by id from repository
// @Author : rasmadibnu
func (s *RoleService) Update(Role entity.Role, ID int) (entity.Role, error) {

	Role, err := s.repository.Update(Role, ID)

	if err != nil {
		return Role, err
	}

	return Role, nil
}

// @Summary : Delete Role
// @Description : Delete Role from repository
// @Author : rasmadibnu
func (s *RoleService) Delete(ID int) (bool, error) {
	Role, err := s.repository.Delete(ID)

	if err != nil {
		return false, err
	}

	return Role, nil
}

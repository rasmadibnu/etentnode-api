package service

import (
	"etentnode-api/app/entity"
	"etentnode-api/app/repository"
	"etentnode-api/security"
)

type UserService struct {
	repository repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return UserService{
		repository: r,
	}
}

// @Summary : List User
// @Description : Get users from repository
// @Author : rasmadibnu
func (s *UserService) List(m map[string]interface{}) ([]entity.User, error) {
	users, err := s.repository.FindAll(m)

	if err != nil {
		return users, err
	}

	return users, nil
}

// @Summary : Insert user
// @Description : Insert user to repository
// @Author : rasmadibnu
func (s *UserService) Insert(req entity.User) (entity.User, error) {
	hash, err := security.HashPassword(req.Password)

	if err != nil {
		return entity.User{}, err
	}

	user := entity.User{
		Username:    req.Username,
		Name:        req.Name,
		Password:    hash,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
		RoleID:      req.RoleID,
		IsActive:    req.IsActive,
	}

	newUser, err := s.repository.Insert(user)

	if err != nil {
		return entity.User{}, err
	}

	return newUser, nil
}

// @Summary : Find user
// @Description : Find user by username
// @Author : rasmadibnu
func (s *UserService) FindByUsername(username string) (entity.User, error) {
	user, err := s.repository.FindByUsername(username)

	if err != nil {
		return user, err
	}

	return user, nil
}

// @Summary : Find user
// @Description : Find user by email
// @Author : rasmadibnu
func (s *UserService) FindByEmail(email string) (entity.User, error) {
	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return user, err
	}

	return user, nil
}

// @Summary : Find user
// @Description : Find user by id
// @Author : rasmadibnu
func (s *UserService) FindById(ID int) (entity.User, error) {
	user, err := s.repository.FindById(ID)

	if err != nil {
		return user, err
	}

	return user, nil
}

// @Summary : Udate user
// @Description : Update user to repository
// @Author : rasmadibnu
func (s *UserService) Update(req entity.User, ID int) (entity.User, error) {
	hash, err := security.HashPassword(req.Password)

	if err != nil {
		return entity.User{}, err
	}

	if req.Password != "" {
		user := entity.User{
			Username:    req.Username,
			Name:        req.Name,
			Password:    hash,
			PhoneNumber: req.PhoneNumber,
			Email:       req.Email,
			RoleID:      req.RoleID,
			IsActive:    req.IsActive,
			Lat:         req.Lat,
			Lng:         req.Lng,
			Location:    req.Location,
		}

		updateUser, err := s.repository.Update(user, ID)
		if err != nil {
			return updateUser, err
		}

		return updateUser, nil
	} else {
		user := entity.User{
			Username:    req.Username,
			Name:        req.Name,
			PhoneNumber: req.PhoneNumber,
			Email:       req.Email,
			RoleID:      req.RoleID,
			IsActive:    req.IsActive,
			Lat:         req.Lat,
			Lng:         req.Lng,
			Location:    req.Location,
		}

		updateUser, err := s.repository.Update(user, ID)
		if err != nil {
			return updateUser, err
		}

		return updateUser, nil
	}
}

// @Summary : Delete user
// @Description : Delete user to repository
// @Author : rasmadibnu
func (s *UserService) Delete(ID int) (bool, error) {
	_, err := s.repository.Delete(ID)
	if err != nil {
		return true, nil
	}

	return true, nil

}

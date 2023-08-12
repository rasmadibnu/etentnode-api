package repository

import (
	"etentnode-api/app/entity"
	"etentnode-api/config"
)

type UserRepository struct {
	config config.Database
}

func NewUserRepository(database config.Database) UserRepository {
	return UserRepository{
		config: database,
	}
}

// @Summary : Insert user
// @Description : Insert user to database
// @Author : rasmadibnu
func (r *UserRepository) Insert(user entity.User) (entity.User, error) {
	err := r.config.DB.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

// @Summary : Get users
// @Description : -
// @Author : rasmadibnu
func (r *UserRepository) FindAll(m map[string]interface{}) ([]entity.User, error) {
	var users []entity.User

	err := r.config.DB.Where(m).Preload("Role").Order("id desc").Find(&users).Error

	if err != nil {
		return users, err
	}

	return users, nil
}

// @Summary : Get user
// @Description : Find user by Username
// @Author : rasmadibnu
func (r *UserRepository) FindByUsername(username string) (entity.User, error) {
	var user entity.User

	err := r.config.DB.Where("username = ?", username).Preload("Role").First(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

// @Summary : Get user
// @Description : Find user by Email
// @Author : rasmadibnu
func (r *UserRepository) FindByEmail(email string) (entity.User, error) {
	var user entity.User

	err := r.config.DB.Where("email = ?", email).Preload("Role").First(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

// @Summary : Get user
// @Description : Find user by ID
// @Author : rasmadibnu
func (r *UserRepository) FindById(ID int) (entity.User, error) {
	var user entity.User
	err := r.config.DB.Preload("Role").Where("id = ?", ID).First(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

// @Summary : Update user
// @Description : Update user
// @Author : rasmadibnu
func (r *UserRepository) Update(user entity.User, ID int) (entity.User, error) {
	err := r.config.DB.Where("id = ?", ID).Updates(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

// @Summary : Delete user
// @Description : Delete user temporary
// @Author : rasmadibnu
func (r *UserRepository) Delete(ID int) (bool, error) {
	var user entity.User

	err := r.config.DB.Where("id = ?", ID).Delete(&user).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

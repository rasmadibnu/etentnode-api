package repository

import (
	"etentnode-api/app/entity"
	"etentnode-api/config"
)

type RoleRepository struct {
	config config.Database
}

func NewRoleRepository(database config.Database) RoleRepository {
	return RoleRepository{
		config: database,
	}
}

// @Summary : Insert Role
// @Description : Insert Role to database
// @Author : rasmadibbnu
func (r *RoleRepository) Insert(Role entity.Role) (entity.Role, error) {
	err := r.config.DB.Create(&Role).Error

	if err != nil {
		return Role, err
	}

	return Role, nil
}

// @Summary : Get Roles
// @Description : -
// @Author : rasmadibbnu
func (r *RoleRepository) FindAll(param map[string]interface{}) ([]entity.Role, error) {
	var Roles []entity.Role

	err := r.config.DB.Where(param).Find(&Roles).Error

	if err != nil {
		return Roles, err
	}

	return Roles, nil
}

// @Summary : Get Role
// @Description : Find Role by ID
// @Author : rasmadibbnu
func (r *RoleRepository) FindById(ID int) (entity.Role, error) {
	var Role entity.Role

	err := r.config.DB.First(&Role).Error

	if err != nil {
		return Role, err
	}

	return Role, nil
}

// @Summary : Update Role
// @Description : Update Role by ID
// @Author : rasmadibbnu
func (r *RoleRepository) Update(Role entity.Role, ID int) (entity.Role, error) {
	err := r.config.DB.Where("id = ?", ID).Updates(&Role).Error

	if err != nil {
		return Role, err
	}

	return Role, nil
}

// @Summary : Delete Role
// @Description : Delete Role temporary
// @Author : rasmadibbnu
func (r *RoleRepository) Delete(ID int) (bool, error) {
	var Role entity.Role

	err := r.config.DB.Where("id = ?", ID).Delete(&Role).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

package repository

import (
	"etentnode-api/app/entity"
	"etentnode-api/config"
)

type StatusRepository struct {
	config config.Database
}

func NewStatusRepository(database config.Database) StatusRepository {
	return StatusRepository{
		config: database,
	}
}

// @Summary : Insert Status
// @Description : Insert Status to database
// @Author : rasmadibbnu
func (r *StatusRepository) Insert(Status entity.Status) (entity.Status, error) {
	err := r.config.DB.Create(&Status).Error

	if err != nil {
		return Status, err
	}

	return Status, nil
}

// @Summary : Get Statuss
// @Description : -
// @Author : rasmadibbnu
func (r *StatusRepository) FindAll(param map[string]interface{}) ([]entity.Status, error) {
	var Statuss []entity.Status

	err := r.config.DB.Where(param).Find(&Statuss).Error

	if err != nil {
		return Statuss, err
	}

	return Statuss, nil
}

// @Summary : Get Status
// @Description : Find Status by ID
// @Author : rasmadibbnu
func (r *StatusRepository) FindById(ID int) (entity.Status, error) {
	var Status entity.Status

	err := r.config.DB.First(&Status).Error

	if err != nil {
		return Status, err
	}

	return Status, nil
}

// @Summary : Update Status
// @Description : Update Status by ID
// @Author : rasmadibbnu
func (r *StatusRepository) Update(Status entity.Status, ID int) (entity.Status, error) {
	err := r.config.DB.Where("id = ?", ID).Updates(&Status).Error

	if err != nil {
		return Status, err
	}

	return Status, nil
}

// @Summary : Delete Status
// @Description : Delete Status temporary
// @Author : rasmadibbnu
func (r *StatusRepository) Delete(ID int) (bool, error) {
	var Status entity.Status

	err := r.config.DB.Where("id = ?", ID).Delete(&Status).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

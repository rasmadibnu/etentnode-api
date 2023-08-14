package repository

import (
	"etentnode-api/app/entity"
	"etentnode-api/config"

	"gorm.io/gorm"
)

type EventRepository struct {
	config config.Database
}

func NewEventRepository(database config.Database) EventRepository {
	return EventRepository{
		config: database,
	}
}

// @Summary : Insert Event
// @Description : Insert Event to database
// @Author : rasmadibbnu
func (r *EventRepository) Insert(Event entity.Event) (entity.Event, error) {
	err := r.config.DB.Create(&Event).Find(&Event).Error

	if err != nil {
		return Event, err
	}

	return Event, nil
}

// @Summary : Get Event
// @Description : -
// @Author : rasmadibbnu
func (r *EventRepository) FindAll(param map[string]interface{}) ([]entity.Event, error) {
	var Event []entity.Event

	err := r.config.DB.Where(param).Preload("EventUserHandling.User.Role").Preload("EventCategory", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Roles.Role", func(db2 *gorm.DB) *gorm.DB {
			return db2.Where("is_service = 1")
		}).Preload("Types").Preload("Fields")
	}).Preload("EventHandling", func(db *gorm.DB) *gorm.DB {
		return db.Preload("EventCategoryType").Preload("UserCreate.Role")
	}).Preload("EventCategoryType").Preload("Status").Preload("UserCreate.Role").Preload("UserAssign.Role").Preload("UserUpdate.Role").Order("created_at desc").Find(&Event).Error

	if err != nil {
		return Event, err
	}

	return Event, nil
}

// @Summary : Get Event
// @Description : -
// @Author : rasmadibbnu
func (r *EventRepository) ListByUserAssign(param map[string]interface{}, id string) ([]entity.Event, error) {
	var Event []entity.Event

	err := r.config.DB.Where(param).Preload("EventUserHandling.User.Role").Preload("EventCategory", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Roles.Role", func(db2 *gorm.DB) *gorm.DB {
			return db2.Where("is_service = 1")
		}).Preload("Types").Preload("Fields")
	}).Preload("EventHandling", func(db *gorm.DB) *gorm.DB {
		return db.Preload("EventCategoryType").Preload("UserCreate.Role")
	}).Preload("EventCategoryType").Preload("Status").Preload("UserCreate.Role").Preload("UserAssign.Role").Preload("UserUpdate.Role").Order("created_at desc").Find(&Event, "id in (select event_id from tr_event_user_handling where user_id = ?)", id).Error

	if err != nil {
		return Event, err
	}

	return Event, nil
}

// @Summary : Get Event
// @Description : Find Event by ID
// @Author : rasmadibbnu
func (r *EventRepository) FindById(ID int) (entity.Event, error) {
	var Event entity.Event

	err := r.config.DB.Preload("EventUserHandling.User.Role").Preload("UserAssign.Role").Preload("EventCategory", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Roles.Role", func(db2 *gorm.DB) *gorm.DB {
			return db2.Where("is_service = 1")
		}).Preload("Types").Preload("Fields")
	}).Preload("EventHandling", func(db *gorm.DB) *gorm.DB {
		return db.Preload("EventCategoryType").Preload("UserCreate.Role")
	}).Preload("EventCategoryType").Preload("Status").Preload("UserCreate.Role").Preload("UserAssign.Role").Preload("UserUpdate.Role").Where("id = ?", ID).First(&Event).Error

	if err != nil {
		return Event, err
	}

	return Event, nil
}

// @Summary : Get Event
// @Description : Find Event
// @Author : rasmadibbnu
func (r *EventRepository) GetCountEvent(param map[string]interface{}, ID int) (int64, error) {
	var count int64

	err := r.config.DB.Raw("select count(e.id) from tr_events e  left join tr_event_handling eh on e.id=eh.event_id where e.id = (select uh.event_id from tr_event_user_handling uh where uh.event_id = e.id and uh.user_id = ?) and e.status_id=2 and eh.id is null", ID).Scan(&count).Error

	if err != nil {
		return count, err
	}

	return count, nil
}

// @Summary : Update Event
// @Description : Update Event by ID
// @Author : rasmadibbnu
func (r *EventRepository) Update(Event entity.Event, ID int) (entity.Event, error) {
	err := r.config.DB.Debug().Where("id = ?", ID).Updates(&Event).Error

	if err != nil {
		return Event, err
	}

	return Event, nil
}

// @Summary : Delete Event
// @Description : Delete Event temporary
// @Author : rasmadibbnu
func (r *EventRepository) Delete(ID int) (bool, error) {
	var Event entity.Event

	err := r.config.DB.Where("id = ?", ID).Delete(&Event).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

// @Summary : Assign User Handling
// @Description : Assign User Handling to database
// @Author : rasmadibbnu
func (r *EventRepository) AssignUser(UserHandling []entity.EventUserHandling) ([]entity.EventUserHandling, error) {

	err := r.config.DB.Create(&UserHandling).Error

	if err != nil {
		return UserHandling, err
	}

	return UserHandling, nil
}

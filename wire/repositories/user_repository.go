package repositories

import (
	"wire/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(input *models.User) error
	FindAll() (users []models.User)
	FindById(id uint) (user *models.User, err error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepositoiry(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(input *models.User) error {
	return r.db.Create(input).Error
}

func (r *userRepository) FindAll() (users []models.User) {
	err := r.db.Find(users)
	if err != nil {
		return
	}

	return
}

func (r *userRepository) FindById(id uint) (user *models.User, err error) {

	err = r.db.First(user, "id", id).Error
	if err != nil {
		return
	}

	return
}

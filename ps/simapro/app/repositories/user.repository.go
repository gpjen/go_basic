package repositories

import (
	"errors"

	"github.com/gpjen/simapro/app/models"
	"github.com/gpjen/simapro/database"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmail(email string) (*models.User, error)
	IsDuplicateEmail(email string) error
	CreateUser(user *models.User) error
	GetUsers(page, perPage int) ([]*models.User, int64, error)
	GetUser(userId uint, preload bool) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(user *models.User) error
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := database.DB.Preload("Roles").Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) IsDuplicateEmail(email string) error {
	var user models.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil
	}

	return errors.New("email already exists")
}

func (r *userRepository) CreateUser(user *models.User) error {
	return database.DB.Create(user).Error
}

func (r *userRepository) GetUsers(page, perPage int) ([]*models.User, int64, error) {
	var users []*models.User
	var totalItems int64

	// count total items
	database.DB.Model(&models.User{}).Count(&totalItems)

	if err := database.DB.Offset((page - 1) * perPage).Limit(perPage).Find(&users).Error; err != nil {
		return nil, 0, err
	}
	return users, totalItems, nil
}

func (r *userRepository) GetUser(userId uint, preload bool) (*models.User, error) {
	var user models.User

	query := database.DB

	if preload {
		query = query.Preload("Roles")
	}

	if err := query.First(&user, userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) UpdateUser(user *models.User) error {
	db := database.DB.Begin()

	if err := db.Save(user).Error; err != nil {
		db.Rollback()
		return err
	}

	if err := db.Model(user).Association("Roles").Replace(user.Roles); err != nil {
		db.Rollback()
		return err
	}

	return db.Commit().Error

}

func (r *userRepository) DeleteUser(user *models.User) error {
	return database.DB.Delete(user).Error
}

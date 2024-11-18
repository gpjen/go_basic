package repository

import "clean_arsitektur/internal/domain/entity"

type UserRepository interface {
	GetAll() ([]entity.User, error)
	GetByID(id uint) (*entity.User, error)
}

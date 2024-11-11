package repository

import "clean_arsitektur/internal/domain/entity"

type UserRepository interface {
	GetUserByID(id uint) (*entity.User, error)
}

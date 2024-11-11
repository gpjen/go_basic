package repository

import (
	"clean_arsitektur/internal/domain/entity"
	"clean_arsitektur/internal/domain/repository"
	"errors"
)

type UserRepo struct {
	users map[uint]*entity.User
}

func NewUserRepo() repository.UserRepository {
	return &UserRepo{
		users: map[uint]*entity.User{
			1: {ID: 1, Name: "John Doe", Email: "john@example.com"},
			2: {ID: 2, Name: "Jane Smith", Email: "jane@example.com"},
		},
	}
}

func (r *UserRepo) GetUserByID(id uint) (*entity.User, error) {
	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

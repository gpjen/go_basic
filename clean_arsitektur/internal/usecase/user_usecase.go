package usecase

import (
	"clean_arsitektur/internal/domain/entity"
	"clean_arsitektur/internal/domain/repository"
)

type UserUsecase interface {
	GetUser(id uint) (*entity.User, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo: repo}
}

func (u *userUsecase) GetUser(id uint) (*entity.User, error) {
	return u.userRepo.GetByID(id)
}

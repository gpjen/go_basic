package usecase

import "gracefullac/internal/domain"

type userUseCase struct {
	userRepo domain.UserRepository
}

func NewUserUseCase(userRepo domain.UserRepository) domain.UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}

func (uc *userUseCase) Create(user *domain.User) error {
	return uc.userRepo.Create(user)
}

func (uc *userUseCase) FindAll() []domain.User {
	return uc.userRepo.FindAll()
}

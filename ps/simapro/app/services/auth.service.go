package services

import (
	"errors"

	"github.com/gpjen/simapro/app/dto"
	"github.com/gpjen/simapro/app/repositories"
	"github.com/gpjen/simapro/utils"
)

type AuthService interface {
	Login(input dto.UserLoginDTO) (res dto.UserLoginResponseDTO, err error)
}

type authService struct {
	userRepository repositories.UserRepository
}

func NewAuthService(userRepository repositories.UserRepository) AuthService {
	return &authService{userRepository}
}

func (s *authService) Login(input dto.UserLoginDTO) (res dto.UserLoginResponseDTO, err error) {

	user, err := s.userRepository.FindByEmail(input.Email)
	if err != nil {
		return res, errors.New("email/password not found")
	}

	ok := utils.CheckPasswordHash(input.Password, user.Password)
	if !ok {
		return res, errors.New("email/password not found")
	}

	userCtx := dto.UserContext{
		ID:    user.ID,
		Email: user.Email,
	}

	for idx, role := range user.Roles {
		if idx == 0 {
			userCtx.SelectedRole = role.ID
		}

		userCtx.Roles = append(userCtx.Roles, role.ID)
	}

	res.Email = user.Email
	res.Token, err = utils.GenerateToken(userCtx)
	if err != nil {
		return
	}

	return

}

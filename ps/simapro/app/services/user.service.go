package services

import (
	"github.com/gpjen/simapro/app/dto"
	"github.com/gpjen/simapro/app/models"
	"github.com/gpjen/simapro/app/repositories"
	"github.com/gpjen/simapro/utils"
)

type UserService interface {
	CreateUser(input *dto.CreateUserDTO) error
	GetAllUsers(query utils.QueryParams) ([]*dto.UserDTO, utils.Meta, error)
	GetUserById(id uint) (*dto.UserDetailDTO, error)
	UpdateUser(id uint, dto *dto.UpdateUserDTO) error
	DeleteUser(id uint) error
}

type userService struct {
	userRepository repositories.UserRepository
	roleRepository repositories.RoleRepository
}

func NewUserService(userRepository repositories.UserRepository, roleRepository repositories.RoleRepository) UserService {
	return &userService{userRepository, roleRepository}
}

func (s *userService) CreateUser(input *dto.CreateUserDTO) error {
	err := s.userRepository.IsDuplicateEmail(input.Email)
	if err != nil {
		return err
	}

	hashed, _ := utils.HashPassword(input.Password)
	user := &models.User{
		Name:      input.Name,
		Email:     input.Email,
		Password:  hashed,
		CreatedBy: input.UserReqEmail,
	}

	var roles []*models.Role

	if len(input.Roles) > 0 {
		roles, err = s.roleRepository.GetRoleByIds(input.Roles)
		if err != nil {
			return err
		}
	}

	for _, role := range roles {
		user.Roles = append(user.Roles, *role)
	}

	err = s.userRepository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) GetAllUsers(query utils.QueryParams) ([]*dto.UserDTO, utils.Meta, error) {
	page, PerPage := utils.GetPaginationParams(query.Page, query.PerPage)

	users, totalItems, err := s.userRepository.GetUsers(page, PerPage)
	if err != nil {
		return nil, utils.Meta{}, err
	}

	res := []*dto.UserDTO{}

	for _, user := range users {
		res = append(res, user.ToUserDTO())
	}

	meta := utils.MetaPagination(
		page, PerPage, len(users), int(totalItems),
	)

	return res, meta, err
}

func (s *userService) GetUserById(id uint) (*dto.UserDetailDTO, error) {
	user, err := s.userRepository.GetUser(id, true)
	if err != nil {
		return nil, err
	}

	return user.ToUserDetailDTO(), err
}

func (s *userService) UpdateUser(id uint, dto *dto.UpdateUserDTO) error {
	user, err := s.userRepository.GetUser(id, false)
	if err != nil {
		return err
	}

	if dto.Name != "" && user.Name != dto.Name {
		user.Name = dto.Name
	}

	if dto.Email != "" && user.Email != dto.Email {
		err = s.userRepository.IsDuplicateEmail(dto.Email)
		if err != nil {
			return err
		}

		user.Email = dto.Email
	}

	if dto.Password != "" {
		hashed, _ := utils.HashPassword(dto.Password)
		user.Password = hashed
	}

	for _, role := range dto.Roles {
		role, err := s.roleRepository.GetRole(role)
		if err != nil {
			return err
		}

		user.Roles = append(user.Roles, *role)
	}

	err = s.userRepository.UpdateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) DeleteUser(id uint) error {
	user, err := s.userRepository.GetUser(id, false)
	if err != nil {
		return err
	}

	err = s.userRepository.DeleteUser(user)
	if err != nil {
		return err
	}

	return nil
}

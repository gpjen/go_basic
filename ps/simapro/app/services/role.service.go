package services

import (
	"github.com/gpjen/simapro/app/dto"
	"github.com/gpjen/simapro/app/models"
	"github.com/gpjen/simapro/app/repositories"
	"github.com/gpjen/simapro/utils"
)

type RoleService interface {
	CreateRole(input *dto.CreateRoleDTO) error
	GetAllRoles(query utils.QueryParams) ([]*dto.RoleDTO, utils.Meta, error)
	GetRoleById(id uint) (*dto.RoleDetailDTO, error)
	UpdateRole(id uint, input *dto.UpdateRoleDTO) error
	DeleteRole(id uint) error

	GetRoleMenuPermision(roleId uint, param *dto.ParamQueryMenuPermissionDTO) (*dto.MenuPermissionDTO, error)
}

type roleService struct {
	roleRepository repositories.RoleRepository
}

func NewRoleService(roleRepository repositories.RoleRepository) RoleService {
	return &roleService{
		roleRepository: roleRepository,
	}
}

func (s *roleService) CreateRole(input *dto.CreateRoleDTO) error {
	err := s.roleRepository.IsDuplicateName(input.Name)
	if err != nil {
		return err
	}

	role := &models.Role{
		Name:        input.Name,
		Description: input.Description,
		Active:      true,
		CreatedBy:   input.RoleReqEmail,
	}

	err = s.roleRepository.CreateRole(role)
	if err != nil {
		return err
	}
	return nil
}

func (s *roleService) GetAllRoles(query utils.QueryParams) ([]*dto.RoleDTO, utils.Meta, error) {
	page, PerPage := utils.GetPaginationParams(query.Page, query.PerPage)

	roles, totalItems, err := s.roleRepository.GetRoles(page, PerPage)
	if err != nil {
		return nil, utils.Meta{}, err
	}

	res := []*dto.RoleDTO{}

	for _, role := range roles {
		res = append(res, role.ToRoleDTO())
	}

	meta := utils.MetaPagination(
		page, PerPage, len(roles), int(totalItems),
	)

	return res, meta, err
}

func (s *roleService) GetRoleById(id uint) (*dto.RoleDetailDTO, error) {
	role, err := s.roleRepository.GetRole(id)
	if err != nil {
		return nil, err
	}

	return role.ToRoleDetailDTO(), err
}

func (s *roleService) UpdateRole(id uint, input *dto.UpdateRoleDTO) error {
	role, err := s.roleRepository.GetRole(id)
	if err != nil {
		return err
	}

	if role.Name != input.Name {
		err := s.roleRepository.IsDuplicateName(input.Name)
		if err != nil {
			return err
		}

		role.Name = input.Name
	}

	role.Description = input.Description
	role.Active = input.Active

	err = s.roleRepository.UpdateRole(role)
	if err != nil {
		return err
	}

	return nil
}

func (s *roleService) DeleteRole(id uint) error {
	role, err := s.roleRepository.GetRole(id)
	if err != nil {
		return err
	}

	err = s.roleRepository.DeleteRole(role)
	if err != nil {
		return err
	}

	return nil
}

func (s *roleService) GetRoleMenuPermision(roleId uint, param *dto.ParamQueryMenuPermissionDTO) (*dto.MenuPermissionDTO, error) {
	role, err := s.roleRepository.GetRoleMenuPermision(roleId, param)
	if err != nil {
		return &dto.MenuPermissionDTO{}, err
	}

	return role.ToRoleMenuPermision(), nil
}

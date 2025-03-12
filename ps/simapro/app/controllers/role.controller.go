package controllers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/gpjen/simapro/app/dto"
	"github.com/gpjen/simapro/app/services"
	"github.com/gpjen/simapro/utils"
)

type RoleController struct {
	service services.RoleService
}

func NewRoleController(service services.RoleService) *RoleController {
	return &RoleController{
		service: service,
	}
}

// Get List of Roles godoc
// @Summary Get List of Roles
// @Description Get List of Roles
// @Tags Roles
// @Accept json
// @Produce json
// @Param page query int false "Page"
// @Param perPage query int false "PerPage"
// @Param sort query string false "Sort"
// @Param search query string false "Search"
// @Param status query string false "Status"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ErrorResponse
// @Router /roles [get]
func (ctrl *RoleController) GetRoles(c *fiber.Ctx) error {
	q := new(utils.QueryParams)
	if err := c.QueryParser(q); err != nil {
		return err
	}

	h := &utils.ResponseHandler{}

	roles, meta, err := ctrl.service.GetAllRoles(*q)
	if err != nil {
		return h.InternalServerError(c, []string{err.Error()})
	}

	return h.Ok(c, roles, "roles fetched successfully", &meta)
}

// Create Role godoc
// @Summary Create Role
// @Description Create Role
// @Tags Roles
// @Accept json
// @Produce json
// @Param role body dto.CreateRoleDTO true "Role"
// @Success 201 {object} utils.ResponseData
// @Failure 400 {object} utils.ErrorResponse
// @Router /roles [post]
func (ctrl *RoleController) CreateRole(c *fiber.Ctx) error {
	h := &utils.ResponseHandler{}
	var dto dto.CreateRoleDTO
	if err := c.BodyParser(&dto); err != nil {
		return h.BadRequest(c, []string{err.Error()})
	}

	dto.RoleReqEmail = utils.GetUserSession(c).Email
	err := ctrl.service.CreateRole(&dto)
	if err != nil {
		return h.BadRequest(c, []string{err.Error()})
	}
	return h.Created(c, nil, "role created successfully")
}

// Get Role godoc
// @Summary Get Role
// @Description Get Role
// @Tags Roles
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ErrorResponse
// @Router /roles/{id} [get]
func (ctrl *RoleController) GetRole(c *fiber.Ctx) error {
	id := utils.StringToUint(c.Params("id"))

	h := &utils.ResponseHandler{}

	role, err := ctrl.service.GetRoleById(id)
	if err != nil {
		return h.NotFound(c, []string{err.Error()})
	}

	return h.Ok(c, role, "roles fetched successfully", nil)
}

// Update Role godoc
// @Summary Update Role
// @Description Update Role
// @Tags Roles
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Param role body dto.UpdateRoleDTO true "Role"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ErrorResponse
// @Router /roles/{id} [put]
func (ctrl *RoleController) UpdateRole(c *fiber.Ctx) error {
	id := utils.StringToUint(c.Params("id"))

	h := &utils.ResponseHandler{}

	var dto *dto.UpdateRoleDTO
	if err := c.BodyParser(&dto); err != nil {
		return h.BadRequest(c, []string{err.Error()})
	}

	err := ctrl.service.UpdateRole(id, dto)
	if err != nil {
		return h.InternalServerError(c, []string{err.Error()})
	}

	return h.Ok(c, nil, "role updated successfully", nil)
}

// Delete Role godoc
// @Summary Delete Role
// @Description Delete Role
// @Tags Roles
// @Produce json
// @Accept json
// @Param id path string true "ID"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ErrorResponse
// @Router /roles/{id} [delete]
func (ctrl *RoleController) DeleteRole(c *fiber.Ctx) error {
	id := utils.StringToUint(c.Params("id"))

	h := &utils.ResponseHandler{}
	err := ctrl.service.DeleteRole(id)
	if err != nil {
		if err.Error() == "role not found" {
			return h.NotFound(c, []string{err.Error()})
		}
		return h.InternalServerError(c, []string{err.Error()})
	}

	return h.Ok(c, nil, "Role deleted successfully", nil)
}

func (ctrl *RoleController) GetRoleMenuPermision(c *fiber.Ctx) error {
	id := utils.StringToUint(c.Params("id"))

	dto := &dto.ParamQueryMenuPermissionDTO{
		CanView:    utils.StringToBool(c.Query("can_view")),
		MenuActive: utils.StringToBool(c.Query("menu_active")),
	}

	h := &utils.ResponseHandler{}

	role, err := ctrl.service.GetRoleMenuPermision(id, dto)
	if err != nil {
		return h.NotFound(c, []string{err.Error()})
	}

	return h.Ok(c, role, "roles fetched successfully", nil)
}

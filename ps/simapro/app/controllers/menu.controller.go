package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gpjen/simapro/app/dto"
	"github.com/gpjen/simapro/app/services"
	"github.com/gpjen/simapro/utils"
)

type MenuController struct {
	MenuService services.MenuService
}

func NewMenuController(menuService services.MenuService) *MenuController {
	return &MenuController{menuService}
}

func (ctrl *MenuController) GetMenuHierarchy(c *fiber.Ctx) error {
	param := &dto.ParamQueryMenuDTO{}
	r := utils.ResponseHandler{}

	if err := c.QueryParser(param); err != nil {
		return r.BadRequest(c, []string{err.Error()})
	}

	menus, err := ctrl.MenuService.GetMenuHierarchy(param)
	if err != nil {
		return r.BadRequest(c, []string{err.Error()})
	}

	return r.Ok(c, menus, "menu hierarchy fetched successfully", nil)
}

func (ctrl *MenuController) MenuUpdates(c *fiber.Ctx) error {

	dto := utils.GetValidatedDTO(c).(dto.UpdateMenusDto)
	r := utils.ResponseHandler{}

	err := ctrl.MenuService.MenuUpdates(dto)
	if err != nil {
		return r.BadRequest(c, []string{err.Error()})
	}

	return r.Ok(c, nil, "menu hierarchy updated successfully", nil)
}

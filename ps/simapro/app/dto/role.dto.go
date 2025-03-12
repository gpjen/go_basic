package dto

type RoleDTO struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type RoleDetailDTO struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedBy   string `json:"created_by"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type CreateRoleDTO struct {
	Name         string `json:"name"  validate:"required,min=3,max=32"`
	Description  string `json:"description" validate:""`
	RoleReqEmail string `json:"-" form:"-"`
}

type UpdateRoleDTO struct {
	Name        string `json:"name"  validate:"required,min=3,max=32"`
	Description string `json:"description" validate:"required"`
	Active      bool   `json:"active"`
}

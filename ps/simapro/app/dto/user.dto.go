package dto

type UserDTO struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"  validate:"required,min=3,max=32"`
	Email string `json:"email" validate:"required,email"`
}

type UserDetailDTO struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedBy string    `json:"created_by"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	Roles     []RoleDTO `json:"roles"`
}

type CreateUserDTO struct {
	Name     string `json:"name"  validate:"required,min=3,max=32"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Roles    []uint `json:"roles" validate:"required"`

	UserReqEmail string `json:"-" form:"-"`
}

type UpdateUserDTO struct {
	Name     string `json:"name"  validate:"omitempty,min=3,max=32"`
	Email    string `json:"email" validate:"omitempty,email"`
	Password string `json:"password" validate:"omitempty"`
	Roles    []uint `json:"roles" validate:"required"`
}

type UserContext struct {
	ID           uint   `json:"id"`
	Email        string `json:"email"`
	SelectedRole uint   `json:"selected_role"`
	Roles        []uint `json:"roles"`
}

type UserLoginDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserLoginResponseDTO struct {
	Email string `json:"email" validate:"required,email"`
	Token string `json:"token"`
}

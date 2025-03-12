package dto

type ParamQueryMenuPermissionDTO struct {
	CanView    bool `json:"can_view"`
	MenuActive bool `json:"menu_active"`
}

type MenuDTO struct {
	MenuID     uint   `json:"menu_id"`
	Name       string `json:"name"`
	Route      string `json:"route"`
	Icon       string `json:"icon"`
	CanView    bool   `json:"can_view"`
	CanShow    bool   `json:"can_show"`
	CanCreate  bool   `json:"can_create"`
	CanUpdate  bool   `json:"can_update"`
	CanDelete  bool   `json:"can_delete"`
	Parent     *uint  `json:"parent"`
	Order      uint16 `json:"order"`
	MenuActive bool   `json:"menu_active"`
}

type MenuPermissionDTO struct {
	RoleID      uint      `json:"role_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Menus       []MenuDTO `json:"menus"`
}

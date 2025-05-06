package dto

type ParamQueryMenuDTO struct {
	Active bool   `json:"active"`
	Order  string `json:"order"`
}

type UpdateMenusDataDTO struct {
	Id       uint   `json:"id" validate:"required,ne_field=IdParent"`
	IdParent *uint  `json:"id_parent"`
	Icon     string `json:"icon" validate:"max=100"`
	Order    uint16 `json:"order"`
	Active   bool   `json:"active"`
}

type UpdateMenusDto struct {
	Data []UpdateMenusDataDTO `json:"data" validate:"required,min=1,dive"`
}

package models

type Menu struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"size:100;not null"`
	Route    string `json:"route" gorm:"size:100;not null"`
	IdParent *uint  `json:"id_parent" gorm:"index"`
	Icon     string `json:"icon" gorm:"size:100;not null"`
	Order    uint16 `json:"order" gorm:"not null"`
	Active   bool   `json:"active" gorm:"default:true"`
	Parent   *Menu  `json:"-" gorm:"foreignKey:IdParent;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Children []Menu `json:"children" gorm:"foreignKey:IdParent"`
}

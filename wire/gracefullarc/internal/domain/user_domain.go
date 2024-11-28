package domain

type User struct {
	ID    uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name  string `json:"name" gorm:"not null;size:100"`
	Email string `json:"email" gorm:"unique;not null;size:100"`
}

type UserRepository interface {
	Create(user *User) error
	FindAll() []User
}

type UserUseCase interface {
	Create(user *User) error
	// FindAll() []User
}

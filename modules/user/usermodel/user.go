package usermodel

import "gofiber-boilerplate/base"

type UserModel struct {
	base.BaseModel
	FirstName   string  `json:"first_name" gorm:"not null;"`
	LastName    string  `json:"last_name" gorm:"not null;"`
	PhoneNumber string  `json:"phone_number" gorm:"not null;unique;"`
	Address     string  `json:"address" gorm:"not null;"`
	Pin         *string `json:"pin" gorm:"not null;"`
	Balance     int     `json:"balance" gorm:"not null;default:0;"`
}

func (UserModel) TableName() string {
	return "users"
}

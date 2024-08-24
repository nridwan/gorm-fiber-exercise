package userdto

import "gofiber-boilerplate/modules/user/usermodel"

type RegisterDTO struct {
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	Pin         string `json:"pin" validate:"required,len=6"`
	PhoneNumber string `json:"phone_number" validate:"required,startswith=08"`
	Address     string `json:"address" validate:"required"`
}

func (dto *RegisterDTO) ToModel() *usermodel.UserModel {
	return &usermodel.UserModel{
		ReadonlyUserModel: usermodel.ReadonlyUserModel{
			FirstName:   dto.FirstName,
			LastName:    dto.LastName,
			PhoneNumber: dto.PhoneNumber,
			Address:     dto.Address,
		},
		Pin: &dto.Pin,
	}
}

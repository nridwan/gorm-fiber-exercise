package userdto

import (
	"gofiber-boilerplate/base"
	"gofiber-boilerplate/modules/user/usermodel"
)

type UserDTO = usermodel.ReadonlyUserModel

func MapUserModelToDTO(model *usermodel.UserModel) *UserDTO {
	return &UserDTO{
		BaseModel: base.BaseModel{
			ID:        model.ID,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
			DeletedAt: model.DeletedAt,
		},
		FirstName:   model.FirstName,
		LastName:    model.LastName,
		PhoneNumber: model.PhoneNumber,
		Address:     model.Address,
	}
}

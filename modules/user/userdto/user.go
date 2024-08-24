package userdto

import (
	"gofiber-boilerplate/modules/user/usermodel"
)

type UserDTO = usermodel.UserModel

func MapUserModelToDTO(model *usermodel.UserModel) *UserDTO {
	return model
}

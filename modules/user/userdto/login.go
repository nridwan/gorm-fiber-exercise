package userdto

import (
	"gofiber-boilerplate/modules/jwt"
)

type LoginDTO struct {
	PhoneNumber string `json:"phone_number" validate:"required,startswith=08"`
	Pin         string `json:"pin" validate:"required"`
}

type LoginResponseDTO = jwt.JWTTokenModel

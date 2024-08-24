package userdto

import (
	"gofiber-boilerplate/modules/user/usermodel"
	"time"

	"github.com/google/uuid"
)

type UserDTO struct {
	ID          uuid.UUID  `json:"id" gorm:"type:uuid;not null;primaryKey;default:uuid_generate_v4()"`
	FirstName   string     `json:"first_name" gorm:"not null;"`
	LastName    string     `json:"last_name" gorm:"not null;"`
	PhoneNumber string     `json:"phone_number" gorm:"not null;unique;"`
	Address     string     `json:"address" gorm:"not null;"`
	CreatedAt   *time.Time `json:"created_date,omitempty" gorm:"not null;"`
	UpdatedAt   *time.Time `json:"updated_date,omitempty" gorm:"not null;"`
}

func MapUserModelToDTO(model *usermodel.UserModel) *UserDTO {
	return &UserDTO{
		ID:          model.ID,
		FirstName:   model.FirstName,
		LastName:    model.LastName,
		PhoneNumber: model.PhoneNumber,
		Address:     model.Address,
		UpdatedAt:   model.UpdatedAt,
	}
}

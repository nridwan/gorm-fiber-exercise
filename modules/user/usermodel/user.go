package usermodel

type UserModel struct {
	ReadonlyUserModel
	Pin *string `json:"pin" gorm:"not null;"`
}

func (UserModel) TableName() string {
	return "users"
}

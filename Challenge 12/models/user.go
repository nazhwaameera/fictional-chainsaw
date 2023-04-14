package models

import (
	"Challenge_12/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// User represents the model for an user
type User struct {
	GormModel
	Username     string        `gorm:"not null;uniqueIndex" form:"username" json:"username" valid:"required~Your username is required"`
	Email        string        `gorm:"not null;uniqueIndex" form:"email" json:"email" valid:"required~Your email is required, email~Invalid email format"`
	Password     string        `gorm:"not null" form:"password" json:"password" valid:"required~Your password is required, minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age          uint          `gorm:"not null" form:"age" json:"age" valid:"required~Your age is required,range(8|100)~You do not meet the age requirement"`
	SocialMedias []SocialMedia `gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL;" json:"social_medias"`
	Photos       []Photo       `gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL;" json:"photos"`
	Comments     []Comment     `gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL;" json:"comments"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	hashedPass, err := helpers.HashPassword(u.Password)

	if err != nil {
		return
	}

	u.Password = hashedPass

	return
}

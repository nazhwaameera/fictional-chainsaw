package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// SocialMedia represents the model for a social media
type SocialMedia struct {
	GormModel
	Name   string `json:"name" form:"name" valid:"required~Your social media's name is required"`
	URL    string `json:"social_media_url" form:"social_media_url" valid:"required~Your social media's URL is required"`
	UserID uint
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (s *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(s)

	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}

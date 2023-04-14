package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// Photo represents the model for a photo
type Photo struct {
	GormModel
	Title   string `json:"title" form:"title" valid:"required~Photo title is required"`
	Caption string `json:"caption" form:"caption"`
	URL     string `json:"photo_url" form:"photo_url" valid:"required~Photo URL is required"`
	UserID  uint
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(p)

	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}

package models

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID        uint      `gorm:"primaryKey json:"id"`
	Name_book string    `gorm:"not null;unique;type:varchar(50)" json:"name_book"`
	Author    string    `gorm:"type:varchar(50)" json:"author"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt`
}

func (b *Book) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Book before create()")

	if len(b.Author) < 2 {
		err = errors.New("Author name is too short.")
	}

	return
}

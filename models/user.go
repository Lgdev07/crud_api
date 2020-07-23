package models

import (
	"github.com/jinzhu/gorm"
)

// User model
type User struct {
	gorm.Model
	Name  string `gorm:"size:100;not null" json:"name"`
	Email string `gorm:"size:100;not null;unique" json:"email"`
	Type  string `gorm:"size:50;not null" json:"type"`
}

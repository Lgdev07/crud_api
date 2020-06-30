package models

import (
	"github.com/jinzhu/gorm"
)

// Store struct
type Store struct {
	gorm.Model
	ID     uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Name   string `gorm:"size:100;not null" json:"name"`
	Type   string `gorm:"size:100;not null" json:"type"`
	Active bool   `gorm:"default:false" json:"ref"`
}

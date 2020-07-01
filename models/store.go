package models

import (
	"strings"

	"github.com/jinzhu/gorm"
)

// Store struct
type Store struct {
	gorm.Model
	Name   string `gorm:"size:100;not null" json:"name"`
	Type   string `gorm:"size:100;not null" json:"type"`
	Active bool   `gorm:"default:false" json:"active"`
}

// Prepare to clean the fields
func (s *Store) Prepare() {
	s.Name = strings.TrimSpace(s.Name)
	s.Type = strings.TrimSpace(s.Type)
}

// Save to create a new register on Database
func (s *Store) Save(db *gorm.DB) (*Store, error) {
	var err error

	err = db.Debug().Create(&s).Error
	if err != nil {
		return &Store{}, err
	}
	return s, nil
}

// GetStores Get all stores in the database
func GetStores(db *gorm.DB) (*[]Store, error) {
	stores := []Store{}

	if err := db.Debug().Table("stores").Find(&stores).Error; err != nil {
		return &[]Store{}, err
	}

	return &stores, nil
}

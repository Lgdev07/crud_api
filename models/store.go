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
func (s *Store) GetStores(db *gorm.DB) (*[]Store, error) {
	stores := []Store{}

	if err := db.Debug().Table("stores").Find(&stores).Error; err != nil {
		return &[]Store{}, err
	}

	return &stores, nil
}

// GetStoreByID return a store by the id parameter
func GetStoreByID(id int, db *gorm.DB) (*Store, error) {
	store := Store{}

	if err := db.Debug().Table("stores").Where("id = ?", id).First(&store).Error; err != nil {
		return &Store{}, err
	}
	return &store, nil

}

// DeleteStore deletes a store based in the id
func DeleteStore(id int, db *gorm.DB) error {
	store := Store{}

	if err := db.Debug().Table("stores").Where("id = ?", id).Delete(store).Error; err != nil {
		return err
	}

	return nil

}

// UpdateStore update a record based in the id
func (s *Store) UpdateStore(id int, db *gorm.DB) (*Store, error) {
	if err := db.Debug().Table("stores").Where("id = ?", id).Updates(Store{
		Name:   s.Name,
		Type:   s.Type,
		Active: s.Active,
	}).Error; err != nil {
		return &Store{}, err
	}

	return s, nil
}

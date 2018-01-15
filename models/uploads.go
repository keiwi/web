package models

import (
	"github.com/jinzhu/gorm"

	// import mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Uploads struct
type Uploads struct {
	gorm.Model    `json:"-"`
	Name          string `gorm:"not null;unique" json:"name"`
	Checksum      string `gorm:"not null" json:"checksum"`
	Version       string `gorm:"not null" json:"version"`
	Patch         bool   `json:"patch"`
	PatchChecksum string `gorm:"not null" json:"patch_checksum"`
}

// UploadsManager struct
type UploadsManager struct {
	db *DB
}

// NewUploadsManager - Creates a new *UploadsManager that can be used for managing uploads.
func NewUploadsManager(db *DB) (*UploadsManager, error) {
	db.AutoMigrate(&Uploads{})
	manager := UploadsManager{}
	manager.db = db
	return &manager, nil
}

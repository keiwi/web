package models

import (
	"github.com/jinzhu/gorm"

	// import mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// AlertsOptions struct
type AlertsOptions struct {
	gorm.Model `json:"-"`
	ClientID   int    `json:"client_id"`
	CommandID  int    `json:"command_id"`
	Alert      string `gorm:"not null" json:"alert"`
	Value      string `gorm:"not null" json:"value"`
	Count      int    `json:"count"`
	Delay      int    `json:"delay"`
	Service    string `gorm:"not null" json:"service"`
}

// AlertsOptionsManager struct
type AlertsOptionsManager struct {
	db *DB
}

// NewAlertsOptionsManager - Creates a new *AlertsOptionsManager that can be used for managing alertsOptions.
func NewAlertsOptionsManager(db *DB) (*AlertsOptionsManager, error) {
	db.AutoMigrate(&AlertsOptions{})
	manager := AlertsOptionsManager{}
	manager.db = db
	return &manager, nil
}

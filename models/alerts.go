package models

import (
	"github.com/jinzhu/gorm"
	"time"

	// import mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Alerts struct
type Alerts struct {
	gorm.Model `json:"-"`
	AlertID    int       `json:"alert_id"`
	ClientID   int       `json:"client_id"`
	Timestamp  time.Time `json:"timestamp"`
	Value      string    `gorm:"not null" json:"value"`
}

// AlertsManager struct
type AlertsManager struct {
	db *DB
}

// NewAlertsManager - Creates a new *AlertsManager that can be used for managing Alerts.
func NewAlertsManager(db *DB) (*AlertsManager, error) {
	db.AutoMigrate(&Alerts{})
	manager := AlertsManager{}
	manager.db = db
	return &manager, nil
}

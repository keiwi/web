package models

import (
	"time"

	// import mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Check struct
type Check struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
	CommandID uint       `json:"command_id"`
	ClientID  uint       `json:"client_id"`
	Response  string     `gorm:"not null" json:"response"`
	Checked   bool       `json:"checked"`
	Error     bool       `json:"error"`
	Finished  bool       `json:"finished"`
}

// ChecksManager struct
type ChecksManager struct {
	db *DB
}

// NewChecksManager - Creates a new *ChecksManager that can be used for managing checks.
func NewChecksManager(db *DB) (*ChecksManager, error) {
	db.AutoMigrate(&Check{})
	manager := ChecksManager{}
	manager.db = db
	return &manager, nil
}

// GetChecksBetweenDate grabs all checks with a specific command id
// between 2 different dates.
func (state *ChecksManager) GetChecksBetweenDate(from, to string, commandID uint) []Check {
	checks := []Check{}
	state.db.Where("command_id = ? AND created_at >= FROM_UNIXTIME(?) AND created_at <= FROM_UNIXTIME(?)", commandID, from, to).Find(&checks)
	return checks
}

// GetChecksBetweenDateClient is similiar to GetChecksBetweenDate
// where it grabs all checks between 2 different dates with a specific
// command id, the difference is that this function also grabs checks
// with a specific command_id
func (state *ChecksManager) GetChecksBetweenDateClient(from, to string, commandID, clientID uint) ([]Check, error) {
	checks := []Check{}
	err := state.db.Where("client_id = ? AND command_id = ? AND created_at >= ? AND created_at <= ?", clientID, commandID, from, to).Find(&checks).Error
	return checks, err
}

// Create inserts a new Check in the database
func (state *ChecksManager) Create(check *Check) error {
	return state.db.Create(check).Error
}

// FindAll grabs all of the existing Checks
func (state *ChecksManager) FindAll() ([]Check, error) {
	checks := []Check{}
	err := state.db.Find(&checks).Error
	return checks, err
}

// Find tries to find an existing Check from a id
func (state *ChecksManager) Find(id uint) (*Check, error) {
	check := Check{}
	err := state.db.Where("id = ?", id).Find(&check).Error
	return &check, err
}

// FindWithClientIDAndCommandID tries to find an existing Check from a client ID and command ID
func (state *ChecksManager) FindWithClientIDAndCommandID(client, command uint) ([]*Check, error) {
	checks := []*Check{}
	err := state.db.Where("client_id = ? AND command_id = ?", client, command).Order("created_at desc").Find(&checks).Error
	return checks, err
}

// Save saves a Check from a Check struct
func (state *ChecksManager) Save(check *Check) error {
	return state.db.Save(&check).Error
}

// Delete removes an existing Check
// from the database.
func (state *ChecksManager) Delete(check *Check) error {
	return state.db.Delete(&check).Error
}

// DeleteWithID is a wrapper for Delete that
// creates a new Check instance from ID and then run the Delete function
func (state *ChecksManager) DeleteWithID(id uint) error {
	m := &Check{}
	m.ID = id
	return state.Delete(m)
}

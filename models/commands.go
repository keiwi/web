package models

import (
	"time"

	// import mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Command struct
type Command struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `sql:"index" json:"-"`
	Command     string     `gorm:"not null" json:"command"`
	Namn        string     `gorm:"not null;unique" json:"namn"`
	Description string     `gorm:"not null" json:"description"`
	Format      string     `json:"format"`
}

// CommandsManager struct
type CommandsManager struct {
	db *DB
}

// NewCommandsManager - Creates a new *CommandsManager that can be used for managing commands.
func NewCommandsManager(db *DB) (*CommandsManager, error) {
	db.AutoMigrate(&Command{})
	manager := CommandsManager{}
	manager.db = db
	return &manager, nil
}

// Create inserts a new command in the database
func (state *CommandsManager) Create(cmd *Command) error {
	return state.db.Create(cmd).Error
}

// FindAll grabs all of the existing commands, note that this
// is base commands not GroupCommands
func (state *CommandsManager) FindAll() ([]Command, error) {
	cmds := []Command{}
	err := state.db.Find(&cmds).Error
	return cmds, err
}

// Find tries to find an existing command from a id
func (state *CommandsManager) Find(id uint) (*Command, error) {
	cmd := Command{}
	err := state.db.Where("id = ?", id).Find(&cmd).Error
	return &cmd, err
}

// Save saves a command from a command struct
func (state *CommandsManager) Save(cmd *Command) error {
	return state.db.Save(&cmd).Error
}

// Delete removes an existing command
// from the database.
func (state *CommandsManager) Delete(cmd *Command) error {
	return state.db.Delete(&cmd).Error
}

// DeleteWithID is a wrapper for Delete that
// creates a new CMD instance from ID and then run the Delete function
func (state *CommandsManager) DeleteWithID(id uint) error {
	m := &Command{}
	m.ID = id
	return state.Delete(m)
}

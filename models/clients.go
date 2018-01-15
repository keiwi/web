package models

import (
	"time"

	// import mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Client struct
type Client struct {
	ID         uint       `gorm:"primary_key" json:"id"`
	CreatedAt  time.Time  `json:"-"`
	UpdatedAt  time.Time  `json:"-"`
	DeletedAt  *time.Time `sql:"index" json:"-"`
	GroupNames string     `gorm:"not null" json:"group_names"`
	IP         string     `gorm:"not null" json:"ip"`
	Namn       string     `gorm:"not null;unique" json:"namn"`
}

// ClientsManager struct
type ClientsManager struct {
	db *DB
}

// NewClientsManager - Creates a new *ClientsManager that can be used for managing clients.
func NewClientsManager(db *DB) (*ClientsManager, error) {
	db.AutoMigrate(&Client{})
	manager := ClientsManager{}
	manager.db = db
	return &manager, nil
}

// Create inserts a new client into the database
func (state *ClientsManager) Create(client *Client) error {
	return state.db.Create(client).Error
}

// Delete removes an existing Client
// from the database.
func (state *ClientsManager) Delete(client *Client) error {
	return state.db.Delete(&client).Error
}

// DeleteWithID is a wrapper for Delete that
// creates a new client instance from ID and then run the Delete function
func (state *ClientsManager) DeleteWithID(id uint) error {
	m := &Client{}
	m.ID = id
	return state.Delete(m)
}

// FindAll grabs all of the existing clients
func (state *ClientsManager) FindAll() ([]Client, error) {
	clients := []Client{}
	err := state.db.Find(&clients).Error
	return clients, err
}

// Find tries to find an existing Client from a id
func (state *ClientsManager) Find(id uint) (*Client, error) {
	client := &Client{}
	err := state.db.Where("id = ?", id).Find(client).Error
	return client, err
}

// FindWithName tries to find an existing Client from a name
func (state *ClientsManager) FindWithName(name string) (*Client, error) {
	client := &Client{}
	err := state.db.Where("namn = ?", name).Find(&client).Error
	return client, err
}

// Save saves a client in the database
func (state *ClientsManager) Save(client *Client) error {
	return state.db.Save(&client).Error
}

package models

import (
	"github.com/jinzhu/gorm"

	// import mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Server struct
type Server struct {
	gorm.Model `json:"-"`
	IP         string `gorm:"not null" json:"ip"`
	Namn       string `gorm:"not null;unique" json:"namn"`
}

// ServersManager struct
type ServersManager struct {
	db *DB
}

// NewServersManager - Creates a new *ServersManager that can be used for managing servers.
func NewServersManager(db *DB) (*ServersManager, error) {
	db.AutoMigrate(&Server{})
	manager := ServersManager{}
	manager.db = db
	return &manager, nil
}

// Create inserts a new Server in the database
func (state *ServersManager) Create(server *Server) bool {
	state.db.Create(server)
	return state.db.NewRecord(server)
}

// FindAll grabs all of the existing Servers
func (state *ServersManager) FindAll() []Server {
	servers := []Server{}
	state.db.Find(&servers)
	return servers
}

// Find tries to find an existing Server from a id
func (state *ServersManager) Find(id uint) *Server {
	server := Server{}
	state.db.Where("id = ?", id).Find(&server)
	return &server
}

// Save saves a Server from a Server struct
func (state *ServersManager) Save(server *Server) {
	state.db.Save(&server)
}

// Delete removes an existing Server
// from the database.
func (state *ServersManager) Delete(server *Server) error {
	return state.db.Delete(&server).Error
}

// DeleteFromID is a wrapper for Delete that
// creates a new Server instance from ID and then run the Delete function
func (state *ServersManager) DeleteFromID(id uint) error {
	m := &Server{}
	m.ID = id
	return state.Delete(m)
}

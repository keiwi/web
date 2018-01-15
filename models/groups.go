package models

import (
	"time"
	// import mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Group struct
type Group struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	CommandID int        `gorm:"command_id" json:"command_id"`
	GroupName string     `gorm:"not null" json:"group_name"`
	NextCheck int        `json:"next_check"`
	StopError bool       `json:"stop_error"`
}

// GroupsManager struct
type GroupsManager struct {
	db *DB
}

// NewGroupsManager - Creates a new *GroupsManager that can be used for managing groups.
func NewGroupsManager(db *DB) (*GroupsManager, error) {
	db.AutoMigrate(&Group{})
	manager := GroupsManager{}
	manager.db = db
	return &manager, nil
}

// Create inserts a new group in the database
func (state *GroupsManager) Create(group *Group) error {
	return state.db.Create(group).Error
}

// ExistsName checks if a group with the name exists
func (state *GroupsManager) ExistsName(name string) bool {
	if err := state.db.Where("group_name=?", name).Find(&Group{}).Error; err != nil {
		return false
	}
	return true
}

// FindAll grabs all of the existing groups
func (state *GroupsManager) FindAll() ([]Group, error) {
	groups := []Group{}
	err := state.db.Find(&groups).Error
	return groups, err
}

// Find tries to find an existing Group from a id
func (state *GroupsManager) Find(id uint) (*Group, error) {
	group := Group{}
	err := state.db.Where("id = ?", id).Find(&group).Error
	return &group, err
}

// Save saves a Group from a Group struct
func (state *GroupsManager) Save(group *Group) error {
	return state.db.Save(&group).Error
}

// Delete removes an existing Group
// from the database.
func (state *GroupsManager) Delete(group *Group) error {
	return state.db.Delete(&group).Error
}

// DeleteWithID is a wrapper for Delete that
// creates a new group instance from ID and then run the Delete function
func (state *GroupsManager) DeleteWithID(id uint) error {
	m := &Group{}
	m.ID = id
	return state.Delete(m)
}

// DeleteWithName is a wrapper for Delete that
// creates a new group instance from ID and then run the Delete function
func (state *GroupsManager) DeleteWithName(name string) error {
	return state.db.Where("group_name = ?", name).Delete(&Group{GroupName: name}).Error
}

// UpdateName update all record names in the database
// with a new value
func (state *GroupsManager) UpdateName(oldname, newname string) (int64, error) {
	u := state.db.Where("name = ?", oldname).Updates(newname)
	return u.RowsAffected, u.Error
}

// HasGroup - Check if given group exists
func (state *GroupsManager) HasGroup(group string) bool {
	if err := state.db.Where("group_name=?", group).Find(&Group{}).Error; err != nil {
		return false
	}
	return true
}

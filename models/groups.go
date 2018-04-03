package models

import (
	"errors"
	"time"

	"github.com/keiwi/utils"
	"github.com/keiwi/utils/models"
	"github.com/nats-io/go-nats"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2/bson"
)

// GroupsManager struct
type GroupsManager struct {
	conn *nats.Conn
}

// NewGroupsManager - Creates a new *GroupsManager that can be used for managing groups.
func NewGroupsManager(conn *nats.Conn) (*GroupsManager, error) {
	manager := GroupsManager{}
	manager.conn = conn
	return &manager, nil
}

// Create inserts a new group in the database
func (state *GroupsManager) Create(group *models.Group) error {
	group.CreatedAt = time.Now()
	group.UpdatedAt = time.Now()

	data, err := bson.MarshalJSON(&group)
	if err != nil {
		return err
	}
	return state.conn.Publish("groups.create.send", data)
}

// FindAll grabs all of the existing groups
func (state *GroupsManager) FindAll() ([]models.Group, error) {
	requestData := utils.FindOptions{
		Sort: utils.Sort{"created_at"},
	}
	data, err := bson.MarshalJSON(&requestData)
	if err != nil {
		return nil, err
	}
	msg, err := state.conn.Request("groups.retrieve.find", data, time.Duration(viper.GetInt("nats_delay"))*time.Second)
	if err != nil {
		return nil, err
	}

	var groups []models.Group
	err = bson.UnmarshalJSON(msg.Data, &groups)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

// Find tries to find an existing Group from a id
func (state *GroupsManager) Find(id string) (*models.Group, error) {
	return state.FindFilter(utils.Filter{"_id": id})
}

// Find tries to find an existing Group from a id
func (state *GroupsManager) FindName(name string) (*models.Group, error) {
	return state.FindFilter(utils.Filter{"name": name})
}

// Find tries to find an existing Group from a id
func (state *GroupsManager) FindFilter(filter utils.Filter) (*models.Group, error) {
	requestData := utils.FindOptions{
		Filter: filter,
		Sort:   utils.Sort{"created_at"},
		Limit:  1,
	}
	data, err := bson.MarshalJSON(&requestData)
	if err != nil {
		return nil, err
	}

	msg, err := state.conn.Request("groups.retrieve.find", data, time.Duration(viper.GetInt("nats_delay"))*time.Second)
	if err != nil {
		return nil, err
	}

	groups := []models.Group{}
	err = bson.UnmarshalJSON(msg.Data, &groups)
	if err != nil {
		return nil, err
	}
	if len(groups) <= 0 {
		return nil, errors.New("could not find any groups")
	}
	return &groups[0], nil
}

// Save saves a Group from a Group struct
func (state *GroupsManager) Save(filter utils.Filter, updates utils.Updates) error {
	if u, ok := updates["$set"]; ok {
		up := u.(map[string]interface{})
		up["updated_at"] = time.Now()
		updates["$set"] = up
	} else {
		updates["$set"] = map[string]interface{}{"updated_at": time.Now()}
	}

	requestData := utils.UpdateOptions{
		Filter:  filter,
		Updates: updates,
	}

	data, err := bson.MarshalJSON(&requestData)
	if err != nil {
		return err
	}

	return state.conn.Publish("groups.update.send", data)
}

// Delete removes an existing Group
// from the database.
func (state *GroupsManager) Delete(group *models.Group) error {
	requestData := utils.DeleteOptions{
		Filter: utils.Filter{"_id": group.ID},
	}

	data, err := bson.MarshalJSON(&requestData)
	if err != nil {
		return err
	}

	return state.conn.Publish("groups.delete.send", data)
}

// DeleteWithID is a wrapper for Delete that
// creates a new group instance from ID and then run the Delete function
func (state *GroupsManager) DeleteWithID(id string) error {
	filter := utils.Filter{"commands.id": bson.ObjectIdHex(id)}
	updates := utils.Updates{
		"$pull": bson.M{
			"commands": bson.M{
				"id": bson.ObjectIdHex(id),
			},
		},
	}
	return state.Save(filter, updates)
}

// DeleteWithName is a wrapper for Delete that
// creates a new group instance from ID and then run the Delete function
func (state *GroupsManager) DeleteWithName(name string) error {
	requestData := utils.DeleteOptions{
		Filter: utils.Filter{"name": name},
	}

	data, err := bson.MarshalJSON(&requestData)
	if err != nil {
		return err
	}

	return state.conn.Publish("groups.delete.send", data)
}

// UpdateName update all record names in the database
// with a new value
func (state *GroupsManager) UpdateName(oldname, newname string) error {

	requestData := utils.UpdateOptions{
		Filter: utils.Filter{"name": oldname},
		Updates: utils.Updates{
			"name":       newname,
			"updated_at": time.Now(),
		},
	}

	data, err := bson.MarshalJSON(&requestData)
	if err != nil {
		return err
	}

	return state.conn.Publish("groups.update.send", data)
}

// ExistsName checks if a group with the name exists
func (state *GroupsManager) HasGroup(name string) bool {
	requestData := utils.HasOptions{
		Filter: utils.Filter{"name": name},
	}

	data, err := bson.MarshalJSON(&requestData)
	if err != nil {
		return false
	}

	msg, err := state.conn.Request("groups.has", data, time.Duration(viper.GetInt("nats_delay"))*time.Second)
	if err != nil {
		return false
	}

	var has bool
	err = bson.UnmarshalJSON(msg.Data, has)
	if err != nil {
		return false
	}

	return has
}

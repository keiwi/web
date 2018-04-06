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

// ServersManager struct
type ServersManager struct {
	conn *nats.Conn
}

// NewServersManager - Creates a new *ServersManager that can be used for managing servers.
func NewServersManager(conn *nats.Conn) (*ServersManager, error) {
	manager := ServersManager{}
	manager.conn = conn
	return &manager, nil
}

// Create inserts a new Server in the database
func (state *ServersManager) Create(server *models.Server) error {
	server.CreatedAt = time.Now()
	server.UpdatedAt = time.Now()

	data, err := bson.MarshalJSON(&server)
	if err != nil {
		return err
	}
	return state.conn.Publish("servers.create.send", data)
}

// FindAll grabs all of the existing Servers
func (state *ServersManager) FindAll() ([]models.Server, error) {
	requestData := utils.FindOptions{
		Sort: utils.Sort{"created_at"},
	}
	data, err := bson.MarshalJSON(&requestData)
	if err != nil {
		return nil, err
	}
	msg, err := state.conn.Request("servers.retrieve.find", data, time.Duration(viper.GetInt("nats.delay"))*time.Second)
	if err != nil {
		return nil, err
	}

	var servers []models.Server
	err = bson.UnmarshalJSON(msg.Data, &servers)
	if err != nil {
		return nil, err
	}
	return servers, nil
}

// Find tries to find an existing Server from a id
func (state *ServersManager) Find(id uint) (*models.Server, error) {
	requestData := utils.FindOptions{
		Filter: utils.Filter{"_id": id},
		Sort:   utils.Sort{"created_at"},
		Limit:  1,
	}
	data, err := bson.MarshalJSON(&requestData)
	if err != nil {
		return nil, err
	}

	msg, err := state.conn.Request("servers.retrieve.find", data, time.Duration(viper.GetInt("nats.delay"))*time.Second)
	if err != nil {
		return nil, err
	}

	servers := []models.Server{}
	err = bson.UnmarshalJSON(msg.Data, &servers)
	if err != nil {
		return nil, err
	}
	if len(servers) <= 0 {
		return nil, errors.New("could not find any servers")
	}
	return &servers[0], nil
}

// Save saves a Server from a Server struct
func (state *ServersManager) Save(id string, updates utils.Updates) error {
	updates["updated_at"] = time.Now()

	requestData := utils.UpdateOptions{
		Filter:  utils.Filter{"_id": id},
		Updates: updates,
	}

	data, err := bson.MarshalJSON(&requestData)
	if err != nil {
		return err
	}

	return state.conn.Publish("servers.update.send", data)
}

// Delete removes an existing Server
// from the database.
func (state *ServersManager) Delete(server *models.Server) error {
	requestData := utils.DeleteOptions{
		Filter: utils.Filter{"_id": server.ID},
	}

	data, err := bson.MarshalJSON(&requestData)
	if err != nil {
		return err
	}

	return state.conn.Publish("servers.delete.send", data)
}

// DeleteFromID is a wrapper for Delete that
// creates a new Server instance from ID and then run the Delete function
func (state *ServersManager) DeleteWithID(id string) error {
	m := &models.Server{}
	m.ID = bson.ObjectId(id)
	return state.Delete(m)
}

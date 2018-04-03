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

// ClientsManager struct
type ClientsManager struct {
	conn *nats.Conn
}

// NewClientsManager - Creates a new *ClientsManager that can be used for managing clients.
func NewClientsManager(conn *nats.Conn) (*ClientsManager, error) {
	manager := ClientsManager{}
	manager.conn = conn

	return &manager, nil
}

// Create inserts a new client into the database
func (state *ClientsManager) Create(client *models.Client) error {
	client.CreatedAt = time.Now()
	client.UpdatedAt = time.Now()

	data, err := bson.MarshalJSON(client)
	if err != nil {
		return err
	}
	return state.conn.Publish("clients.create.send", data)
}

// Delete removes an existing Client
// from the database.
func (state *ClientsManager) Delete(client *models.Client) error {
	requestData := utils.DeleteOptions{
		Filter: utils.Filter{"_id": client.ID},
	}

	data, err := bson.MarshalJSON(requestData)
	if err != nil {
		return err
	}

	return state.conn.Publish("clients.delete.send", data)
}

// DeleteWithID is a wrapper for Delete that
// creates a new client instance from ID and then run the Delete function
func (state *ClientsManager) DeleteWithID(id string) error {
	m := &models.Client{}
	m.ID = bson.ObjectIdHex(id)
	return state.Delete(m)
}

// FindAll grabs all of the existing clients
func (state *ClientsManager) FindAll() ([]models.Client, error) {
	requestData := utils.FindOptions{
		Sort: utils.Sort{"created_at"},
	}
	data, err := bson.MarshalJSON(requestData)
	if err != nil {
		return nil, err
	}
	msg, err := state.conn.Request("clients.retrieve.find", data, time.Duration(viper.GetInt("nats_delay"))*time.Second)
	if err != nil {
		return nil, err
	}

	var clients []models.Client
	err = bson.UnmarshalJSON(msg.Data, &clients)
	if err != nil {
		return nil, err
	}
	return clients, nil
}

func (state *ClientsManager) FindFilter(filter utils.Filter) (*models.Client, error) {
	requestData := utils.FindOptions{
		Filter: filter,
		Sort:   utils.Sort{"created_at"},
		Limit:  1,
	}
	data, err := bson.MarshalJSON(requestData)
	if err != nil {
		return nil, err
	}

	msg, err := state.conn.Request("clients.retrieve.find", data, time.Duration(viper.GetInt("nats_delay"))*time.Second)
	if err != nil {
		return nil, err
	}

	clients := []models.Client{}
	err = bson.UnmarshalJSON(msg.Data, &clients)
	if err != nil {
		return nil, err
	}
	if len(clients) <= 0 {
		return nil, errors.New("could not find any clients")
	}
	return &clients[0], nil
}

// Find tries to find an existing Client from a id
func (state *ClientsManager) Find(id string) (*models.Client, error) {
	return state.FindFilter(utils.Filter{"_id": bson.ObjectIdHex(id)})
}

// FindWithName tries to find an existing Client from a name
func (state *ClientsManager) FindWithName(name string) (*models.Client, error) {
	return state.FindFilter(utils.Filter{"name": name})
}

// Save saves a client in the database
func (state *ClientsManager) Save(id string, updates utils.Updates) error {
	if u, ok := updates["$set"]; ok {
		up := u.(bson.M)
		up["updated_at"] = time.Now()
		updates["$set"] = up
	} else {
		updates["$set"] = map[string]interface{}{"updated_at": time.Now()}
	}

	requestData := utils.UpdateOptions{
		Filter:  utils.Filter{"_id": bson.ObjectIdHex(id)},
		Updates: updates,
	}

	data, err := bson.MarshalJSON(requestData)
	if err != nil {
		return err
	}

	return state.conn.Publish("clients.update.send", data)
}

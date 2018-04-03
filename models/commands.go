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

// CommandsManager struct
type CommandsManager struct {
	conn *nats.Conn
}

// NewCommandsManager - Creates a new *CommandsManager that can be used for managing commands.
func NewCommandsManager(conn *nats.Conn) (*CommandsManager, error) {
	manager := CommandsManager{}
	manager.conn = conn
	return &manager, nil
}

// Create inserts a new command in the database
func (state *CommandsManager) Create(cmd *models.Command) error {
	cmd.CreatedAt = time.Now()
	cmd.UpdatedAt = time.Now()

	data, err := bson.MarshalJSON(&cmd)
	if err != nil {
		return err
	}
	return state.conn.Publish("commands.create.send", data)
}

func (state *CommandsManager) FindFilter(filter utils.Filter) (*models.Command, error) {
	requestData := utils.FindOptions{
		Filter: filter,
		Sort:   utils.Sort{"created_at"},
		Limit:  1,
	}
	data, err := bson.MarshalJSON(&requestData)
	if err != nil {
		return nil, err
	}

	msg, err := state.conn.Request("commands.retrieve.find", data, time.Duration(viper.GetInt("nats_delay"))*time.Second)
	if err != nil {
		return nil, err
	}

	commands := []models.Command{}
	err = bson.UnmarshalJSON(msg.Data, &commands)
	if err != nil {
		return nil, err
	}
	if len(commands) <= 0 {
		return nil, errors.New("could not find any commands")
	}
	return &commands[0], nil
}

// FindAll grabs all of the existing commands, note that this
// is base commands not GroupCommands
func (state *CommandsManager) FindAll() ([]models.Command, error) {
	requestData := utils.FindOptions{
		Sort: utils.Sort{"created_at"},
	}
	data, err := bson.MarshalJSON(&requestData)
	if err != nil {
		return nil, err
	}
	msg, err := state.conn.Request("commands.retrieve.find", data, time.Duration(viper.GetInt("nats_delay"))*time.Second)
	if err != nil {
		return nil, err
	}

	var commands []models.Command
	err = bson.UnmarshalJSON(msg.Data, &commands)
	if err != nil {
		return nil, err
	}
	return commands, nil
}

// Find tries to find an existing command from a id
func (state *CommandsManager) Find(id string) (*models.Command, error) {
	return state.FindFilter(utils.Filter{"_id": bson.ObjectIdHex(id)})
}

// Save saves a command from a command struct
func (state *CommandsManager) Save(id string, updates utils.Updates) error {
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

	return state.conn.Publish("commands.update.send", data)
}

// Delete removes an existing command
// from the database.
func (state *CommandsManager) Delete(cmd *models.Command) error {
	requestData := utils.DeleteOptions{
		Filter: utils.Filter{"_id": cmd.ID},
	}

	data, err := bson.MarshalJSON(requestData)
	if err != nil {
		return err
	}

	return state.conn.Publish("commands.delete.send", data)
}

// DeleteWithID is a wrapper for Delete that
// creates a new CMD instance from ID and then run the Delete function
func (state *CommandsManager) DeleteWithID(id string) error {
	m := &models.Command{}
	m.ID = bson.ObjectIdHex(id)
	return state.Delete(m)
}

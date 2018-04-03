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

// ChecksManager struct
type ChecksManager struct {
	conn *nats.Conn
}

// NewChecksManager - Creates a new *ChecksManager that can be used for managing checks.
func NewChecksManager(conn *nats.Conn) (*ChecksManager, error) {
	manager := ChecksManager{}
	manager.conn = conn
	return &manager, nil
}

// GetChecksBetweenDate grabs all checks with a specific command id
// between 2 different dates.
func (state *ChecksManager) GetChecksBetweenDate(from, to time.Time, commandID string) ([]models.Check, error) {
	requestData := utils.FindOptions{
		Filter: utils.Filter{"command_id": bson.ObjectIdHex(commandID), "created_at": bson.M{"$gte": from, "$lte": to}},
		Sort:   utils.Sort{"created_at"},
	}
	data, err := bson.MarshalJSON(requestData)
	if err != nil {
		return nil, err
	}

	msg, err := state.conn.Request("checks.retrieve.find", data, time.Duration(viper.GetInt("nats_delay"))*time.Second)
	if err != nil {
		return nil, err
	}

	var checks []models.Check
	err = bson.UnmarshalJSON(msg.Data, &checks)
	if err != nil {
		return nil, err
	}
	return checks, nil
}

// GetChecksBetweenDateClient is similiar to GetChecksBetweenDate
// where it grabs all checks between 2 different dates with a specific
// command id, the difference is that this function also grabs checks
// with a specific command_id
func (state *ChecksManager) GetChecksBetweenDateClient(from, to time.Time, commandID, clientID string, max int) ([]models.Check, error) {
	requestData := utils.FindOptions{
		Filter: utils.Filter{"command_id": bson.ObjectIdHex(commandID), "client_id": bson.ObjectIdHex(clientID), "created_at": bson.M{"$gte": from, "$lte": to}},
		Sort:   utils.Sort{"created_at"},
		Max:    utils.Max(max),
	}
	data, err := bson.MarshalJSON(requestData)
	if err != nil {
		return nil, err
	}

	msg, err := state.conn.Request("checks.retrieve.find", data, time.Duration(viper.GetInt("nats_delay"))*time.Second)
	if err != nil {
		return nil, err
	}

	var checks []models.Check
	err = bson.UnmarshalJSON(msg.Data, &checks)
	if err != nil {
		return nil, err
	}
	return checks, nil
}

// Create inserts a new Check in the database
func (state *ChecksManager) Create(check *models.Check) error {
	check.CreatedAt = time.Now()
	check.UpdatedAt = time.Now()

	data, err := bson.MarshalJSON(check)
	if err != nil {
		return err
	}
	return state.conn.Publish("checks.create.send", data)
}

// FindAll grabs all of the existing Checks
func (state *ChecksManager) FindAll() ([]models.Check, error) {
	requestData := utils.FindOptions{
		Sort: utils.Sort{"created_at"},
	}
	data, err := bson.MarshalJSON(requestData)
	if err != nil {
		return nil, err
	}
	msg, err := state.conn.Request("checks.retrieve.find", data, time.Duration(viper.GetInt("nats_delay"))*time.Second)
	if err != nil {
		return nil, err
	}

	var checks []models.Check
	err = bson.UnmarshalJSON(msg.Data, &checks)
	if err != nil {
		return nil, err
	}
	return checks, nil
}

// Find tries to find an existing Check from a id
func (state *ChecksManager) Find(id string) (*models.Check, error) {
	requestData := utils.FindOptions{
		Filter: utils.Filter{"_id": bson.ObjectIdHex(id)},
		Sort:   utils.Sort{"created_at"},
		Limit:  1,
	}
	data, err := bson.MarshalJSON(requestData)
	if err != nil {
		return nil, err
	}

	msg, err := state.conn.Request("checks.retrieve.find", data, time.Duration(viper.GetInt("nats_delay"))*time.Second)
	if err != nil {
		return nil, err
	}

	checks := []models.Check{}
	err = bson.UnmarshalJSON(msg.Data, &checks)
	if err != nil {
		return nil, err
	}
	if len(checks) <= 0 {
		return nil, errors.New("could not find any checks")
	}
	return &checks[0], nil
}

// FindWithClientIDAndCommandID tries to find an existing Check from a client ID and command ID
func (state *ChecksManager) FindWithClientIDAndCommandID(clientID, commandID string) ([]models.Check, error) {
	requestData := utils.FindOptions{
		Filter: utils.Filter{"command_id": bson.ObjectIdHex(commandID), "client_id": bson.ObjectIdHex(clientID)},
		Sort:   utils.Sort{"-created_at"},
		Limit:  1,
	}
	data, err := bson.MarshalJSON(requestData)
	if err != nil {
		return nil, err
	}
	msg, err := state.conn.Request("checks.retrieve.find", data, time.Duration(viper.GetInt("nats_delay"))*time.Second)
	if err != nil {
		return nil, err
	}

	var checks []models.Check
	err = bson.UnmarshalJSON(msg.Data, &checks)
	if err != nil {
		return nil, err
	}
	return checks, nil
}

// Save saves a Check from a Check struct
func (state *ChecksManager) Save(id string, updates utils.Updates) error {
	updates["updated_at"] = time.Now()

	requestData := utils.UpdateOptions{
		Filter:  utils.Filter{"_id": bson.ObjectIdHex(id)},
		Updates: updates,
	}

	data, err := bson.MarshalJSON(requestData)
	if err != nil {
		return err
	}

	return state.conn.Publish("checks.update.send", data)
}

// Delete removes an existing Check
// from the database.
func (state *ChecksManager) Delete(check *models.Check) error {
	requestData := utils.DeleteOptions{
		Filter: utils.Filter{"_id": check.ID},
	}

	data, err := bson.MarshalJSON(requestData)
	if err != nil {
		return err
	}

	return state.conn.Publish("checks.delete.send", data)
}

// DeleteWithID is a wrapper for Delete that
// creates a new Check instance from ID and then run the Delete function
func (state *ChecksManager) DeleteWithID(id string) error {
	m := &models.Check{}
	m.ID = bson.ObjectId(id)
	return state.Delete(m)
}

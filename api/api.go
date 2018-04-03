package api

import (
	"encoding/json"
	"errors"
	"io"
	"strconv"

	"github.com/keiwi/utils"
	"github.com/keiwi/web/models"
)

// API struct
type API struct {
	handler *models.Handler
}

// NewAPI - Creates a new API
func NewAPI(handler *models.Handler) *API {
	return &API{handler: handler}
}

// MessageJSON - json data for outputting
type MessageJSON struct {
	Success bool        `json:"success"` // Wether an error occured or not
	Message string      `json:"message"` // The message
	Data    interface{} `json:"data"`    // Extra data, generally it will contain a struct
}

// EditJSON is the values you can send
// in the POST request body to modify an existing
// database record.
type EditJSON struct {
	ID     string      `json:"id"`
	Option string      `json:"option"`
	Value  interface{} `json:"value"`
}

// IDOptions is the expected data when dealing with only ID
// for example when deleting or trying to find a specific
// data in the database
type IDOptions struct {
	ID string `json:"id"`
}

func outputJSON(w io.Writer, success bool, message string, data interface{}) {
	out, _ := json.Marshal(MessageJSON{Success: success, Message: message, Data: data})
	if _, err := w.Write(out); err != nil {
		utils.Log.Fatal(err.Error())
	}
}

func convertToInt(i interface{}) (int64, error) {
	switch ci := i.(type) {
	case int64:
		return ci, nil
	case float64:
		return int64(ci), nil
	case string:
		pi, err := strconv.ParseInt(ci, 10, 64)
		if err != nil {
			return 0, err
		}
		return pi, nil
	default:
		return 0, errors.New("invalid number type")
	}
}

// Request is the struct that will be used when sending a request to a Keiwi server
type Request struct {
	Type      string `json:"type"`
	Command   string `json:"command"`
	GroupName string `json:"group_name"`
	ID        string `json:"id"`
	CommandID string `json:"command_id"`
	GroupID   string `json:"group_id"`
	Save      bool   `json:"save"`
}

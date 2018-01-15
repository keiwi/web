package api

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"strconv"

	"github.com/keiwi/utils"
	"github.com/keiwi/web/models"
)

// API struct
type API struct {
	users         *models.UserManager
	groups        *models.GroupsManager
	alerts        *models.AlertsManager
	alertsOptions *models.AlertsOptionsManager
	checks        *models.ChecksManager
	commands      *models.CommandsManager
	clients       *models.ClientsManager
	servers       *models.ServersManager
	uploads       *models.UploadsManager
}

// NewAPI - Creates a new API
func NewAPI(db *models.DB) *API {
	utils.Log.Debug("Creating internal API")

	utils.Log.Debug("Initialzing user manager")
	usermgr, err := models.NewUserManager(db)
	if err != nil {
		utils.Log.WithField("error", err).Fatal("error when initializing user manager")
	}
	utils.Log.Debug("User manager initialization done")

	utils.Log.Debug("Initialzing alerts manager")
	alertsmgr, err := models.NewAlertsManager(db)
	if err != nil {
		utils.Log.WithField("error", err).Fatal("error when initializing alerts manager")
	}
	utils.Log.Debug("Alerts manager initialization done")

	utils.Log.Debug("Initialzing alert options manager")
	alertsoptsmgr, err := models.NewAlertsOptionsManager(db)
	if err != nil {
		utils.Log.WithField("error", err).Fatal("error when initializing alert options manager")
	}
	utils.Log.Debug("Alert options manager initialization done")

	utils.Log.Debug("Initialzing checks manager")
	checksmgr, err := models.NewChecksManager(db)
	if err != nil {
		utils.Log.WithField("error", err).Fatal("error when initializing checks manager")
	}
	utils.Log.Debug("Checks manager initialization done")

	utils.Log.Debug("Initialzing clients manager")
	clientsmgr, err := models.NewClientsManager(db)
	if err != nil {
		utils.Log.WithField("error", err).Fatal("error when initializing clients manager")
	}
	utils.Log.Debug("Clients manager initialization done")

	utils.Log.Debug("Initialzing commands manager")
	commandsmgr, err := models.NewCommandsManager(db)
	if err != nil {
		utils.Log.WithField("error", err).Fatal("error when initializing commands manager")
	}
	utils.Log.Debug("Commands manager initialization done")

	utils.Log.Debug("Initialzing groups manager")
	groupsmgr, err := models.NewGroupsManager(db)
	if err != nil {
		utils.Log.WithField("error", err).Fatal("error when initializing groups manager")
	}
	utils.Log.Debug("Groups manager initialization done")

	utils.Log.Debug("Initialzing servers manager")
	serversmgr, err := models.NewServersManager(db)
	if err != nil {
		utils.Log.WithField("error", err).Fatal("error when initializing servers manager")
	}
	utils.Log.Debug("Servers manager initialization done")

	utils.Log.Debug("Initialzing upload manager")
	uploadsmgr, err := models.NewUploadsManager(db)
	if err != nil {
		utils.Log.WithField("error", err).Fatal("error when initializing upload manager")
	}
	utils.Log.Debug("Upload manager initialization done")

	utils.Log.Debug("API initialization done")
	return &API{
		users:         usermgr,
		groups:        groupsmgr,
		alerts:        alertsmgr,
		alertsOptions: alertsoptsmgr,
		checks:        checksmgr,
		commands:      commandsmgr,
		clients:       clientsmgr,
		servers:       serversmgr,
		uploads:       uploadsmgr,
	}
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
	ID     uint        `json:"id"`
	Option string      `json:"option"`
	Value  interface{} `json:"value"`
}

// IDOptions is the expected data when dealing with only ID
// for example when deleting or trying to find a specific
// data in the database
type IDOptions struct {
	ID uint `json:"id"`
}

func outputJSON(w io.Writer, success bool, message string, data interface{}) {
	out, _ := json.Marshal(MessageJSON{Success: success, Message: message, Data: data})
	if _, err := w.Write(out); err != nil {
		log.Fatal(err)
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
		return 0, errors.New("Invalid number type")
	}
}

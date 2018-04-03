package models

import (
	"github.com/nats-io/go-nats"
	"github.com/pkg/errors"

	"github.com/keiwi/utils"
)

// Handler handles all of the tasks related to models
type Handler struct {
	Conn          *nats.Conn
	Users         *UserManager
	Groups        *GroupsManager
	Alerts        *AlertsManager
	AlertsOptions *AlertsOptionsManager
	Checks        *ChecksManager
	Commands      *CommandsManager
	Clients       *ClientsManager
	Servers       *ServersManager
	Uploads       *UploadsManager
}

// NewHandler - will connect to NATS and return a Handler
func NewHandler(url string) (*Handler, error) {
	utils.Log.WithField("URL", url).Debug("connecting to NATS")
	conn, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	utils.Log.Info("successful connection to NATS")

	utils.Log.Debug("creating all of the managers")
	utils.Log.Debug("initialzing user manager")
	usermgr, err := NewUserManager(conn)
	if err != nil {
		return nil, errors.Wrap(err, "error when initializing user manager")
	}
	utils.Log.Debug("user manager initialization done")

	utils.Log.Debug("initialzing alerts manager")
	alertsmgr, err := NewAlertsManager(conn)
	if err != nil {
		return nil, errors.Wrap(err, "error when initializing alerts manager")
	}
	utils.Log.Debug("alerts manager initialization done")

	utils.Log.Debug("initialzing alert options manager")
	alertsoptsmgr, err := NewAlertsOptionsManager(conn)
	if err != nil {
		return nil, errors.Wrap(err, "error when initializing alert options manager")
	}
	utils.Log.Debug("alert options manager initialization done")

	utils.Log.Debug("initialzing checks manager")
	checksmgr, err := NewChecksManager(conn)
	if err != nil {
		return nil, errors.Wrap(err, "error when initializing checks manager")
	}
	utils.Log.Debug("checks manager initialization done")

	utils.Log.Debug("initialzing clients manager")
	clientsmgr, err := NewClientsManager(conn)
	if err != nil {
		return nil, errors.Wrap(err, "error when initializing clients manager")
	}
	utils.Log.Debug("clients manager initialization done")

	utils.Log.Debug("initialzing commands manager")
	commandsmgr, err := NewCommandsManager(conn)
	if err != nil {
		return nil, errors.Wrap(err, "error when initializing commands manager")
	}
	utils.Log.Debug("commands manager initialization done")

	utils.Log.Debug("initialzing groups manager")
	groupsmgr, err := NewGroupsManager(conn)
	if err != nil {
		return nil, errors.Wrap(err, "error when initializing groups manager")
	}
	utils.Log.Debug("groups manager initialization done")

	utils.Log.Debug("initialzing servers manager")
	serversmgr, err := NewServersManager(conn)
	if err != nil {
		return nil, errors.Wrap(err, "error when initializing servers manager")
	}
	utils.Log.Debug("servers manager initialization done")

	utils.Log.Debug("initialzing upload manager")
	uploadsmgr, err := NewUploadsManager(conn)
	if err != nil {
		return nil, errors.Wrap(err, "error when initializing upload manager")
	}
	utils.Log.Debug("upload manager initialization done")

	utils.Log.Debug("Handler initialization done")
	return &Handler{
		Conn:          conn,
		Users:         usermgr,
		Groups:        groupsmgr,
		Alerts:        alertsmgr,
		AlertsOptions: alertsoptsmgr,
		Checks:        checksmgr,
		Commands:      commandsmgr,
		Clients:       clientsmgr,
		Servers:       serversmgr,
		Uploads:       uploadsmgr,
	}, nil
}

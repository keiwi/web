package models

import (
	"github.com/keiwi/utils/log"
	"github.com/nats-io/go-nats"
	"github.com/pkg/errors"
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
	log.WithField("URL", url).Debug("connecting to NATS")
	conn, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	log.Info("successful connection to NATS")

	log.Debug("creating all of the managers")
	log.Debug("initialzing user manager")
	usermgr, err := NewUserManager(conn)
	if err != nil {
		return nil, errors.Wrap(err, "error when initializing user manager")
	}
	log.Debug("user manager initialization done")

	log.Debug("initialzing alerts manager")
	alertsmgr, err := NewAlertsManager(conn)
	if err != nil {
		return nil, errors.Wrap(err, "error when initializing alerts manager")
	}
	log.Debug("alerts manager initialization done")

	log.Debug("initialzing alert options manager")
	alertsoptsmgr, err := NewAlertsOptionsManager(conn)
	if err != nil {
		return nil, errors.Wrap(err, "error when initializing alert options manager")
	}
	log.Debug("alert options manager initialization done")

	log.Debug("initialzing checks manager")
	checksmgr, err := NewChecksManager(conn)
	if err != nil {
		return nil, errors.Wrap(err, "error when initializing checks manager")
	}
	log.Debug("checks manager initialization done")

	log.Debug("initialzing clients manager")
	clientsmgr, err := NewClientsManager(conn)
	if err != nil {
		return nil, errors.Wrap(err, "error when initializing clients manager")
	}
	log.Debug("clients manager initialization done")

	log.Debug("initialzing commands manager")
	commandsmgr, err := NewCommandsManager(conn)
	if err != nil {
		return nil, errors.Wrap(err, "error when initializing commands manager")
	}
	log.Debug("commands manager initialization done")

	log.Debug("initialzing groups manager")
	groupsmgr, err := NewGroupsManager(conn)
	if err != nil {
		return nil, errors.Wrap(err, "error when initializing groups manager")
	}
	log.Debug("groups manager initialization done")

	log.Debug("initialzing servers manager")
	serversmgr, err := NewServersManager(conn)
	if err != nil {
		return nil, errors.Wrap(err, "error when initializing servers manager")
	}
	log.Debug("servers manager initialization done")

	log.Debug("initialzing upload manager")
	uploadsmgr, err := NewUploadsManager(conn)
	if err != nil {
		return nil, errors.Wrap(err, "error when initializing upload manager")
	}
	log.Debug("upload manager initialization done")

	log.Debug("Handler initialization done")
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

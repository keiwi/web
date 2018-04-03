package models

import (
	"github.com/nats-io/go-nats"
)

// AlertsOptionsManager struct
type AlertsOptionsManager struct {
	conn *nats.Conn
}

// NewAlertsOptionsManager - Creates a new *AlertsOptionsManager that can be used for managing alertsOptions.
func NewAlertsOptionsManager(conn *nats.Conn) (*AlertsOptionsManager, error) {
	manager := AlertsOptionsManager{}
	manager.conn = conn
	return &manager, nil
}

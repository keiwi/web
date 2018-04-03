package models

import (
	"github.com/nats-io/go-nats"
)

// AlertsManager struct
type AlertsManager struct {
	conn *nats.Conn
}

// NewAlertsManager - Creates a new *AlertsManager that can be used for managing Alerts.
func NewAlertsManager(conn *nats.Conn) (*AlertsManager, error) {
	manager := AlertsManager{}
	manager.conn = conn
	return &manager, nil
}

package models

import (
	"github.com/nats-io/go-nats"
)

// UploadsManager struct
type UploadsManager struct {
	conn *nats.Conn
}

// NewUploadsManager - Creates a new *UploadsManager that can be used for managing uploads.
func NewUploadsManager(conn *nats.Conn) (*UploadsManager, error) {
	manager := UploadsManager{}
	manager.conn = conn
	return &manager, nil
}

package events

import (
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Checksum string
	HostsID  int
	Message  string
	StatusID int
}

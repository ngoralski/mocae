package events_status

import (
	"gorm.io/gorm"
)

type EventStatus struct {
	gorm.Model
	Checksum  string `gorm:"index"`
	CommentID int
	StatusID  int
}

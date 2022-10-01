package logs

import (
	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	Log     string
	UserID  int
	EventID int
}

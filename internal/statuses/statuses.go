package statuses

import (
	"gorm.io/gorm"
)

type Status struct {
	gorm.Model
	Status string
}

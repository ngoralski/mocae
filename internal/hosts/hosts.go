package hosts

import "gorm.io/gorm"

type Host struct {
	gorm.Model
	Hostname string
	Comments string
}

package comments

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Comment string
	UserID  int
}

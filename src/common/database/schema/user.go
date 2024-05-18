package schema

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID string `gorm:"not null"`
	// Money int
	// Name string
}

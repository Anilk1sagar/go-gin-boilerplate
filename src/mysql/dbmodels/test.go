package dbmodels

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // _ mysql driver
)

// Test model
type Test struct {
	gorm.Model
	Name  string `gorm:"unique" json:"name"`
	Email string `json:"email"`
}

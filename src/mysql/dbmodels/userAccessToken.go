package dbmodels

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // _ mysql driver
)

// UserAccessToken model
type UserAccessToken struct {
	gorm.Model
	UserID    uint   `gorm:"unique" json:"userId"`
	Token     string `json:"token"`
	IsExpired bool   `json:"isExpired"`
}

package dbmodels

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {

	fmt.Println("Migration runs")

	// Creating tables
	db.AutoMigrate(&Test{}, &User{})
	return db
}

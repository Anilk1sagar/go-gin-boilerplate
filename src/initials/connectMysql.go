package initials

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // _ For mysql driver
	"gitlab.com/anilk1sagar/go_gin_test/src/mysql"
	"gitlab.com/anilk1sagar/go_gin_test/src/mysql/dbmodels"
	"gitlab.com/anilk1sagar/go_gin_test/src/utils"
)

// ConnectMysqlDb for connection
func (a *App) ConnectMysqlDb(mysql *mysql.MysqlConfig) {

	fmt.Println("**==== Connecting Mysqldb ====**")

	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
		mysql.MysqlDb.Username,
		mysql.MysqlDb.Password,
		mysql.MysqlDb.Name,
		mysql.MysqlDb.Charset)

	db, err := gorm.Open(mysql.MysqlDb.Dialect, dbURI)

	if err != nil {
		utils.Logger().Fatal("Could not connect database", err.Error())
	} else {
		fmt.Println("Database connect successfully")
	}

	// Migrate db and assign
	a.MysqlDb = dbmodels.DBMigrate(db)

}

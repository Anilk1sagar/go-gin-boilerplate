package main

import (
	"gitlab.com/anilk1sagar/go_gin_test/src/initials"
	"gitlab.com/anilk1sagar/go_gin_test/src/mysql"
	"gitlab.com/anilk1sagar/go_gin_test/src/utils"
)

func main() {

	//Environment Setup
	utils.InitializeEnvSetup()

	// Get Mysql Config
	mysqlConfig := mysql.GetMysqlConfig()

	//Initialize App
	app := &initials.App{}
	app.InitializeApp(mysqlConfig)

	/* Run Server */
	host := "localhost:" + utils.APIPort()
	app.RunServer(host)
}

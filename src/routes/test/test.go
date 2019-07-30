package test

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	SqlHandler "gitlab.com/anilk1sagar/go_gin_test/src/mysql/handler"
	"gitlab.com/anilk1sagar/go_gin_test/src/utils"
)

// MysqlDb for test routes
var MysqlDb *gorm.DB

// InitializeMysqlTest Routes
func InitializeMysqlTest(mysqlDb *gorm.DB) {
	MysqlDb = mysqlDb
}

// GetMysqlTest for single test
func GetMysqlTest(c *gin.Context) {

	SqlHandler.GetSqlTest(MysqlDb, c)
}

// GetAllMysqlTests for all test
func GetAllMysqlTests(c *gin.Context) {

	SqlHandler.GetAllSqlTests(MysqlDb, c)
}

// AddMysqlTest for adding new test
func AddMysqlTest(c *gin.Context) {

	checkBody := []string{"name", "email"}

	err := utils.ValidateRequestBody(c, checkBody)
	if err != nil {
		utils.Logger().Errorln(err.Error())
		utils.RespondError(c, http.StatusNotFound, err.Error())
		return
	}

	fmt.Println("aftre...........")

	// if c.PostForm("name") == "" {
	// 	utils.Logger().Errorln("name not found!")
	// 	utils.RespondError(c, http.StatusNotFound, "Required name")
	// 	return
	// }

	// if c.PostForm("email") == "" {
	// 	utils.Logger().Errorln("email not found!")
	// 	utils.RespondError(c, http.StatusNotFound, "Required email")
	// 	return
	// }

	// var name, email string

	// name = c.PostForm("name")
	// email = c.PostForm("email")

	// /* Creating test model */
	// dbModel := dbmodels.Test{}
	// dbModel.Name = name
	// dbModel.Email = email

	/* Calling Handler For Adding*/
	// SqlHandler.AddSqlTest(MysqlDb, c, dbModel)
	utils.RespondJSON(c, http.StatusOK, "okay")

}

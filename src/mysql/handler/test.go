package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"gitlab.com/anilk1sagar/go_gin_test/src/mysql/dbmodels"
	utils "gitlab.com/anilk1sagar/go_gin_test/src/utils"
)

// GetAllSqlTests Handler
func GetAllSqlTests(db *gorm.DB, c *gin.Context) {

	tests := []dbmodels.Test{}
	db.Find(&tests)

	utils.RespondJSON(c, http.StatusOK, tests)

}

// GetSqlTest Handler
func GetSqlTest(db *gorm.DB, c *gin.Context) {

	name := c.Param("name")

	test := getSqlTestOr404(db, name, c)

	if test == nil {
		return
	}

	utils.RespondJSON(c, http.StatusOK, test)

}

// AddSqlTest Handler
func AddSqlTest(db *gorm.DB, c *gin.Context, pModel dbmodels.Test) {

	/* Add to database */
	err := db.Save(&pModel).Error

	if err != nil {

		utils.Logger().Errorln("Handler Save Error: ", err.Error())
		utils.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	log.Println("Test Added Successfully: ", pModel)

	utils.RespondJSON(c, http.StatusCreated, pModel)

}

// getProjectOr404 gets a project instance if exists, or respond the 404 error otherwise
func getSqlTestOr404(db *gorm.DB, name string, c *gin.Context) *dbmodels.Test {

	test := dbmodels.Test{}

	if err := db.First(&test, dbmodels.Test{Name: name}).Error; err != nil {

		utils.Logger().Errorln(err.Error())
		utils.RespondError(c, http.StatusNotFound, err.Error())
		return nil
	}

	return &test

}

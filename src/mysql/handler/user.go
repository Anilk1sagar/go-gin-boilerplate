package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	dbmodels "gitlab.com/anilk1sagar/go_gin_test/src/mysql/dbmodels"
	utils "gitlab.com/anilk1sagar/go_gin_test/src/utils"
)

// Sqldb for User Handler
var Sqldb *gorm.DB

// InitializeUserHandler Routes
func InitializeUserHandler(mysqlDb *gorm.DB) {
	Sqldb = mysqlDb
}

/*
 */
// GetAllUsers Sql Handler
func GetAllUsers() []dbmodels.User {

	users := []dbmodels.User{}
	Sqldb.Find(&users)

	return users

}

/*
 */
// GetUserByEmail Sql Handler
func GetUserByEmail(email string) (dbmodels.User, error) {

	// Getting Params
	// vars := mux.Vars(r)
	// email := vars["email"]

	user := dbmodels.User{}

	// Get User
	err := Sqldb.First(&user, dbmodels.User{Email: email}).Error

	if err != nil {
		return user, err
	}

	if &user == nil {
		return user, nil
	}

	return user, nil

}

/*
 */
// AddUser Sql Handler
func AddUser(c *gin.Context, pModel dbmodels.User) {

	/* Add to database */
	err := Sqldb.Save(&pModel).Error

	if err != nil {

		utils.Logger().Errorln("Handler Save Error: ", err.Error())
		utils.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	pModel.Password = ""

	fmt.Println("\nUser Added Successfully: ", pModel)

	utils.RespondJSON(c, http.StatusCreated, pModel)

}

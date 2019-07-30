package user

import (
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	models "gitlab.com/anilk1sagar/go_gin_test/src/models"
	SqlModels "gitlab.com/anilk1sagar/go_gin_test/src/mysql/dbmodels"
	SqlHandler "gitlab.com/anilk1sagar/go_gin_test/src/mysql/handler"
	utils "gitlab.com/anilk1sagar/go_gin_test/src/utils"
	"golang.org/x/crypto/bcrypt"
)

// InitializeUser Routes
func InitializeUser(mysqlDb *gorm.DB) {

	//Handlers
	SqlHandler.InitializeUserHandler(mysqlDb)
}

// RegisterUser for adding user to db
func RegisterUser(c *gin.Context) {

	if c.PostForm("email") == "" {
		utils.Logger().Errorln("email not found!")
		utils.RespondError(c, http.StatusNotFound, "Required email")
		return
	}

	if c.PostForm("password") == "" {
		utils.Logger().Errorln("password not found!")
		utils.RespondError(c, http.StatusNotFound, "Required password")
		return
	}

	var name, email, password string
	name = c.PostForm("name")
	email = c.PostForm("email")
	password = c.PostForm("password")

	/* Check User if exist */
	_, err := SqlHandler.GetUserByEmail(email)

	if err == nil {
		utils.Logger().Errorln("User already exist with given email.")
		utils.RespondError(c, http.StatusNotFound, "User already exist with given email.")
		return
	}

	/* Creating User Model */
	dbModel := SqlModels.User{}
	dbModel.Name = name
	dbModel.Email = email
	dbModel.Password = password

	/* Bcrypting password */
	dbModel.Password = _bcryptingPassword(dbModel.Password)

	fmt.Println("dbModel is: ", dbModel)

	/* Calling Handler For Adding*/
	SqlHandler.AddUser(c, dbModel)

}

// AuthenticateUser for login/token
func AuthenticateUser(c *gin.Context) {

	fmt.Println("Req body: ", c.Request.ParseForm())

	if c.PostForm("email") == "" {
		utils.Logger().Errorln("email not found!")
		utils.RespondError(c, http.StatusNotFound, "Required email")
		return
	}

	if c.PostForm("password") == "" {
		utils.Logger().Errorln("password not found!")
		utils.RespondError(c, http.StatusNotFound, "Required password")
		return
	}

	var email, password string
	email = c.PostForm("email")
	password = c.PostForm("password")

	/* Creating User Model */
	dbModel := SqlModels.User{}
	dbModel.Email = email
	dbModel.Password = password

	fmt.Println("req. email: ", email, c.PostForm("email"))

	/* Check if User exist */
	user, err := SqlHandler.GetUserByEmail(email)
	fmt.Println("User is: ", user)

	if err != nil {
		utils.Logger().Errorln(err.Error())
		utils.RespondError(c, http.StatusNotFound, err.Error())
		return
	}

	/* Checking password */
	isMatch := _comparePassword(user.Password, dbModel.Password)

	if !isMatch {
		utils.Logger().Errorln("Invalid Password.")
		utils.RespondError(c, http.StatusNotFound, "Invalid Password.")
		return
	}

	//Create JWT token
	tokenString, expiresAt := _generateAccessToken(user.ID)

	oModel := models.AuthToken{Token: tokenString, TokenType: "access-token", ExpiresIn: expiresAt}

	fmt.Println("\noModel: ", oModel)

	utils.RespondJSON(c, http.StatusOK, oModel)

}

/**
 * ==================================================================================//
 *                                In File Custom Functions
 * ==================================================================================//
 */

// Bcrypting password
func _bcryptingPassword(password string) string {

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}

// Compare Hashed Password
func _comparePassword(userPassword, password string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(password))

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return false
	}

	return true
}

// Generate Access Token
func _generateAccessToken(userID uint) (string, int64) {

	expiresAt := time.Now().Add(time.Hour * 72).Unix()

	token := jwt.New(jwt.SigningMethodHS256)

	token.Claims = &models.AuthTokenClaim{
		&jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
		userID,
	}

	tokenString, err := token.SignedString([]byte(utils.JwtSecret()))

	if err != nil {
		fmt.Println(err.Error())
	}

	return tokenString, expiresAt
}

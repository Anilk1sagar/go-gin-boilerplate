package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	models "gitlab.com/anilk1sagar/go_gin_test/src/models"
	utils "gitlab.com/anilk1sagar/go_gin_test/src/utils"

	jwt "github.com/dgrijalva/jwt-go"
)

// UserFind from verifying jwt token
var UserFind = func(c *gin.Context) {

	// List of endpoints that doesn't require auth
	notAuth := []string{
		// "/api/test/mysql/add",
		"/api/test/mysql/getAll", "/api/test/mysql/get/{name}", //test mysql routes
		"/api/user/register", "/api/user/auth", //User routes
	}

	//current request path
	requestPath := c.Request.URL.Path

	// Check if request does not need authentication, serve the request if it doesn't need it
	for _, value := range notAuth {

		if value == requestPath {
			c.Next()
			return
		}
	}

	// Grab the token from the header
	tokenHeader := c.Request.Header.Get("Authorization")

	// Token is missing, returns with error code 403 Unauthorized
	if tokenHeader == "" {
		utils.RespondError(c, http.StatusForbidden, "Missing auth token")
		c.Abort()
		return
	}

	// The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
	splitted := strings.Split(tokenHeader, " ")
	if len(splitted) != 2 {
		utils.RespondError(c, http.StatusForbidden, "Invalid/Malformed auth token")
		c.Abort()
		return
	}

	// Grab the token part
	tokenPart := splitted[1]
	tk := &models.AuthTokenClaim{}

	token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(utils.JwtSecret()), nil
	})

	if err != nil {
		fmt.Println("error: ", err.Error())
		utils.RespondError(c, http.StatusForbidden, err.Error())
		c.Abort()
		return
	}

	// If Token is invalid
	if !token.Valid {
		utils.RespondError(c, http.StatusForbidden, "Token is not valid.")
		c.Abort()
		return
	}

	// Parse token
	// fmt.Println("req UserId is: ", tk.UserID) //Useful for monitoring

	c.Set("userID", tk.UserID)
	// userID, _ := c.Get("userID")
	// fmt.Println("body get: ", userID)

	c.Next()
}

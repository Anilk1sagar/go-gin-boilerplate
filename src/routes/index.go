package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"gitlab.com/anilk1sagar/go_gin_test/src/middleware"
	Test "gitlab.com/anilk1sagar/go_gin_test/src/routes/test"
	User "gitlab.com/anilk1sagar/go_gin_test/src/routes/user"
)

// InitializeRoutes sets the all required routers
func InitializeRoutes(router *gin.Engine, mysqlDb *gorm.DB) {

	fmt.Println("___Routes Initialized")

	// Attach JWT auth middleware
	router.Use(middleware.UserFind)

	/*
	 **===== Mysql Test Routes =====
	 */
	Test.InitializeMysqlTest(mysqlDb)
	router.POST("/api/test/mysql/add", Test.AddMysqlTest)
	router.GET("/api/test/mysql/get/:name", Test.GetMysqlTest)
	router.GET("/api/test/mysql/getAll", Test.GetAllMysqlTests)

	/*
	 **===== User Auth Routes =====
	 */
	User.InitializeUser(mysqlDb)
	router.POST("/api/user/register", User.RegisterUser)
	router.POST("/api/user/auth", User.AuthenticateUser)
}

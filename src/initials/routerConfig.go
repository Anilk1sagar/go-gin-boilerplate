package initials

import (
	"fmt"

	"github.com/gin-gonic/gin"
	routes "gitlab.com/anilk1sagar/go_gin_test/src/routes"
)

// InitializeRouter for initializing Router
func (a *App) InitializeRouter() {

	fmt.Println("**==== Initializing Router ====**")

	gin.SetMode(gin.ReleaseMode)

	// Creating gin Instance
	a.Router = gin.New()

	//Router Global Middlewares
	a.Router.Use(gin.Logger())
	a.Router.Use(gin.Recovery())

	routes.InitializeRoutes(a.Router, a.MysqlDb)

}

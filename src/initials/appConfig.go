package initials

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"gitlab.com/anilk1sagar/go_gin_test/src/mysql"
	utils "gitlab.com/anilk1sagar/go_gin_test/src/utils"
	"golang.org/x/net/context"
)

// App has router and db instances
type App struct {
	Router  *gin.Engine
	MysqlDb *gorm.DB
}

// InitializeApp for initializing app configurations
func (a *App) InitializeApp(mysqlConfig *mysql.MysqlConfig) {

	// Connect Mysqldb
	a.ConnectMysqlDb(mysqlConfig)

	// Initialize Middlewares
	a.InitializeMiddlewares()

	// Initializing Router
	a.InitializeRouter()

}

// RunServer for starting Server
func (a *App) RunServer(host string) {

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	server := &http.Server{
		Addr: host,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      a.Router,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {

		fmt.Println("Server Started on port: ", host)

		if err := a.Router.Run(host); err != nil {
			utils.Logger().Errorln(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	utils.Logger().Warningln("Shutting down Server ...")

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		utils.Logger().Fatal("Server Shutdown err: ", err)
	}

	utils.Logger().Println("Server exiting")

}

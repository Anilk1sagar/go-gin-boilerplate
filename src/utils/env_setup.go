package utils

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload" // _ godotenv autoload
)

// InitializeEnvSetup for app
func InitializeEnvSetup() {

	Logger().Info("Success message")
	Logger().Warning("Warning message")
	Logger().Error("Error message")

	fmt.Println("**===== Setting up Environment =====**")

	fmt.Println("Node Env: ", os.Getenv("NODE_ENV"))

	// Load Env File
	envFile := path.Join(".env." + os.Getenv("NODE_ENV"))
	fmt.Println("Selected Env File is: ", envFile)

	err := godotenv.Load(envFile)

	if err != nil {
		log.Fatal("Error loading .env file", err.Error())
	}

	fmt.Println("Node Env is: ", os.Getenv("NODE_ENV"))

	fmt.Println("isProduction?: ", IsProduction())
	fmt.Println("Firebase DB Ref: ", FirebaseDbRef())
}

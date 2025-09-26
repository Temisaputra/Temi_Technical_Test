// @title           My REST API
// @version         1.0
// @description     Documentation for my REST API
// @termsOfService  http://swagger.io/terms/

// @contact.name   Temi Saputra
// @contact.email  youremail@example.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8085
// @BasePath  /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

package main

import (
	"fmt"
	"os"

	"github.com/Temisaputra/warOnk/cmd"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	// Jalankan Cobra root command
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

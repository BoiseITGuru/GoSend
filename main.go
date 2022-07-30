package main

import (
	"github.com/eurekadao/gosend/internal/database"
	"github.com/eurekadao/gosend/internal/server"
)

func main() {
	// Load Config File
	LoadAppConfig()

	// Initialize Database
	database.Connect(AppConfig.DB.GormConnection)
	database.Migrate()

	// Start Web Server
	server.StartWebServer(AppConfig.WebServer.Port)
}

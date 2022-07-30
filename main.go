package main

import (
	"github.com/eurekadao/gosend/internal/amazon"
	"github.com/eurekadao/gosend/internal/auth"
	"github.com/eurekadao/gosend/internal/database"
	"github.com/eurekadao/gosend/internal/server"
)

func main() {
	// Load Config File
	LoadAppConfig()

	// Initialize Database
	database.Connect(AppConfig.DB.GormConnection)
	database.Migrate()

	// Initialize AWS Defaults
	keys := amazon.KmsKeys{
		EncryptionKey: AppConfig.Aws.Kms.EncryptionKey,
		JwtKey:        AppConfig.Aws.Kms.JwtKey,
	}
	aws := amazon.BuildAWSClient(AppConfig.Aws.Region, keys)
	aws.Init()

	// Initialize Auth System
	auth.Init()

	// Start Web Server
	server.StartWebServer(AppConfig.WebServer.Port)
}

package main

import (
	"github.com/eurekadao/gosend/internal/amazon"
	"github.com/eurekadao/gosend/internal/auth"
	"github.com/eurekadao/gosend/internal/controllers"
	"github.com/eurekadao/gosend/internal/database"
	"github.com/gin-gonic/gin"
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

	r := gin.Default()

	// SendGrid Email Receiver Endpoint
	r.POST("/sendgrid", controllers.EmailReceiver)

	r.Run(":" + AppConfig.WebServer.Port)
}

package main

import (
	"github.com/eurekadao/gosend/internal/database"
	"github.com/eurekadao/gosend/sdk"
)

func main() {
	// Load Config File
	LoadAppConfig()

	// Initialize Database
	database.Connect(AppConfig.DB.GormConnection)
	database.Migrate()

	sdk := sdk.SdkConfig{
		Instance:      database.Instance,
		Region:        AppConfig.Aws.Region,
		EncryptionKey: AppConfig.Aws.Kms.EncryptionKey,
		JwtKey:        AppConfig.Aws.Kms.JwtKey,
		Port:          AppConfig.WebServer.Port,
	}

	sdk.Configure().Start()
}

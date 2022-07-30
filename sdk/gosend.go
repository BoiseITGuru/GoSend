package sdk

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SDK struct {
	Database *gorm.DB
}

func Connect(connectionString string) SDK {
	var sdk *SDK
	var dbError error

	sdk.Database, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to GoSend Database")
	}
	log.Println("Connected to GoSend Database!")

	return *sdk
}

func (sdk *SDK) CreateMailbox(name string, addresses []string) bool {
	return false
}

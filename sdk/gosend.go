package sdk

import (
	"github.com/eurekadao/gosend/internal/amazon"
	"github.com/eurekadao/gosend/internal/auth"
	"github.com/eurekadao/gosend/internal/controllers"
	"github.com/eurekadao/gosend/internal/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SdkConfig struct {
	Instance      *gorm.DB
	Region        string
	EncryptionKey string
	JwtKey        string
	Port          string
}

type Server struct {
	port string
	keys amazon.KmsKeys
	aws  *amazon.AwsSvc
}

func (config *SdkConfig) Configure() Server {
	var s Server

	database.SdkDbStart(config.Instance)
	s.port = config.Port

	// Initialize AWS Defaults
	s.keys = amazon.KmsKeys{
		EncryptionKey: config.EncryptionKey,
		JwtKey:        config.JwtKey,
	}
	s.aws = amazon.BuildAWSClient(config.Region, s.keys)

	return s
}

func (s Server) Start() {
	s.aws.Init()

	// Initialize Auth System
	auth.Init()

	r := gin.Default()

	// SendGrid Email Receiver Endpoint
	r.POST("/sendgrid", controllers.EmailReceiver)

	r.Run(":" + s.port)
}

func (s *Server) CreateMailbox(name string, addresses []string) bool {
	return false
}

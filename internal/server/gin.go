package server

import (
	"github.com/eurekadao/gosend/internal/server/controllers"
	"github.com/gin-gonic/gin"
)

func StartWebServer(port string) {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// SendGrid Email Receiver Endpoint
	r.POST("/sendgrid", controllers.EmailReceiver)

	r.Run(":" + port)
}

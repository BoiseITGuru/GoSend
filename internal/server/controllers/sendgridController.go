package controllers

import (
	"net/http"

	"github.com/eurekadao/gosend/internal/database"
	"github.com/eurekadao/gosend/internal/models"
	"github.com/gin-gonic/gin"
)

func EmailReceiver(context *gin.Context) {
	var email *models.Email

	if err := context.ShouldBind(&email); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	database.Instance.Save(email)
	context.JSON(http.StatusOK, gin.H{"status": "OK"})
}

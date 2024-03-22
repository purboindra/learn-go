package routes

import (
	"example/rest_api/db"
	"example/rest_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signUp(context *gin.Context) {
	db.InitDB()
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	context.JSON(http.StatusOK, gin.H{"message": "User created successfully!"})
}

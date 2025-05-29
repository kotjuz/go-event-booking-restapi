package routes

import (
	"net/http"

	"example.com/eventapi/models"
	"github.com/gin-gonic/gin"
)

func createUser(context *gin.Context) {
	var newUser models.User
	err := context.ShouldBindJSON(&newUser)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
	}

	err = newUser.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user"})
	}
	context.JSON(http.StatusOK, gin.H{"message": "Created new user"})
}

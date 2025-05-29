package routes

import (
	"net/http"

	"example.com/eventapi/models"
	"example.com/eventapi/utils"
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

func loginUser(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not auth user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login succesful!", "token:": token})
}

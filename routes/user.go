package routes

import (
	"billboard/controllers"

	"github.com/gin-gonic/gin"
)

func userHandler(router *gin.Engine, userController *controllers.User) {
	router.POST("/auth/signup", userController.Signup)
	router.POST("/auth/login", userController.Login)
}

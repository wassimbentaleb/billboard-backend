package routes

import (
	"billboard/controllers"

	"github.com/gin-gonic/gin"
)

func userHandler(router *gin.Engine, userController *controllers.User) {
	router.POST("/clientCreate", userController.Signup)
	router.GET("/clients", userController.FindAll)
	router.PUT("/client/:id", userController.Update)
	router.DELETE("/client/:id", userController.Delete)

	router.POST("/auth/signup", userController.Signup)

	router.POST("/auth/login", userController.Login)
}

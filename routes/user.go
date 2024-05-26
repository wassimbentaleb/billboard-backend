package routes

import (
	"billboard/controllers"

	"github.com/gin-gonic/gin"
)

func userHandler(router *gin.Engine, userController *controllers.User) {
	router.POST("/client", userController.Create)
	router.GET("/clients", userController.FindAll)
	router.GET("/client/current", userController.FindCurrent)
	router.PUT("/client/:id", userController.Update)
	router.DELETE("/client/:id", userController.Delete)

	router.POST("/auth/login", userController.Login)
}

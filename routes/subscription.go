package routes

import (
	"billboard/controllers"
	"github.com/gin-gonic/gin"
)

func subscriptionHandler(router *gin.Engine, subscriptionController *controllers.Subscription) {
	router.POST("/subscription", subscriptionController.Create)
	router.GET("/subscriptions", subscriptionController.FindAll)
	router.GET("/subscription/:id", subscriptionController.FindByID)
	router.PUT("/subscription/:id", subscriptionController.Update)
	router.DELETE("/subscription/:id", subscriptionController.Delete)
}

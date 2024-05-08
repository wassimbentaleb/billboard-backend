package routes

import (
	"billboard/controllers"

	"github.com/gin-gonic/gin"
)

func planHandler(router *gin.Engine, planController *controllers.Plan) {
	router.POST("/plan", planController.Create)
	router.GET("/plans", planController.FindAll)
	router.GET("/plans/:boardId", planController.FindByBoardID)
	router.GET("/plan/:id", planController.FindByID)
	router.PUT("/plan/:id", planController.Update)
	router.DELETE("/plan/:id", planController.Delete)
}

package routes

import (
	"billboard/controllers"

	"github.com/gin-gonic/gin"
)

func boardHandler(router *gin.Engine, boardController *controllers.Board) {
	router.POST("/board", boardController.Create)
	router.GET("/boards", boardController.FindAll)
	router.GET("/board/:id", boardController.FindByID)
	router.PUT("/board/:id", boardController.Update)
	router.DELETE("/board/:id", boardController.Delete)

	router.PUT("/board", boardController.LinkBoardUpdate)
}

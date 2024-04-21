package routes

import (
	"billboard/controllers"

	"github.com/gin-gonic/gin"
)

func uploadHandler(router *gin.Engine, uploadController *controllers.Upload) {
	router.POST("/upload", uploadController.FileUpload)
}

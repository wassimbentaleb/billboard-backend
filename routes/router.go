package routes

import (
	"billboard/controllers"
	"billboard/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(
	userController *controllers.User,
	boardController *controllers.Board,
	planController *controllers.Plan,
	uploadController *controllers.Upload,
	middleware *middleware.AuthMiddleware,
) *gin.Engine {
	// create a new gin router
	router := gin.Default()

	// enable cors
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"*"},
		AllowHeaders:    []string{"*"},
	}))

	// use auth middleware
	router.Use(middleware.RequireAuth)

	// serve static files
	router.Static("/media", "./static")

	// define routes
	userHandler(router, userController)
	boardHandler(router, boardController)
	planHandler(router, planController)
	uploadHandler(router, uploadController)

	return router
}

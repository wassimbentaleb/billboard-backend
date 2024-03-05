package main

import (
	"awesomeProject1/controllers"
	"awesomeProject1/initializers"
	"awesomeProject1/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()

}

func main() {
	r := gin.Default()

	// Middleware CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"X-Requested-With", "Content-Type", "Authorization", "ngrok-skip-browser-warning"},
	}))

	// Définition des routes
	//Auth
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate/:id", middleware.RequireAuth, controllers.Validate)
	r.DELETE("/delete/:id", middleware.RequireAuth, controllers.UserDelete)

	//CreateNewUser
	r.POST("/newuser", controllers.PostsCreate)
	r.PUT("/newuser/:Id", controllers.PostsUpdate)
	r.GET("/newuser", controllers.PostsIndex)
	r.GET("/newuser/:Id", controllers.PostsShow)
	r.DELETE("/newuser/:Id", controllers.PostsDelete)

	// Lancement du serveur
	r.Run() // écoute et sert sur 0.0.0.0:8080
}

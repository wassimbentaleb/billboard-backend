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
	r.Static("/media", "./static")
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	// Middleware CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{
			"X-Requested-With",
			"Content-Type",
			"Authorization",
			"ngrok-skip-browser-warning",
			"Access-Control-Allow-Headers",
			"Custom-Headers",
		},
	}))

	// Définition des routes
	//Auth
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate/:id", middleware.RequireAuth, controllers.Validate)
	r.DELETE("/delete/:id", middleware.RequireAuth, controllers.UserDelete)

	//CreateNewUser
	r.POST("/clientCreate", controllers.ClientCreate)
	r.PUT("/clientUpdate/:Id", controllers.ClientUpdate)
	r.GET("/clientsGet", controllers.ClientIndex)
	r.GET("/clientGet/:Id", controllers.ClientShow)
	r.DELETE("/clientDelete/:Id", controllers.ClientDelete)

	//UploadImage
	r.POST("/handleFileUpload", controllers.HandleFileUpload)
	r.POST("/HandleAddPlan", controllers.HandleAddPlan)

	// Lancement du serveur
	r.Run() // écoute et sert sur 0.0.0.0:8080
}

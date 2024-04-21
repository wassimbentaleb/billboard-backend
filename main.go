package main

import (
	"billboard/database"
	"billboard/middleware"
	"billboard/routes"
	"fmt"
	"log"
	"net/http"
	"os"

	"billboard/controllers"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	pg, err := database.NewPostgres()
	if err != nil {
		log.Fatalf("Failed to connect to database %v", err)
	}

	userController := controllers.NewUser(pg)
	boardController := controllers.NewBoard(pg)
	planController := controllers.NewPlan(pg)
	uploadController := controllers.NewUpload()
	authMiddleware := middleware.NewAuthMiddleware(pg)

	router := routes.NewRouter(
		userController,
		boardController,
		planController,
		uploadController,
		authMiddleware,
	)

	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	err = http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}

	// Définition des routes
	//Auth
	//r.POST("/signup", controllers.Signup)
	//r.POST("/login", controllers.Login)
	//r.GET("/validate/:id", middleware.RequireAuth, controllers.Validate)
	//r.DELETE("/delete/:id", middleware.RequireAuth, controllers.UserDelete)
	//
	////CreateNewUser
	//r.POST("/clientCreate", controllers.ClientCreate)
	//r.PUT("/clientUpdate/:Id", controllers.ClientUpdate)
	//r.GET("/clientsGet", controllers.ClientIndex)
	//r.GET("/clientGet/:Id", controllers.ClientShow)
	//r.DELETE("/clientDelete/:Id", controllers.ClientDelete)
	//
	////CreateNewBillboard

	//
	////UploadImage
	//r.POST("/handleFileUpload", controllers.HandleFileUpload)
	//r.POST("/HandleAddPlan", controllers.HandleAddPlan)
	//r.GET("/GetAllPlans", controllers.GetAllPlans)
	//r.GET("/GetPlansByBoardId/:boardId", controllers.GetPlansByBoardId)

}

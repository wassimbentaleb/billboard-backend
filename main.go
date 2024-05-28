package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"billboard/controllers"
	"billboard/database"
	"billboard/middleware"
	"billboard/routes"
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
	subscriptionController := controllers.NewSubscription(pg)
	uploadController := controllers.NewUpload()
	authMiddleware := middleware.NewAuthMiddleware(pg)

	router := routes.NewRouter(
		userController,
		boardController,
		planController,
		uploadController,
		subscriptionController,
		authMiddleware,
	)

	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	err = http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}
}

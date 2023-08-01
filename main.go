package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/saifuljnu/crudApp/database"
	"github.com/saifuljnu/crudApp/handler"
	"github.com/saifuljnu/crudApp/service"
)

func main() {
	// Set Gin to release mode.
	gin.SetMode(gin.ReleaseMode)

	// Initialize MongoDB connection
	dbURI := "localhost:27017/crudDB"
	dbClient, err := database.NewMongoClient(dbURI)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}
	defer dbClient.Disconnect(context.Background())

	// Create a new Gin router.
	r := gin.Default()

	// Create instances of services and handlers
	userService := service.NewUserService(dbClient)
	userHandler := handler.NewUserHandler(userService)

	// Define API routes and handlers
	r.POST("/api/collectionUser", userHandler.InsertUser)
	r.GET("/api/collectionUser", userHandler.GetAllUsers)
	r.DELETE("/api/collectionUser/:id", userHandler.DeleteUser)
	r.PUT("/api/collectionUser/:id", userHandler.UpdateUser)

	// Start the server
	fmt.Println("Server started on http://localhost:3000")
	r.Run(":3000")
}

package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/saifuljnu/crudApp/database"
	"github.com/saifuljnu/crudApp/handler"
	"github.com/saifuljnu/crudApp/middlewares"
	"github.com/saifuljnu/crudApp/service"
)

// LoggerMiddleware logs the client IP, Response HTTP Status Code, and Latency for every request
// func LoggerMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		startTime := time.Now()

// 		// for taking client ip
// 		clientIP := c.ClientIP()

// 		// for calling the next handler in the chain
// 		c.Next()

// 		// for latency
// 		latency := time.Since(startTime)

// 		// Log the info
// 		log.Printf("Client IP: %s | Status Code: %d | Latency: %s | RequestMethod: %s | RequestURI: %s",
// 			clientIP, c.Writer.Status(), latency, c.Request.Method, c.Request.RequestURI)
// 	}
// }

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

	// Using the custom middleware for every route
	//r.Use(LoggerMiddleware())

	r.Use(gin.Recovery(), middlewares.Logger())

	// Create instances of services and handlers
	userService := service.NewUserService(dbClient)
	userHandler := handler.NewUserHandler(userService)

	r.POST("/api/collectionUser", userHandler.InsertUser)
	r.GET("/api/collectionUser", userHandler.GetAllUsers)
	r.DELETE("/api/collectionUser/:id", userHandler.DeleteUser)
	r.PUT("/api/collectionUser/:id", userHandler.UpdateUser)

	fmt.Println("Server started on http://localhost:3000")
	r.Run(":3000")
}

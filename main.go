package main

import (
	"github.com/ThomasPhilipp/go-event-api/db"
	"github.com/ThomasPhilipp/go-event-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// init database
	db.InitDB()

	// init HTTP server
	server := gin.Default()

	// register API routes
	routes.RegisterRoutes(server)

	// start server and listen to incoming requests
	server.Run(":8080")
}

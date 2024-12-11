package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"zwickit.com/go-event-api/models"
)

func main() {
	// init HTTP server
	server := gin.Default()

	// define request paths
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	// start server and listen to incoming requests
	server.Run(":8080")
}

func getEvents(ctx *gin.Context) {
	events := models.GetAllEvents()
	// send JSON response
	ctx.JSON(http.StatusOK /*status code*/, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event) // make some magic and map the values to the event

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	event.ID = 1
	event.UserID = 1

	event.Save()

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

package routes

import (
	"net/http"
	"strconv"

	"github.com/ThomasPhilipp/go-event-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10 /*base*/, 64 /*bit size*/)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user for event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered"})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10 /*base*/, 64 /*bit size*/)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	var event models.Event
	event.ID = eventId

	event.CancelRegistration(userId)

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration for event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Canceled"})
}

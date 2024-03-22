package routes

import (
	"example/rest_api/db"
	"example/rest_api/models"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvent(context *gin.Context) {

	param := context.Param("id")

	log.Println(param)

	id, err := strconv.ParseInt(param, 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	event, err := models.GetEventById(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, event)
}

func getEvents(context *gin.Context) {

	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	db.InitDB()
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		fmt.Println("SOMETHING WENT WRONG", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	event.ID = 1
	event.UserId = 1

	err = event.Save()

	if err != nil {
		fmt.Println("SOMETHING WENT WRONG 2", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}

func updateEvent(context *gin.Context) {
	param := context.Param("id")

	log.Println(param)

	id, err := strconv.ParseInt(param, 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	_, err = models.GetEventById(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	var updatedEvent models.Event

	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	updatedEvent.ID = id
	err = updatedEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Event updated successfully",
	})
}

func deleteEvent(context *gin.Context) {
	parseId := context.Param("id")
	id, err := strconv.ParseInt(parseId, 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	event, err := models.GetEventById(id)

	log.Println("EVENTTTTT", event)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	err = event.Delete()

	log.Println("EVENTTTTT AFTER DLEETE", event)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Event deleted successfully",
	})
}

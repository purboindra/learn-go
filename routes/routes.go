package routes

import (
	"example/rest_api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	authenticated.GET("/events", getEvents)
	authenticated.GET("/events/:id", getEvent)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	// TODO FIX DELETE
	server.DELETE("/events/:id", deleteEvent)
	server.POST("/signup", signUp)
	server.POST("/login", login)
}

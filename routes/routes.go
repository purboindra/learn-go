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
	authenticated.DELETE("/events/:id", deleteEvent)

	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", signUp)
	server.POST("/login", login)
}

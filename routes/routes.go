package routes

import (
	"github.com/gin-gonic/gin"
	"osho.com/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventById)
	//  way to authenticate all routes
	authenticated :=server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvents)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register",registerForEvent)
	authenticated.DELETE("/events/:id/unregister",cancelRegisteration)

    //  way to authenticate 1 route
	// server.POST("/events",middlewares.Authenticate, createEvents)
	// server.PUT("/events/:id", updateEvent)
	// server.DELETE("/events/:id", deleteEvent)

	server.POST("/signup",signup)
	server.POST("/login",login)
}

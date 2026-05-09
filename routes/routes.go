package routes

import (
	"database/sql"
	"event_booking/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine, DB *sql.DB) {
	server.GET("/events", func(ctx *gin.Context) { GetEvents(ctx, DB) })
	server.GET("/events/:id", func(ctx *gin.Context) { GetEvent(ctx, DB) })

	authenticated := server.Group("/")
	authenticated.Use(func(ctx *gin.Context) { middleware.Authentication(ctx) })
	authenticated.POST("/events", func(ctx *gin.Context) { CreateEvents(ctx, DB) })
	authenticated.PUT("/events/:id", func(ctx *gin.Context) { UpdateEvent(ctx, DB) })
	authenticated.DELETE("/events/:id", func(ctx *gin.Context) { DeleteEvent(ctx, DB) })

	server.POST("/signup", func(ctx *gin.Context) { Signup(ctx, DB) })
	server.POST("/login", func(ctx *gin.Context) { Login(ctx, DB) })
}

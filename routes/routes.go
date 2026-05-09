package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine, DB *sql.DB) {
	server.GET("/events", func(ctx *gin.Context) { GetEvents(ctx, DB) })
	server.GET("/events/:id", func(ctx *gin.Context) { GetEvent(ctx, DB) })
	server.POST("/events", func(ctx *gin.Context) { CreateEvents(ctx, DB) })      // Protected
	server.PUT("/events/:id", func(ctx *gin.Context) { UpdateEvent(ctx, DB) })    // Protected
	server.DELETE("/events/:id", func(ctx *gin.Context) { DeleteEvent(ctx, DB) }) // Protected
	server.POST("/signup", func(ctx *gin.Context) { Signup(ctx, DB) })
	server.POST("/login", func(ctx *gin.Context) { Login(ctx, DB) })
}

package routes

import (
	"database/sql"
	_ "event_booking/docs" // hasil swag init
	"event_booking/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//

func RegisterRoutes(server *gin.Engine, DB *sql.DB) {
	server.GET("/events", func(ctx *gin.Context) { GetEvents(ctx, DB) })
	server.GET("/events/:id", func(ctx *gin.Context) { GetEvent(ctx, DB) })

	authenticated := server.Group("/")
	authenticated.Use(func(ctx *gin.Context) { middleware.Authentication(ctx) })
	authenticated.POST("/events", func(ctx *gin.Context) { CreateEvents(ctx, DB) })
	authenticated.PUT("/events/:id", func(ctx *gin.Context) { UpdateEvent(ctx, DB) })
	authenticated.DELETE("/events/:id", func(ctx *gin.Context) { DeleteEvent(ctx, DB) })
	authenticated.POST("/events/:id/register", func(ctx *gin.Context) { RegisterForEvent(ctx, DB) })
	authenticated.DELETE("/events/:id/register", func(ctx *gin.Context) { CancelForEvent(ctx, DB) })

	server.POST("/signup", func(ctx *gin.Context) { Signup(ctx, DB) })
	server.POST("/login", func(ctx *gin.Context) { Login(ctx, DB) })

	// Swagger
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

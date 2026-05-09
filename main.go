package main

import (
	"event_booking/db"
	"event_booking/routes"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

// @title           Event Booking API
// @version         1.0
// @description     REST API untuk manajemen event dan registrasi
// @host            event-bookings.burhanudin.com
// @BasePath        /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	DB, err := db.InitDB()
	if err != nil {
		fmt.Println("Failed Connect To Database")
		log.Fatal(err)
		return
	}

	server := gin.Default()
	routes.RegisterRoutes(server, DB)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err = server.Run(":" + port); err != nil {
		panic(err)
	}

}

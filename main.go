package main

import (
	"event_booking/db"
	"event_booking/routes"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	DB, err := db.InitDB()
	if err != nil {
		fmt.Println("Failed Connect To Database")
		log.Fatal(err)
		return
	}

	server := gin.Default()
	routes.RegisterRoutes(server, DB)

	err = server.Run(":8080")
	if err != nil {
		panic(err)
	}

}

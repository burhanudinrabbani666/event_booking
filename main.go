package main

import (
	"database/sql"
	"event_booking/db"
	"event_booking/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var DB *sql.DB

func main() {
	var err error
	DB, err = db.InitDB()
	if err != nil {
		fmt.Println(err)
		return
	}

	server := gin.Default()

	server.GET("/events", GetEvents)
	server.GET("/events/:id", GetEvent)
	server.POST("/events", CreateEvents)

	err = server.Run(":8080")
	if err != nil {
		panic(err)
	}

}

func GetEvents(ctx *gin.Context) {

	events, err := models.GetAllEvents(DB)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]any{
			"code":   http.StatusInternalServerError,
			"status": "INTERNAL SERVER ERROR!",
		})
		return

	}

	ctx.JSON(http.StatusOK, events)
}

func GetEvent(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"code":   http.StatusBadRequest,
			"status": "BAD REQUEST!",
		})
		return
	}

	event, err := models.GetEventById(DB, id)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, map[string]any{
			"code":   http.StatusInternalServerError,
			"status": "FAILED GET EVENT BY ID. INTERNAL SERVER ERROR",
			"error":  err,
		})
		return
	}

	if event == nil {
		ctx.JSON(http.StatusNotFound, map[string]any{
			"code":   http.StatusNotFound,
			"status": "EVENT NOT FOUND",
		})
		return

	}

	ctx.JSON(http.StatusOK, map[string]any{
		"code":   http.StatusOK,
		"status": "SUCCESS GET EVENT BY ID",
		"data":   event,
	})

}

func CreateEvents(ctx *gin.Context) {

	var event models.Event
	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":   http.StatusBadRequest,
			"status": "FAILED TO CREATE EVENT!",
		})

		return
	}

	// TODO: Change later
	event.User_id = 1

	err = event.Save(DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":   http.StatusBadRequest,
			"status": "FAILED TO CREATE EVENT!",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"code":   http.StatusOK,
		"status": "SUCCESS TO CREATE EVENT!",
		"data":   event,
	})

}

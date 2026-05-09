package routes

import (
	"database/sql"
	"event_booking/models"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetEvents(ctx *gin.Context, DB *sql.DB) {

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

func GetEvent(ctx *gin.Context, DB *sql.DB) {
	id, err := strconv.Atoi(ctx.Param("id"))
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

func CreateEvents(ctx *gin.Context, DB *sql.DB) {

	var event models.Event
	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":   http.StatusBadRequest,
			"status": "FAILED TO CREATE EVENT!",
			"error":  err,
		})

		return
	}

	// TODO: Change later
	userId := ctx.GetInt("userId")
	event.UserId = userId

	err = event.Create(DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":   http.StatusBadRequest,
			"status": "FAILED TO CREATE EVENT!",
			"error":  err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"code":   http.StatusOK,
		"status": "SUCCESS TO CREATE EVENT!",
		"data":   event,
	})

}

func UpdateEvent(ctx *gin.Context, DB *sql.DB) {
	eventId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"code":   http.StatusBadRequest,
			"status": "BAD REQUEST!",
		})
		return
	}

	event, err := models.GetEventById(DB, eventId)
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

	userId := ctx.GetInt("userId")
	if userId != event.UserId {
		ctx.JSON(http.StatusUnauthorized, map[string]any{
			"code":     http.StatusUnauthorized,
			"status":   "NOT AUTHORIZED FOR UPDATED EVENT!",
			"expected": event.UserId,
			"insert":   userId,
		})
		return

	}

	var updateEvent models.EventCompleteData
	err = ctx.ShouldBindJSON(&updateEvent)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"code":   http.StatusBadRequest,
			"status": "FAILED UPDATE EVENT. BAD REQUEST.",
		})
		return
	}

	updateEvent.Id = eventId
	now := time.Now()
	updateEvent.UpdatedAt = &now

	err = updateEvent.Update(DB)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]any{
			"code":   http.StatusInternalServerError,
			"status": "FAILED UPDATE EVENT. INTERNAL SERVER ERROR.",
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string]any{
		"code":   http.StatusOK,
		"status": "SUCCESS UPDATE EVENT.",
		"data":   updateEvent,
	})

}

func DeleteEvent(ctx *gin.Context, DB *sql.DB) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"code":   http.StatusBadRequest,
			"status": "BAD REQUEST!",
		})
		return
	}

	event, err := models.GetEventById(DB, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]any{
			"code":   http.StatusInternalServerError,
			"status": "FAILED TO CHECK EVENT. INTERNAL SERVER ERROR!",
		})
		return
	}

	if event == nil {
		ctx.JSON(http.StatusNotFound, map[string]any{
			"code":   http.StatusNotFound,
			"status": "FAILED TO DELETE EVENT. NOT FOUND.",
		})
		return
	}

	userId := ctx.GetInt("userId")
	if userId != event.UserId {
		ctx.JSON(http.StatusUnauthorized, map[string]any{
			"code":   http.StatusUnauthorized,
			"status": "NOT AUTHORIZED FOR UPDATED EVENT!",
		})
		return

	}

	err = models.DeleteEventById(DB, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]any{
			"code":   http.StatusInternalServerError,
			"status": "FAILED TO DELETE EVENT. INTERNAL SERVER ERROR.",
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string]any{
		"code":   http.StatusOK,
		"status": "SUCCESS DELETE EVENT!",
	})

}

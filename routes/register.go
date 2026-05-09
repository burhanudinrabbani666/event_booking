package routes

import (
	"database/sql"
	"event_booking/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterForEvent(ctx *gin.Context, DB *sql.DB) {
	userId := ctx.GetInt("userId")
	eventId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"code":   http.StatusBadRequest,
			"status": "BAD REQUEST!",
		})
		return
	}

	events, err := models.GetEventById(DB, eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]any{
			"code":   http.StatusInternalServerError,
			"status": "FAILED REGISTER TO EVENT, INTERNAL SERVER ERROR.",
		})
		return
	}

	if events == nil {
		ctx.JSON(http.StatusNotFound, map[string]any{
			"code":   http.StatusNotFound,
			"status": "FAILED REGISTER TO EVENT, EVENTS NOT FOUND.",
		})
		return
	}

	err = events.Register(userId, DB)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]any{
			"code":   http.StatusInternalServerError,
			"status": "FAILED REGISTER TO EVENT, INTERNAL SERVER ERROR.",
		})
		return
	}

	ctx.JSON(http.StatusCreated, map[string]any{
		"code":   http.StatusCreated,
		"status": "SUCCESS REGISTER TO EVENT.",
	})

}

func CancelForEvent(ctx *gin.Context, DB *sql.DB) {
	userId := ctx.GetInt("userId")
	eventId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"code":   http.StatusBadRequest,
			"status": "BAD REQUEST!",
		})
		return
	}

	var event models.Event
	event.Id = eventId
	err = event.CancelRegister(userId, DB)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]any{
			"code":   http.StatusInternalServerError,
			"status": "FAILED REGISTER TO EVENT, INTERNAL SERVER ERROR.",
		})
		return
	}

	ctx.JSON(http.StatusCreated, map[string]any{
		"code":   http.StatusCreated,
		"status": "SUCCESS CANCEL REGISTER.",
	})

}

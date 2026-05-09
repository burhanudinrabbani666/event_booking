package routes

import (
	"database/sql"
	"event_booking/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// RegisterForEvent godoc
// @Summary      Register ke event
// @Tags         Registrations
// @Security     BearerAuth
// @Produce      json
// @Param        id    path      int  true  "Event ID"
// @Success      201   {object}  map[string]any
// @Failure      400   {object}  map[string]any
// @Failure      401   {object}  map[string]any
// @Failure      404   {object}  map[string]any
// @Failure      500   {object}  map[string]any
// @Router       /events/{id}/register [post]
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

// CancelForEvent godoc
// @Summary      Batalkan registrasi event
// @Tags         Registrations
// @Security     BearerAuth
// @Produce      json
// @Param        id    path      int  true  "Event ID"
// @Success      200   {object}  map[string]any
// @Failure      400   {object}  map[string]any
// @Failure      500   {object}  map[string]any
// @Router       /events/{id}/register [delete]
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

	ctx.JSON(http.StatusOK, map[string]any{
		"code":   http.StatusOK,
		"status": "SUCCESS CANCEL REGISTER.",
	})

}

package routes

import (
	"database/sql"
	"event_booking/models"
	"event_booking/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Signup godoc
// @Summary      Signup new User
// @Tags         Auth
// @Accept		 json
// @Param        body   body    models.User true "User Data"
// @Success      200  {object}  map[string]any
// @Failure      400  {object}  map[string]any
// @Failure      500  {object}  map[string]any
// @Router       /signup [post]
func Signup(ctx *gin.Context, DB *sql.DB) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":   http.StatusBadRequest,
			"status": "FAILED TO CREATE USER, BAD REQUEST.",
		})

		return
	}

	err = user.SignUp(DB)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":   http.StatusInternalServerError,
			"status": "FAILED TO CREATE USER, INTERNAL SERVER ERROR.",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"code":   http.StatusCreated,
		"status": "SUCCESS TO CREATE NEW USER.",
	})

}

// Login godoc
// @Summary      Login User
// @Tags         Auth
// @Accept		 json
// @Produce		 json
// @Param        body   body    models.User true "User Data"
// @Success      200  {object}  map[string]any
// @Failure      400  {object}  map[string]any
// @Failure      401  {object}  map[string]any
// @Failure      500  {object}  map[string]any
// @Router       /login [post]
func Login(ctx *gin.Context, DB *sql.DB) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":   http.StatusBadRequest,
			"status": "FAILED TO LOGIN, BAD REQUEST.",
		})

		return
	}

	isValid, err := user.ValidateCredential(DB)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":   http.StatusInternalServerError,
			"status": "FAILED TO LOGIN, INTERNAL SERVER ERROR.",
		})
		return
	}

	if !isValid {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code":   http.StatusUnauthorized,
			"status": "FAILED TO LOGIN, INVALID CREDENTIALS.",
		})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":   http.StatusInternalServerError,
			"status": "FAILED TO LOGIN, INTERNAL SERVER ERROR.",
		})
		return

	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"status": "SUCCESS TO LOGIN.",
		"token":  token,
	})

}

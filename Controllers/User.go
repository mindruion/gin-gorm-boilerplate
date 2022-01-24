package Controllers

import (
	"github.com/gin-gonic/gin"
	"gorm-gin/ApiHelpers"
	"gorm-gin/Middlewares"
	_ "gorm-gin/Models"
	"net/http"
)

// Login
// @Summary Login user
// @Description Get jwt token with credentials
// @Tags Auth
// @Accept */*
// @Book json
// @Param        message  body      Models.Login  true  "Login"
// @Router /login [post]
func Login(c *gin.Context) {
	return
}

// Me
// @Summary get info about current logged user.
// @Description get info about current logged user.
// @Tags User
// @Accept */*
// @Book json
// @Router /me [get]
// @Security ApiKeyAuth
// @Success 200 {object} Models.User
func Me(c *gin.Context) {
	currentUser := Middlewares.GetLoggedUser(c)
	ApiHelpers.RespondJSON(c, http.StatusOK, currentUser)
}

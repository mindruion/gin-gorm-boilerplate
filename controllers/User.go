package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm-gin/apiHelpers"
	"gorm-gin/middlewares"
	_ "gorm-gin/models"
	"net/http"
)

// Login
// @Summary Login user
// @Description Get jwt token with credentials
// @Tags Auth
// @Accept */*
// @Book json
// @Param        message  body      models.Login  true  "Login"
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
// @Success 200 {object} models.User
func Me(c *gin.Context) {
	currentUser := middlewares.GetLoggedUser(c)
	apiHelpers.RespondJSON(c, http.StatusOK, currentUser)
}

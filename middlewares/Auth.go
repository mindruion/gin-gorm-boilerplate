package middlewares

import (
	"errors"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"gorm-gin/config"
	"gorm-gin/models"
	"log"
	"time"
)

func getJwtMiddleware() *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "test zone",
		Key:             []byte(config.EnvConfigs["SECRET_KEY"]),
		Timeout:         time.Hour,
		MaxRefresh:      time.Hour,
		IdentityKey:     config.EnvConfigs["IDENTITY_KEY"],
		PayloadFunc:     handlePayload,
		IdentityHandler: handleIdentity,
		Authenticator:   handleAuthentication,
		Authorizator:    handleAuthorized,
		Unauthorized:    handleUnauthorized,
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	return authMiddleware
}
func handlePayload(data interface{}) jwt.MapClaims {
	if v, ok := data.(*models.User); ok {
		return jwt.MapClaims{
			config.EnvConfigs["IDENTITY_KEY"]: v.Email,
		}
	}
	return jwt.MapClaims{}
}
func handleUnauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

func handleIdentity(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	var user models.User
	config.DB.Model(&models.User{}).First(&user, "email = ?", claims[config.EnvConfigs["IDENTITY_KEY"]].(string))
	return &user
}

func handleAuthentication(c *gin.Context) (interface{}, error) {
	var loginVals models.Login
	if err := c.ShouldBindJSON(&loginVals); err != nil {
		return "", errors.New("missing Email or Password")
	}
	userEmail := loginVals.Email
	password := loginVals.Password
	var user models.User
	errDB := config.DB.Model(&models.User{}).First(&user, "email = ?", userEmail)
	if errDB.Error != nil {
		return nil, jwt.ErrFailedAuthentication
	}
	invalidPassword := user.VerifyPassword(password)
	if invalidPassword != nil {
		return nil, jwt.ErrFailedAuthentication
	}
	return &user, nil
}

func handleAuthorized(data interface{}, c *gin.Context) bool {
	if _, ok := data.(*models.User); ok {
		return true
	}

	return false
}

func InitAuthMiddleware() *jwt.GinJWTMiddleware {
	authMiddleware := getJwtMiddleware()
	errInit := authMiddleware.MiddlewareInit()
	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}
	return authMiddleware
}

func GetLoggedUser(c *gin.Context) *models.User {
	user, _ := c.Get(config.EnvConfigs["IDENTITY_KEY"])
	return user.(*models.User)
}

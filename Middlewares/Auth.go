package Middlewares

import (
	"errors"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"gorm-gin/Config"
	"gorm-gin/Models"
	"log"
	"time"
)

func getJwtMiddleware() *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "test zone",
		Key:             []byte(Config.EnvConfigs["SECRET_KEY"]),
		Timeout:         time.Hour,
		MaxRefresh:      time.Hour,
		IdentityKey:     Config.EnvConfigs["IDENTITY_KEY"],
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
	if v, ok := data.(*Models.User); ok {
		return jwt.MapClaims{
			Config.EnvConfigs["IDENTITY_KEY"]: v.Email,
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
	var user Models.User
	Config.DB.Model(&Models.User{}).First(&user, "email = ?", claims[Config.EnvConfigs["IDENTITY_KEY"]].(string))
	return &user
}

func handleAuthentication(c *gin.Context) (interface{}, error) {
	var loginVals Models.Login
	if err := c.ShouldBindJSON(&loginVals); err != nil {
		return "", errors.New("missing Email or Password")
	}
	userEmail := loginVals.Email
	password := loginVals.Password
	var user Models.User
	errDB := Config.DB.Model(&Models.User{}).First(&user, "email = ?", userEmail)
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
	if _, ok := data.(*Models.User); ok {
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

func GetLoggedUser(c *gin.Context) *Models.User {
	user, _ := c.Get(Config.EnvConfigs["IDENTITY_KEY"])
	return user.(*Models.User)
}

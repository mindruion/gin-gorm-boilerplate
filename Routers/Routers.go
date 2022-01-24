package Routers

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm-gin/Controllers"
	"gorm-gin/Middlewares"
	_ "gorm-gin/docs"
	"log"
)

func SetupNoAuthRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	v1.POST("/login", authMiddleware.LoginHandler)
	v1.GET("/refresh-token", authMiddleware.RefreshHandler)
}

func SetupAuthRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	v1.Use(authMiddleware.MiddlewareFunc())
	{
		v1.GET("/me", Controllers.Me)
		books := v1.Group("books")
		{
			books.GET("", Controllers.ListBook)
			books.POST("", Controllers.AddNewBook)
			books.GET("/:id", Controllers.GetOneBook)
			books.PUT("/:id", Controllers.PutOneBook)
			books.DELETE("/:id", Controllers.DeleteBook)
		}
	}
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	authMiddleware := Middlewares.InitAuthMiddleware()
	v1 := r.Group("/v1")
	SetupNoAuthRouter(v1, authMiddleware)
	SetupAuthRouter(v1, authMiddleware)
	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	return r
}

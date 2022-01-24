package routers

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm-gin/controllers"
	_ "gorm-gin/docs"
	"gorm-gin/middlewares"
	"log"
)

func SetupNoAuthRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	v1.POST("/login", authMiddleware.LoginHandler)
	v1.GET("/refresh-token", authMiddleware.RefreshHandler)
}

func SetupAuthRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	v1.Use(authMiddleware.MiddlewareFunc())
	{
		v1.GET("/me", controllers.Me)
		books := v1.Group("books")
		{
			books.GET("", controllers.ListBook)
			books.POST("", controllers.AddNewBook)
			books.GET("/:id", controllers.GetOneBook)
			books.PUT("/:id", controllers.PutOneBook)
			books.DELETE("/:id", controllers.DeleteBook)
		}
	}
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	authMiddleware := middlewares.InitAuthMiddleware()
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

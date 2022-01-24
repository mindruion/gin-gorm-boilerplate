package main

import (
	"gorm-gin/config"
	_ "gorm-gin/docs"
	"gorm-gin/routers"
	"gorm-gin/seed"
)

// @title           Boilerplate Gin and Gorm
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /v1

//@securityDefinitions.apikey ApiKeyAuth
//@in header
//@name Authorization
func main() {
	config.LoadEnv()
	config.Init()
	defer config.CloseDB()
	r := routers.SetupRouter()
	seed.Load(config.DB)
	err := r.Run(":8080")
	if err != nil {
		return
	}
}

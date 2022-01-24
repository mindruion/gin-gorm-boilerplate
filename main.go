package main

import (
	"gorm-gin/Config"
	"gorm-gin/Routers"
	"gorm-gin/Seed"
	_ "gorm-gin/docs"
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
	Config.LoadEnv()
	Config.Init()
	defer Config.CloseDB()
	r := Routers.SetupRouter()
	Seed.Load(Config.DB)
	err := r.Run(":8080")
	if err != nil {
		return
	}
}

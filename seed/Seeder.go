package seed

import (
	"gorm-gin/config"
	"gorm-gin/models"
	"log"
)

var books = []models.Book{
	{
		Name: "Start with why",
		Author: models.User{
			Email:    "steven@gmail.com",
			Password: "password",
		},
		Category: "Educational",
	},
	{
		Name: "Psychology of money",
		Author: models.User{
			Email:    "luther@gmail.com",
			Password: "password",
		},
		Category: "Educational",
	},
}

func Load() {
	err := config.DB.Migrator().DropTable(&models.User{}, models.Book{})
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = config.DB.Debug().AutoMigrate(&models.User{}, &models.Book{})
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}
	for i, _ := range books {
		err = config.DB.Debug().Model(&models.Book{}).Create(&books[i]).Error
		if err != nil {
			log.Fatalf("cannot seed books table: %v", err)
		}

	}
}

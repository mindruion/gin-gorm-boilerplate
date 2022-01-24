package seed

import (
	"gorm-gin/models"
	"gorm.io/gorm"
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

func Load(db *gorm.DB) {
	err := db.Migrator().DropTable(&models.User{}, models.Book{})
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Book{})
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}
	for i, _ := range books {
		err = db.Debug().Model(&models.Book{}).Create(&books[i]).Error
		if err != nil {
			log.Fatalf("cannot seed books table: %v", err)
		}

	}
}

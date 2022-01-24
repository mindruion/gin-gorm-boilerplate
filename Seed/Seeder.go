package Seed

import (
	"gorm-gin/Models"
	"gorm.io/gorm"
	"log"
)

var books = []Models.Book{
	{
		Name: "Start with why",
		Author: Models.User{
			Email:    "steven@gmail.com",
			Password: "password",
		},
		Category: "Educational",
	},
	{
		Name: "Psychology of money",
		Author: Models.User{
			Email:    "luther@gmail.com",
			Password: "password",
		},
		Category: "Educational",
	},
}

func Load(db *gorm.DB) {
	err := db.Migrator().DropTable(&Models.User{}, Models.Book{})
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&Models.User{}, &Models.Book{})
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}
	for i, _ := range books {
		err = db.Debug().Model(&Models.Book{}).Create(&books[i]).Error
		if err != nil {
			log.Fatalf("cannot seed books table: %v", err)
		}

	}
}

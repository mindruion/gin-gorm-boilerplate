package models

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm-gin/apiHelpers"
	"gorm-gin/config"
)

func GetAllBook(b *[]Book, pagination *apiHelpers.Pagination, user *User) (err error) {
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := config.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	result := queryBuider.Model(&Book{}).Where("author_id = ?", user.ID).Find(&b)
	if result.Error != nil {
		msg := result.Error
		return msg
	}
	return nil
}

func AddNewBook(b *Book) (err error) {
	if err = config.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOneBook(b *Book, id string, user *User) (err error) {
	if err := config.DB.Where("id = ?", id).Where("author_id = ?", user.ID).First(b).Error; err != nil {
		return err
	}
	return nil
}

func PutOneBook(b *Book, id string) (err error) {
	config.DB.Save(b)
	return nil
}

func DeleteBook(b *Book, id string, user *User) (err error) {
	config.DB.Where("id = ?", id).Where("author_id = ?", user.ID).Delete(b)
	return nil
}

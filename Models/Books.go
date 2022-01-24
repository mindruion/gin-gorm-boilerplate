package Models

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm-gin/ApiHelpers"
	"gorm-gin/Config"
)

func GetAllBook(b *[]Book, pagination *ApiHelpers.Pagination, user *User) (err error) {
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := Config.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	result := queryBuider.Model(&Book{}).Where("author_id = ?", user.ID).Find(&b)
	if result.Error != nil {
		msg := result.Error
		return msg
	}
	return nil
}

func AddNewBook(b *Book) (err error) {
	if err = Config.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOneBook(b *Book, id string, user *User) (err error) {
	if err := Config.DB.Where("id = ?", id).Where("author_id = ?", user.ID).First(b).Error; err != nil {
		return err
	}
	return nil
}

func PutOneBook(b *Book, id string) (err error) {
	Config.DB.Save(b)
	return nil
}

func DeleteBook(b *Book, id string, user *User) (err error) {
	Config.DB.Where("id = ?", id).Where("author_id = ?", user.ID).Delete(b)
	return nil
}

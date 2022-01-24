package Models

type Book struct {
	BaseModel
	Name     string `json:"name" binding:"required"`
	AuthorID uint32 `json:"author" biding:"required"`
	Category string `json:"category"`
	Author   User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
}

func (b *Book) TableName() string {
	return "book"
}

package books

import (
	"bookstore_api/internal/authors"
	"bookstore_api/internal/categories"
	"bookstore_api/internal/database"
)

type Book struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	Title      string `json:"title"`
	Synopsis   string `json:"synopsis"`
	AuthorID   uint   `json:"-"`
	CategoryID uint   `json:"-"`
	Author     authors.Author    `json:"author" gorm:"foreignKey:AuthorID"`
	Category   categories.Category `json:"category" gorm:"foreignKey:CategoryID"`
}

func GetOrCreateAuthor(name string) (authors.Author, error) {
	var author authors.Author
	if err := database.DB.Where("name = ?", name).First(&author).Error; err != nil {
		author = authors.Author{Name: name}
		if err := database.DB.Create(&author).Error; err != nil {
			return authors.Author{}, err
		}
	}
	return author, nil
}

func GetOrCreateCategory(name string) (categories.Category, error) {
	var category categories.Category
	if err := database.DB.Where("name = ?", name).First(&category).Error; err != nil {
		category = categories.Category{Name: name}
		if err := database.DB.Create(&category).Error; err != nil {
			return categories.Category{}, err
		}
	}
	return category, nil
}

package books

import (
	"bookstore_api/internal/authors"
	"bookstore_api/internal/categories"
)

type Book struct {
	ID         uint           `gorm:"primaryKey"`
	Title      string         `gorm:"not null"`
	Synopsis   string         `gorm:"type:text"`
	CategoryID uint           `gorm:"not null"`
	Category   categories.Category `gorm:"foreignKey:CategoryID"`
	AuthorID   uint           `gorm:"not null"`
	Author     authors.Author `gorm:"foreignKey:AuthorID"`
}

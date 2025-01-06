package books

import (
	"bookstore_api/internal/authors"
	"bookstore_api/internal/categories"
)

type Book struct {
	ID         uint           `gorm:"primaryKey"`
	Title      string         `gorm:"not null" binding:"required,unique,min=2,max=100"`
	Synopsis   string         `gorm:"type:text" binding:"required,min=2,max=1000"`
	CategoryID uint           `gorm:"not null"`
	Category   categories.Category `gorm:"foreignKey:CategoryID"`
	AuthorID   uint           `gorm:"not null"`
	Author     authors.Author `gorm:"foreignKey:AuthorID"`
}

package books

import (
	//"bookstore_api/internal/authors"
	//"bookstore_api/internal/categories"
	"bookstore_api/internal/database"
	"errors"
)

type CreateBookInput struct {
	Title    string `json:"title" binding:"required,min=3,max=100"`
	Synopsis string `json:"synopsis" binding:"required,min=10,max=500"`
	Author   struct {
		Name string `json:"name" binding:"required"`
	} `json:"author"`
	Category struct {
		Name string `json:"name" binding:"required"`
	} `json:"category"`
}

func CreateBookService(input CreateBookInput) (Book, error) {
	author, err := GetOrCreateAuthor(input.Author.Name)
	if err != nil {
		return Book{}, errors.New("Erro ao salvar autor")
	}

	category, err := GetOrCreateCategory(input.Category.Name)
	if err != nil {
		return Book{}, errors.New("Erro ao salvar categoria")
	}

	book := Book{
		Title:      input.Title,
		Synopsis:   input.Synopsis,
		AuthorID:   author.ID,
		CategoryID: category.ID,
	}

	if err := database.DB.Create(&book).Error; err != nil {
		return Book{}, errors.New("Erro ao salvar livro")
	}

	return book, nil
}

func GetAllBooksService() ([]Book, error) {
	var books []Book
	if err := database.DB.Preload("Author").Preload("Category").Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func GetBookByIDService(id string) (Book, error) {
	var book Book
	if err := database.DB.Preload("Author").Preload("Category").First(&book, id).Error; err != nil {
		return Book{}, errors.New("Livro não encontrado")
	}
	return book, nil
}

func DeleteBookService(id string) error {
	var book Book
	if err := database.DB.First(&book, id).Error; err != nil {
		return errors.New("Livro não encontrado")
	}

	if err := database.DB.Delete(&book).Error; err != nil {
		return errors.New("Erro ao deletar livro")
	}
	return nil
}

func UpdateBookService(id string, input map[string]interface{}) (Book, error) {
	var book Book
	if err := database.DB.First(&book, id).Error; err != nil {
		return Book{}, errors.New("Livro não encontrado")
	}

	if authorData, ok := input["author"].(map[string]interface{}); ok {
		if name, exists := authorData["name"].(string); exists {
			author, err := GetOrCreateAuthor(name)
			if err != nil {
				return Book{}, errors.New("Erro ao salvar autor")
			}
			book.AuthorID = author.ID
		}
	}

	if categoryData, ok := input["category"].(map[string]interface{}); ok {
		if name, exists := categoryData["name"].(string); exists {
			category, err := GetOrCreateCategory(name)
			if err != nil {
				return Book{}, errors.New("Erro ao salvar categoria")
			}
			book.CategoryID = category.ID
		}
	}

	if title, ok := input["title"].(string); ok {
		book.Title = title
	}
	if synopsis, ok := input["synopsis"].(string); ok {
		book.Synopsis = synopsis
	}

	if err := database.DB.Save(&book).Error; err != nil {
		return Book{}, errors.New("Erro ao atualizar livro")
	}

	return book, nil
}

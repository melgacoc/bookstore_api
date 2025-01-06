package books

import (
	"bookstore_api/internal/authors"
	"bookstore_api/internal/categories"
	"bookstore_api/internal/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	var input struct {
		Title    string `json:"title" binding:"required,min=3,max=100"`
		Synopsis string `json:"synopsis" binding:"required,min=10,max=500"`
		Author   struct {
			Name string `json:"name" binding:"required"`
		} `json:"author"`
		Category struct {
			Name string `json:"name" binding:"required"`
		} `json:"category"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var author authors.Author
	if err := database.DB.Where("name = ?", input.Author.Name).First(&author).Error; err != nil {
		author = authors.Author{Name: input.Author.Name}
		if err := database.DB.Create(&author).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar autor"})
			return
		}
	}

	var category categories.Category
	if err := database.DB.Where("name = ?", input.Category.Name).First(&category).Error; err != nil {
		category = categories.Category{Name: input.Category.Name}
		if err := database.DB.Create(&category).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar categoria"})
			return
		}
	}

	book := Book{
		Title:      input.Title,
		Synopsis:   input.Synopsis,
		AuthorID:   author.ID,
		CategoryID: category.ID,
	}

	if err := database.DB.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar livro"})
		return
	}

	c.JSON(http.StatusCreated, book)
}

func GetAllBooks(c *gin.Context) {
	var books []Book
	if err := database.DB.Preload("Author").Preload("Category").Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar livros"})
		return
	}
	c.JSON(http.StatusOK, books)
}

func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	var book Book
	if err := database.DB.Preload("Author").Preload("Category").First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Livro não encontrado"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	var book Book
	if err := database.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Livro não encontrado"})
		return
	}

	if err := database.DB.Delete(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar livro"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Livro deletado com sucesso"})
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book Book
	if err := database.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Livro não encontrado"})
		return
	}

	var input map[string]interface{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if authorData, ok := input["author"].(map[string]interface{}); ok {
		if name, exists := authorData["name"].(string); exists {
			var author authors.Author
			if err := database.DB.Where("name = ?", name).First(&author).Error; err != nil {
				author = authors.Author{Name: name}
				if err := database.DB.Create(&author).Error; err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar autor"})
					return
				}
			}
			book.AuthorID = author.ID
		}
	}

	if categoryData, ok := input["category"].(map[string]interface{}); ok {
		if name, exists := categoryData["name"].(string); exists {
			var category categories.Category
			if err := database.DB.Where("name = ?", name).First(&category).Error; err != nil {
				category = categories.Category{Name: name}
				if err := database.DB.Create(&category).Error; err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar categoria"})
					return
				}
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar livro"})
		return
	}

	c.JSON(http.StatusOK, book)
}


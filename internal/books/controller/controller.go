package books

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// CreateBookInput represents the input data for creating a new book.
type CreateBookInput struct {
	Title     string `json:"title" binding:"required"`
	Synopsis  string `json:"synopsis" binding:"required"`
	Author    struct {
		Name string `json:"name" binding:"required"`
	} `json:"author" binding:"required"`
	Category  struct {
		Name string `json:"name" binding:"required"`
	} `json:"category" binding:"required"`
}

// CreateBook creates a new book.
// @Summary Create a new book
// @Description Create a new book
// @Tags books
// @Accept  json
// @Produce  json
// @Param input body CreateBookInput true "Book information" example: {"title": "Book 1", "synopsis": "Synopsis of Book 1", "author": {"name": "Author 1"}, "category": {"name": "Category 1"}}
// @Success 201 {object} Book
// @Failure 400 {object} map[string]interface{}
// @Router /books [post]
func CreateBook(c *gin.Context) {
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := CreateBookService(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, book)
}

// GetAllBooks retrieves all books.
// @Summary Get all books
// @Description Get all books
// @Tags books
// @Accept  json
// @Produce  json
// @Success 200 {array} Book
// @Failure 500 {object} map[string]interface{}
// @Router /books [get]
func GetAllBooks(c *gin.Context) {
	books, err := GetAllBooksService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving books"})
		return
	}
	c.JSON(http.StatusOK, books)
}

// GetBookByID retrieves a book by its ID.
// @Summary Get a book by its ID
// @Description Get a book by its ID
// @Tags books
// @Accept  json
// @Produce  json
// @Param id path string true "Book ID"
// @Success 200 {object} Book
// @Failure 404 {object} map[string]interface{}
// @Router /books/{id} [get]
func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	book, err := GetBookByIDService(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

// DeleteBook deletes a book by its ID.
// @Summary Delete a book by its ID
// @Description Delete a book by its ID
// @Tags books
// @Accept  json
// @Produce  json
// @Param id path string true "Book ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]interface{}
// @Router /books/{id} [delete]
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	if err := DeleteBookService(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}

// UpdateBook updates a book by its ID.
// @Summary Update a book by its ID
// @Description Update a book by its ID
// @Tags books
// @Accept  json
// @Produce  json
// @Param id path string true "Book ID"
// @Param input body map[string]interface{} true "Book information" example: {"title": "Book 1", "synopsis": "Synopsis of Book 1", "author": {"name": "Author 1"}, "category": {"name": "Category 1"}}
// @Success 200 {object} Book
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /books/{id} [put]
func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var input map[string]interface{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := UpdateBookService(id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

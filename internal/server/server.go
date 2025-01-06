package server

import (
	"bookstore_api/internal/books"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()

	router.GET("/api/books", books.GetAllBooks)
	router.GET("/api/books/:id", books.GetBookByID)
	router.POST("/api/books", books.CreateBook)
	router.PUT("/api/books/:id", books.UpdateBook)
	router.DELETE("/api/books/:id", books.DeleteBook)

	router.Run(":8080")
}

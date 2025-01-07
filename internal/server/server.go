package server

import (
	"github.com/gin-gonic/gin"
	"bookstore_api/internal/books"
)

func SetupRoutes(router *gin.RouterGroup) {
	router.GET("/books", books.GetAllBooks)
	router.GET("/books/:id", books.GetBookByID)
	router.POST("/books", books.CreateBook)
	router.PUT("/books/:id", books.UpdateBook)
	router.DELETE("/books/:id", books.DeleteBook)
	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}

func StartServer() {
	router := gin.Default()
	api := router.Group("/api")
	SetupRoutes(api)
	router.Run(":8080")
}

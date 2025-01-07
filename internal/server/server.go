package server

import (
	"github.com/gin-gonic/gin"
	"bookstore_api/internal/books"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(router *gin.RouterGroup) {
	router.GET("/books", books.GetAllBooks)
	router.GET("/books/:id", books.GetBookByID)
	router.POST("/books", books.CreateBook)
	router.PUT("/books/:id", books.UpdateBook)
	router.DELETE("/books/:id", books.DeleteBook)
}

func StartServer(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := r.Group("/api")
	SetupRoutes(api)
	r.Run(":8080")
}

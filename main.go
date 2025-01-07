package main

import (
    "bookstore_api/internal/authors"
    "bookstore_api/internal/books"
    "bookstore_api/internal/categories"
    "bookstore_api/internal/database"
    "bookstore_api/internal/server"
    "log"
    _ "bookstore_api/docs"
    "github.com/gin-gonic/gin"
)

// @title Bookstore API
// @version 1.0
// @description An CRUD to manage a bookstore.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

func main() {
    database.ConnectDatabase()

    err := database.DB.AutoMigrate(&books.Book{}, &authors.Author{}, &categories.Category{})
    if err != nil {
        log.Fatal("Erro ao migrar o banco de dados:", err)
    }

    log.Println("Tabelas criadas com sucesso!")

    r := gin.Default()

    server.StartServer(r)
}
package main

import (
	"bookstore_api/internal/authors"
	"bookstore_api/internal/books"
	"bookstore_api/internal/categories"
	"bookstore_api/internal/database"
	"bookstore_api/internal/server"
	"log"
)

func main() {
	database.ConnectDatabase()

	err := database.DB.AutoMigrate(&books.Book{}, &authors.Author{}, &categories.Category{})
	if err != nil {
		log.Fatal("Erro ao migrar o banco de dados:", err)
	}

	log.Println("Tabelas criadas com sucesso!")

	server.StartServer()
}

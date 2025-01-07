package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"bookstore_api/internal/books"
	"bookstore_api/internal/database"
	"bookstore_api/internal/server"
)


func TestMain(m *testing.M) {
	buildTesteDB()
	defer database.DB.Migrator().DropTable(&books.Book{})

	exitCode := m.Run()

	os.Exit(exitCode)
}

func buildTesteDB() {
	database.ConnectTestDatabase()
    database.DB.Migrator().DropTable(&books.Book{})
    database.DB.AutoMigrate(&books.Book{})
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	server.SetupRoutes(router.Group("/api"))
	return router
}

func TestCreateBook(t *testing.T) {
    buildTesteDB()
    router := setupRouter()

    payload := map[string]interface{}{
        "title":    "Harry Potter e a Câmara Secreta",
        "synopsis": "Harry Potter descobre que é um bruxo e vai para Hogwarts",
        "author": map[string]interface{}{
            "name": "J.K. Rowling",
        },
        "category": map[string]interface{}{
            "name": "Fantasia",
        },
    }

    body, _ := json.Marshal(payload)
    req, _ := http.NewRequest(http.MethodPost, "/api/books", bytes.NewBuffer(body))
    req.Header.Set("Content-Type", "application/json")
    rec := httptest.NewRecorder()

    router.ServeHTTP(rec, req)

    assert.Equal(t, http.StatusCreated, rec.Code)
}

func TestCreateBookWithInvalidPayload(t *testing.T) {
	buildTesteDB()
	router := setupRouter()

	payload := map[string]interface{}{
		"title": "Harry Potter e a Câmara Secreta",
	}

	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPost, "/api/books", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	t.Logf("Response body: %s", rec.Body.String())

	assert.Equal(t, rec.Code, http.StatusBadRequest)
}

func TestGetAllBooks(t *testing.T) {
	buildTesteDB()
	router := setupRouter()

	req, _ := http.NewRequest(http.MethodGet, "/api/books", nil)
	rec := httptest.NewRecorder()
	
	router.ServeHTTP(rec, req)
	t.Logf("Response body: %s", rec.Body.String())

	assert.Equal(t, rec.Code, http.StatusOK)
}


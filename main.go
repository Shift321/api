package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

// Return all books
func getBooks(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, books)
}

// Return book with current id
func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

// Checkout book
func checkoutBook(c *gin.Context) {
	book := checkParams(c)
	if book == nil {
		return
	}
	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

// Check if all params ok
func checkParams(c *gin.Context) *book {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parammetr"})
		return nil
	}
	book, err := getBookById(id)
	info := fmt.Sprintf("Book with id %s dont exit", id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": info})
		return nil
	}
	return book
}
// Return book back
func returnBook(c *gin.Context) {
	book := checkParams(c)
	if book == nil {
		return
	}
	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}

// Check if book with current id exists
func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil

		}
	}
	return nil, errors.New("book not found")
}

// Create a book
func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

// Main func create a router
func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.PATCH("/checkout", checkoutBook)
	router.PATCH("/return", returnBook)
	router.GET("/books/:id", bookById)
	router.Run("localhost:8080")
}

# Simple API on GO
API handling books using [gin](https://github.com/gin-gonic/gin)

## Book structure
Struct to keep data of books
```go
type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}
```

## Main function to start API
Createing router using gin
```go
func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.Run("localhost:8080")
}
```

## Create books
Create a book and adds it to struct
```go
func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}
```
## Get books
Get books using gin context
```go
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}
```
## Check if book exists
Check that book with id exists
```go
func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil

		}
	}
	return nil, errors.New("book not found")
}
```

## Get book by id
Get book by id if not found error 404
```go
func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}
```
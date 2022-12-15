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
## Checkout book
Checkout book that reduce quantity by 1
```go
func checkoutBook(c *gin.Context) {
	book := checkParams(c)
	if book == nil {
		return
	}
	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}
```
## Return book
Rturn book that increase quantity by 1
```go
func returnBook(c *gin.Context) {
	book := checkParams(c)
	if book == nil {
		return
	}
	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}
```
## Util funcs

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
Check params
```go
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
```

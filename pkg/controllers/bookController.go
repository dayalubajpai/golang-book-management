package controllers

import (
	"database/sql"
	"dayalubajpai/pkg/config"
	"dayalubajpai/pkg/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBooks() gin.HandlerFunc {
	return func(c *gin.Context) {

		var book models.Book
		// Query the database to get all books
		rows, err := config.DB.Query("SELECT id, title, author FROM books")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong while fetching data from database"})
			return
		}
		defer rows.Close()
		// Create a slice to hold the books
		var books []models.Book
		// Iterate through the rows and scan each book into the slice
		for rows.Next() {
			err := rows.Scan(&book.ID, &book.Title, &book.Author)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong while fetching data from database"})
				return
			}
			books = append(books, book)
		}
		// Check for any errors encountered during iteration
		if err := rows.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong while fetching data from database"})
			return
		}
		// Return the list of books as JSON
		if len(books) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "No books found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "List of books",
			"books":   books,
		})

	}
}

func CreateBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		var book models.Book
		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result, err := config.DB.Exec("INSERT INTO books (id, title, author) VALUES (?, ?, ?)", book.ID, book.Title, book.Author)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong while adding data on database"})
		}

		rowsAffected, err := result.RowsAffected()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong while adding data on database"})
		}

		if rowsAffected > 0 {
			c.JSON(200, gin.H{
				"message":      "Book created within database",
				"rowsAffected": rowsAffected,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong while adding data on database"})
		}

	}
}

func GetBookById() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Get the book ID from the URL parameter
		bookID := c.Param("id")

		// Query the database to get the book by ID
		var book models.Book
		err := config.DB.QueryRow("SELECT id, title, author FROM books WHERE id = ?", bookID).Scan(&book.ID, &book.Title, &book.Author)

		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"}) // If no rows were returned, return a 404 error database sql function
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong while fetching data from database"})
			return
		}
		// Return the book details as JSON

		c.JSON(200, gin.H{
			"message":     "Get book by ID - " + bookID,
			"book author": book.Author,
			"book title":  book.Title,
			"book id":     book.ID,
		})
	}
}

func UpdateBook() gin.HandlerFunc {
	return func(c *gin.Context) {

		bookId := c.Param("id")

		var books models.Book
		err := c.ShouldBindJSON(&books)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Please check json with valid formate { 'title': '', 'author': ' }"})
			return
		}

		_, err = config.DB.Exec("UPDATE `books` SET `title`= ?,`author`= ? WHERE `id` = ?", &books.Author, &books.Title, bookId)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusFound, gin.H{
			"message": fmt.Sprintf("Record Updated - %v", bookId),
		})
	}
}

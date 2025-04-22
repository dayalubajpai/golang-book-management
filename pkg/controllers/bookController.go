package controllers

import (
	"dayalubajpai/pkg/config"
	"dayalubajpai/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBooks() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "List of books",
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

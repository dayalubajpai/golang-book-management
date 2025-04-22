package routes

import (
	"dayalubajpai/pkg/controllers"

	"github.com/gin-gonic/gin"
)

var RegisterBookRoutes = func(importRoutes *gin.Engine) {
	importRoutes.GET("/book", controllers.GetBooks())
	importRoutes.POST("/book", controllers.CreateBook())
	// importRoutes.GET("/book/:id", controllers.GetBookById)
	// importRoutes.PUT("/book/:id", controllers.UpdateBook)
	// importRoutes.DELETE("/book/:id", controllers.DeleteBook)
}

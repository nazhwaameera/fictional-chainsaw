package routers

import (
	"Challenge_6/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/books", controllers.GetAllBooks)

	router.GET("/books/:bookID", controllers.GetBookbyID)

	router.POST("/books", controllers.CreateBook)

	router.PUT("/books/:bookID", controllers.UpdateBookbyID)

	router.DELETE("/books/:bookID", controllers.DeleteBookbyID)

	return router
}

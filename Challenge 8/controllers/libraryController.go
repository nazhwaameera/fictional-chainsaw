package controllers

import (
	"fmt"
	"net/http"

	"Challenge_8/models"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// fungsi untuk memasukkan data buku baru
func (c *Controllers) CreateBook(ctx *gin.Context) {

	var newBook models.Book

	if err := ctx.ShouldBindWith(&newBook, binding.JSON); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newBookData := models.Book{Name_book: newBook.Name_book, Author: newBook.Author}
	err := c.masterDB.Create(&newBookData).Error

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"book": newBookData,
	})
}

// fungsi untuk mendapatkan data buku berdasarkan ID
func (c *Controllers) GetBookbyID(ctx *gin.Context) {
	bookID := ctx.Param("bookID")

	book := models.Book{}

	err := c.masterDB.First(&book, "id = ?", bookID).Error

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": book,
	})
}

// Fungsi untuk mengambil data semua buku
func (c *Controllers) GetAllBooks(ctx *gin.Context) {
	books := []models.Book{}

	err := c.masterDB.Find(&books).Error

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if len(books) == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("There is no book data saved."),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": books,
	})
}

// fungsi untuk mengupdate data buku
func (c *Controllers) UpdateBookbyID(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	book := models.Book{}

	err1 := c.masterDB.Where("id = ?", bookID).First(&book).Error

	if err1 != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err1)
		return
	}

	var input models.Book

	if err := ctx.ShouldBindWith(&input, binding.JSON); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := c.masterDB.Model(&book).Updates(input).Error

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": book,
	})
}

// fungsi untuk menghapus data buku
func (c *Controllers) DeleteBookbyID(ctx *gin.Context) {
	bookID := ctx.Param("bookID")

	book := models.Book{}

	err := c.masterDB.Where("id = ?", bookID).Delete(&book).Error

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book deleted successfully."),
	})
}

package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Book struct {
	BookID string `json:"book_id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Copies string `string:"copies"`
}

var BookDatas = []Book{}

// fungsi untuk mendapatkan data buku (keseluruhan)
func GetAllBooks(ctx *gin.Context) {
	if len(BookDatas) == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("There is no book data saved."),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": BookDatas,
	})
}

// fungsi untuk mendapatkan data buku berdasarkan ID
func GetBookbyID(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	condition := false

	var bookData Book

	for i, book := range BookDatas {
		if bookID == book.BookID {
			condition = true
			bookData = BookDatas[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Book with ID %v not found.", bookID),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": bookData,
	})
}

// fungsi untuk memasukkan data buku baru
func CreateBook(ctx *gin.Context) {
	var newBook Book

	if err := ctx.ShouldBindWith(&newBook, binding.JSON); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newBook.BookID = fmt.Sprintf("c%d", len(BookDatas)+1)
	BookDatas = append(BookDatas, newBook)

	ctx.JSON(http.StatusCreated, gin.H{
		"book": newBook,
	})
}

// fungsi untuk mengupdate data buku
func UpdateBookbyID(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	condition := false

	var updatedBook Book

	if err := ctx.ShouldBindWith(&updatedBook, binding.JSON); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, book := range BookDatas {
		if bookID == book.BookID {
			condition = true
			BookDatas[i] = updatedBook
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Book with ID %v not found.", bookID),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with ID %v has been successfully updated.", bookID),
		"book":    updatedBook,
	})
}

// fungsi untuk menghapus data buku
func DeleteBookbyID(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	condition := false

	var bookIndex int

	for i, book := range BookDatas {
		if bookID == book.BookID {
			condition = true
			bookIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Book with ID %v not found.", bookID),
		})

		return
	}

	copy(BookDatas[bookIndex:], BookDatas[bookIndex+1:])
	BookDatas[len(BookDatas)-1] = Book{}
	BookDatas = BookDatas[:len(BookDatas)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with ID %v has been successfully deleted.", bookID),
	})
}

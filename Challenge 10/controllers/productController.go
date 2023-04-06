package controllers

import (
	"Challenge_10/database"
	"Challenge_10/helpers"
	"Challenge_10/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// function to create a new product's data.
func CreateProduct(c *gin.Context) {
	db := database.GetDB()
	// merupakan method yang digunakan untuk mendapatkan claim dari token yang telah disimpan oleh middleware authentication.
	// method ini mengembalikan empty interface sehingga perlu dilakukan casting untuk menggembalikan ke tipe data jwt.MapClaims.
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Product := models.Product{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	Product.UserID = userID

	err := db.Debug().Create(&Product).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, Product)
}

// function to update existing product's data.
func UpdateProduct(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Product := models.Product{}

	productId, _ := strconv.Atoi(c.Param("productId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	Product.UserID = userID
	Product.ID = uint(productId)

	err := db.Model(&Product).Where("id = ?", productId).First(&Product).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Data Not Found",
			"message": "Data doesn't exist",
		})

		return
	}

	err1 := db.Model(&Product).Updates(models.Product{Title: Product.Title, Description: Product.Description}).Error

	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err1.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, Product)
}

// function to get product data by ID
func GetProductbyID(c *gin.Context) {
	db := database.GetDB()
	Product := models.Product{}

	productId, _ := strconv.Atoi(c.Param("productId"))

	err := db.First(&Product, "id = ?", productId).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "Data Not Found",
			"message": "Data doesn't exist",
		})

		return
	}

	c.JSON(http.StatusOK, Product)
}

// function to delete product data by ID
func DeleteProduct(c *gin.Context) {
	db := database.GetDB()
	Product := models.Product{}

	productId, _ := strconv.Atoi(c.Param("productId"))

	err := db.Where("id = ?", productId).Delete(&Product).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data has been successfully deleted.",
	})
}

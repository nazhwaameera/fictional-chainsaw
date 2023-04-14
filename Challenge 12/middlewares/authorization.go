package middlewares

import (
	"Challenge_12/database"
	"Challenge_12/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func SocialMediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()

		// mendapatkan route parameter berupa productId.
		socialmediaId, err := strconv.Atoi(c.Param("socialmediaId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Invalid parameter.",
			})

			return
		}

		// melakukan claim token yang telah disimpan oleh proses autentikasi.
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		SocialMedia := models.SocialMedia{}

		// mendapatkan data berdasarkan productId yang didapatkan melalui param
		err = db.Select("user_id").First(&SocialMedia, uint(socialmediaId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "Data doesn't exist",
			})

			return
		}

		if SocialMedia.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data.",
			})

			return
		}

		c.Next()
	}
}

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()

		// mendapatkan route parameter berupa productId.
		photoId, err := strconv.Atoi(c.Param("photoId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Invalid parameter.",
			})

			return
		}

		// melakukan claim token yang telah disimpan oleh proses autentikasi.
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Photo := models.Photo{}

		// mendapatkan data berdasarkan productId yang didapatkan melalui param
		err = db.Select("user_id").First(&Photo, uint(photoId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "Data doesn't exist",
			})

			return
		}

		if Photo.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data.",
			})

			return
		}

		c.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()

		// mendapatkan route parameter berupa productId.
		commentId, err := strconv.Atoi(c.Param("commentId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Invalid parameter.",
			})

			return
		}

		// melakukan claim token yang telah disimpan oleh proses autentikasi.
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Comment := models.Comment{}

		// mendapatkan data berdasarkan productId yang didapatkan melalui param
		err = db.Select("user_id").First(&Comment, uint(commentId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "Data doesn't exist",
			})

			return
		}

		if Comment.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data.",
			})

			return
		}

		c.Next()
	}
}

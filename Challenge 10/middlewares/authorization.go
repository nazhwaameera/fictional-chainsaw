package middlewares

import (
	"Challenge_10/database"
	"Challenge_10/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ProductAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()

		// mendapatkan route parameter berupa productId.
		productId, err := strconv.Atoi(c.Param("productId"))

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
		userRole := string(userData["role"].(string))
		Product := models.Product{}

		// mendapatkan data berdasarkan productId yang didapatkan melalui param
		err = db.Select("user_id").First(&Product, uint(productId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "Data doesn't exist",
			})

			return
		}

		if Product.UserID != userID {
			if userRole == "admin" {
				c.Next()
			} else {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error":   "Unauthorized",
					"message": "You are not allowed to access this data.",
				})

			}

			return
		}

		c.Next()
	}
}

func AdminAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {

		// melakukan claim token yang telah disimpan oleh proses autentikasi.
		userData := c.MustGet("userData").(jwt.MapClaims)
		userRole := string(userData["role"].(string))

		if userRole != "admin" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data.",
			})

			return
		}

		c.Next()
	}
}

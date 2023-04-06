package middlewares

import (
	"Challenge_10/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthenticated",
				"message": err.Error(),
			})

			return
		}

		// Menyimpan claim dari token ke dalam data request agar dapat diambil pada endpoint berikutnya.
		c.Set("userData", verifyToken)
		// method agar dapat melanjutkan proses ke endpoint berikutnya.
		c.Next()
	}
}

package helpers

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey = "password"

// jwt.mapClaims digunakan untuk menyimpan data user yang terdapat dalam token JWT.
func GenerateToken(id uint, email string) string {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
	}

	// method jwt.NewWithClaims digunakan untuk emmasukkan data-data user yang telah disimpan pada jwt.mapClaims.
	// kita juga menentukan metode enskripsi data tersebut di sini.
	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// method jwt.SignedString merupakan method untuk mem-parsing token menjadi sebuah string panjang
	signedToken, _ := parseToken.SignedString([]byte(secretKey))

	return signedToken
}

func VerifyToken(c *gin.Context) (interface{}, error) {
	errResponse := errors.New("Sign in to proceed")
	// token pada client di simpan dalam headers di variabel Authorization
	headerToken := c.Request.Header.Get("Authorization")
	// token yang tidak memiliki prefix bearer akan dianggap tidak valid.
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, errResponse
	}

	// mengambil token tanpa prefix Bearernya
	stringToken := strings.Split(headerToken, " ")[1]

	// melakukan parsing token menjadi sebuah struct pointer dan memeriksa metode enskripsi tokennya dengan mengcasting metodenya menjadi tipe data pointer dari struct jwt.SigningMehodHMAC
	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}

		return []byte(secretKey), nil
	})

	// mengcasting claim token menjadi tipe data jwt.mMpClaims.
	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errResponse
	}

	// me-return claim dari token
	return token.Claims.(jwt.MapClaims), nil
}

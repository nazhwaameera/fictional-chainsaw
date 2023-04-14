package controllers

import (
	"Challenge_12/database"
	"Challenge_12/helpers"
	"Challenge_12/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

//	UserRegister godoc
//
// @Summary Create user
// @Description Create new user data
// @Tags json
// @Accept json
// @Produce json
// @Param models.Comment body models.User true "create user"
// @Success 200 {object} models.User
// @Router /user/register [post]
func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType

	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":       User.ID,
		"username": User.Username,
		"email":    User.Email,
	})
}

//	UserLogin godoc
//
// @Summary Login as a user
// @Description Login as a user using existing data
// @Tags json
// @Accept json
// @Produce json
// @Param models.Comment body models.User true "login user"
// @Success 200 {object} models.User
// @Router /user/login [post]
func UserLogin(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType

	User := models.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email/password",
		})
	}

	comparePass := helpers.PasswordValid(User.Password, password)

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email/password",
		})

		return
	}

	token := helpers.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

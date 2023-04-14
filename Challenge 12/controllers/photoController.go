package controllers

import (
	"Challenge_12/database"
	"Challenge_12/helpers"
	"Challenge_12/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//	CreatePhoto godoc
//
// @Summary Create photo
// @Description Create new photo data
// @Tags json
// @Accept json
// @Produce json
// @Param models.Comment body models.Photo true "create photo"
// @Success 200 {object} models.Photo
// @Router /photo/create [post]
func CreatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID

	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, Photo)
}

// UpdatePhoto godoc
// @Summary Update photo
// @Description Update existing photo's data.
// @Tags json
// @Accept json
// @Produce json
// @Param Id path int true "ID of the photo"
// @Success 200 {object} models.Photo
// @Router /photo/update/{Id} [patch]
func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))

	err := db.Where("id = ?", photoId).First(&Photo).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Data Not Found",
			"message": "Photo doesn't exist",
		})

		return
	}

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID
	Photo.ID = uint(photoId)

	err1 := db.Model(&Photo).Updates(models.Photo{Title: Photo.Title, URL: Photo.URL, Caption: Photo.Caption}).Error

	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err1.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, Photo)
}

// GetPhotoByID godoc
// @Summary Get photo by photo id
// @Description Get details of specific photo
// @Tags json
// @Accept json
// @Produce json
// @Param Id path int true "ID of the photo"
// @Success 200 {object} models.Photo
// @Router /photo/getone/{Id} [get]
func GetPhotobyID(c *gin.Context) {
	db := database.GetDB()
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))

	err := db.First(&Photo, "id = ?", photoId).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "Data Not Found",
			"message": "Photo doesn't exist",
		})

		return
	}

	c.JSON(http.StatusOK, Photo)
}

//	GetAllPhotos godoc
//
// @Summary Get all photos
// @Description Get details of all photos data
// @Tags json
// @Accept json
// @Produce json
// @Success 200 {object} models.Photo
// @Router /photo/getall [get]
func GetAllPhotos(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	Photos := []models.Photo{}

	userID := uint(userData["id"].(float64))

	err := db.Find(&Photos, "user_id = ?", userID).Error

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if len(Photos) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "Data Not Found",
			"message":      "This user doesn't have any photo.",
		})

		return
	}

	c.JSON(http.StatusOK, Photos)
}

// DeletePhoto godoc
// @Summary Delete photo
// @Description Delete photo data
// @Tags json
// @Accept json
// @Produce json
// @Param Id path int true "ID of the photo"
// @Success 200
// @Router /photo/delete/{Id} [delete]
func DeletePhoto(c *gin.Context) {
	db := database.GetDB()
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))

	err := db.Where("id = ?", photoId).Delete(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data has been successfully deleted.",
	})
}

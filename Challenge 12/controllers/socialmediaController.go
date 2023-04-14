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

//	CreateSocialMedia godoc
//
// @Summary Create social media
// @Description Create new social media data
// @Tags json
// @Accept json
// @Produce json
// @Param models.Comment body models.SocialMedia true "create social media"
// @Success 200 {object} models.Photo
// @Router /socialmedia/create [post]
func CreateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	SocialMedia := models.SocialMedia{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID

	err := db.Debug().Create(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, SocialMedia)
}

// UpdateSocialMedia godoc
// @Summary Update social media
// @Description Update existing social media's data.
// @Tags json
// @Accept json
// @Produce json
// @Param Id path int true "ID of the social media"
// @Success 200 {object} models.SocialMedia
// @Router /socialmedia/update/{Id} [patch]
func UpdateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	SocialMedia := models.SocialMedia{}

	socialmediaId, _ := strconv.Atoi(c.Param("socialmediaId"))
	userID := uint(userData["id"].(float64))

	err := db.Where("id = ?", socialmediaId).First(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Data Not Found",
			"message": "Data doesn't exist",
		})

		return
	}

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID
	SocialMedia.ID = uint(socialmediaId)

	err1 := db.Model(&SocialMedia).Updates(models.SocialMedia{Name: SocialMedia.Name, URL: SocialMedia.URL}).Error

	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err1.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, SocialMedia)
}

// GetSocialMediaByID godoc
// @Summary Get social media by id
// @Description Get details of specific social media
// @Tags json
// @Accept json
// @Produce json
// @Param Id path int true "ID of the social media"
// @Success 200 {object} models.SocialMedia
// @Router /socialmedia/getone/{Id} [get]
func GetSocialMediabyID(c *gin.Context) {
	db := database.GetDB()
	SocialMedia := models.SocialMedia{}

	socialmediaId, _ := strconv.Atoi(c.Param("socialmediaId"))

	err := db.First(&SocialMedia, "id = ?", socialmediaId).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "Data Not Found",
			"message": "Data doesn't exist",
		})

		return
	}

	c.JSON(http.StatusOK, SocialMedia)
}

//	GetAllSocialMedias godoc
//
// @Summary Get all social medias
// @Description Get details of all social media datas
// @Tags json
// @Accept json
// @Produce json
// @Success 200 {object} models.SocialMedia
// @Router /socialmedia/getall [get]
func GetAllSocialMedias(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	SocialMedias := []models.SocialMedia{}

	userID := uint(userData["id"].(float64))

	err := db.Find(&SocialMedias, "user_id = ?", userID).Error

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if len(SocialMedias) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "Data Not Found",
			"message":      "This user doesn't have any social media.",
		})

		return
	}

	c.JSON(http.StatusOK, SocialMedias)
}

// DeleteSocialMedia godoc
// @Summary Delete social media
// @Description Delete social media data
// @Tags json
// @Accept json
// @Produce json
// @Param Id path int true "ID of the social media"
// @Success 200
// @Router /socialmedia/delete/{Id} [delete]
func DeleteSocialMedia(c *gin.Context) {
	db := database.GetDB()
	SocialMedia := models.SocialMedia{}

	socialmediaId, _ := strconv.Atoi(c.Param("socialmediaId"))

	err := db.Where("id = ?", socialmediaId).Delete(&SocialMedia).Error

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

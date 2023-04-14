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

//	CreateComment godoc
//
// @Summary Create comment
// @Description Create new comment data
// @Tags json
// @Accept json
// @Produce json
// @Param models.Comment body models.Comment true "create comment"
// @Success 200 {object} models.Comment
// @Router /photo/{photoId}/comment/create [post]
func CreateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	Comment := models.Comment{}
	userID := uint(userData["id"].(float64))
	photoId, _ := strconv.Atoi(c.Param("photoId"))

	err := db.Where("id = ?", photoId).First(&Photo).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Data Not Found",
			"message": "Photo doesn't exist",
		})

		return
	}

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID
	Comment.PhotoID = uint(photoId)

	err1 := db.Debug().Create(&Comment).Error

	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err1.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, Comment)
}

// UpdateComment godoc
// @Summary Update comment
// @Description Update existing comment's data.
// @Tags json
// @Accept json
// @Produce json
// @Param Id path int true "ID of the comment"
// @Success 200 {object} models.Comment
// @Router /photo/{photoId}/comment/update/{Id} [patch]
func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userID := uint(userData["id"].(float64))

	err := db.Where("id = ?", photoId).First(&Photo).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Data Not Found",
			"message": "Photo doesn't exist",
		})

		return
	}

	err1 := db.Where("photo_id = ?", photoId).Where("id = ?", commentId).First(&Comment).Error

	if err1 != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Data Not Found",
			"message": "Comment doesn't exist",
		})

		return
	}

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID
	Comment.ID = uint(commentId)
	Comment.PhotoID = uint(photoId)

	err2 := db.Model(&Comment).Updates(models.Comment{Message: Comment.Message}).Error

	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err2.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, Comment)
}

// GetCommentByID godoc
// @Summary Get comment by photo and comment id
// @Description Get details of specific comment
// @Tags json
// @Accept json
// @Produce json
// @Param Id path int true "ID of the comment and photo"
// @Success 200 {object} models.Comment
// @Router /photo/{photoId}/comment/getone/{Id} [get]
func GetCommentbyID(c *gin.Context) {
	db := database.GetDB()
	Comment := models.Comment{}
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	commentId, _ := strconv.Atoi(c.Param("commentId"))

	err := db.Where("id = ?", photoId).First(&Photo).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Data Not Found",
			"message": "Photo doesn't exist",
		})

		return
	}

	err1 := db.Where("photo_id = ?", photoId).Where("id = ?", commentId).First(&Comment).Error

	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "Data Not Found",
			"message": "Comment doesn't exist",
		})

		return
	}

	c.JSON(http.StatusOK, Comment)
}

//	GetAllComments godoc
//
// @Summary Get all comments
// @Description Get details of all comments data
// @Tags json
// @Accept json
// @Produce json
// @Success 200 {object} models.Comment
// @Router /photo/{photoId}/comment/getall [get]
func GetAllComments(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	Comments := []models.Comment{}

	userID := uint(userData["id"].(float64))

	err := db.Find(&Comments, "user_id = ?", userID).Error

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if len(Comments) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "Data Not Found",
			"message":      "This user doesn't have any comment.",
		})

		return
	}

	c.JSON(http.StatusOK, Comments)
}

// DeleteComment godoc
// @Summary Delete comment
// @Description Delete comment data
// @Tags json
// @Accept json
// @Produce json
// @Param Id path int true "ID of the comment and photo"
// @Success 200
// @Router /photo/{photoId}/comment/delete/{Id} [delete]
func DeleteComment(c *gin.Context) {
	db := database.GetDB()
	Comment := models.Comment{}
	Photo := models.Photo{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	photoId, _ := strconv.Atoi(c.Param("photoId"))

	err := db.Where("id = ?", photoId).First(&Photo).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Data Not Found",
			"message": "Photo doesn't exist",
		})

		return
	}

	err1 := db.Where("photo_id = ?", photoId).Where("id = ?", commentId).Delete(&Comment).Error

	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err1.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data has been successfully deleted.",
	})
}

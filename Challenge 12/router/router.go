package router

import (
	"Challenge_12/controllers"
	"Challenge_12/middlewares"

	"github.com/gin-gonic/gin"

	_ "Challenge_12/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Stargram API
// @version 1.0
// @description Service to manage user's social media data
// @termsOfService https://google.com
// @contact.name API Support
// @contact.email admin@mail.me
// @lisence.name Apache 2.0
// @lisence.url https://google.com
// @host localhost:8080
// @BasePath /
func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		// Register
		userRouter.POST("/register", controllers.UserRegister)

		// Login
		userRouter.POST("/login", controllers.UserLogin)
	}

	socialmediaRouter := r.Group("/socialmedia")
	{

		socialmediaRouter.Use(middlewares.Authentication())

		// Create Social Media
		socialmediaRouter.POST("/create", controllers.CreateSocialMedia)

		// Read All Social Media
		socialmediaRouter.GET("/getall", controllers.GetAllSocialMedias)

		// Read Social Media
		socialmediaRouter.GET("/getone/:socialmediaId", middlewares.SocialMediaAuthorization(), controllers.GetSocialMediabyID)

		// Update Social Media
		socialmediaRouter.PUT("/update/:socialmediaId", middlewares.SocialMediaAuthorization(), controllers.UpdateSocialMedia)

		// Delete Social Media
		socialmediaRouter.DELETE("/delete/:socialmediaId", middlewares.SocialMediaAuthorization(), controllers.DeleteSocialMedia)
	}

	photoRouter := r.Group("/photo")
	{
		photoRouter.Use(middlewares.Authentication())

		// Create Photo
		photoRouter.POST("/create", controllers.CreatePhoto)

		// Read All Photos
		photoRouter.GET("/getall", controllers.GetAllPhotos)

		// Read Photo
		photoRouter.GET("/getone/:photoId", middlewares.PhotoAuthorization(), controllers.GetPhotobyID)

		// Update Photo
		photoRouter.PUT("/update/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)

		// Delete Photo
		photoRouter.DELETE("/delete/:photoId", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}

	commentRouter := r.Group("/photo/:photoId/comment")
	{
		commentRouter.Use(middlewares.Authentication())

		// Create Comment
		commentRouter.POST("/create", controllers.CreateComment)

		// Read All Comments
		commentRouter.GET("/getall", controllers.GetAllComments)

		// Read Comment
		commentRouter.GET("/getone/:commentId", middlewares.PhotoAuthorization(), controllers.GetCommentbyID)

		// Update Comment
		commentRouter.PUT("/update/:commentId", middlewares.PhotoAuthorization(), controllers.UpdateComment)

		// Delete Comment
		commentRouter.DELETE("/delete/:commentId", middlewares.PhotoAuthorization(), controllers.DeleteComment)

	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

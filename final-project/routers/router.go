package routers

import (
	"final-project/controllers"
	"final-project/db"
	"final-project/middleware"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	db := db.ConnectDB()
	userController := controllers.NewUserController(db)
	photoController := controllers.NewPhotoController(db)
	commentController := controllers.NewCommentController(db)
	socialMediaController := controllers.NewSocialMediaController(db)

	// router.Use(middleware.Auth())
	userGroup := router.Group("/users")
	{
		userGroup.POST("/register", userController.Register)
		userGroup.POST("/login", userController.Login)
		userGroup.Use(middleware.Auth())
		userGroup.PUT("/", userController.UpdateUser)
		userGroup.DELETE("/", userController.DeleteUser)
	}

	photoGroup := router.Group("/photos")
	{

		photoGroup.Use(middleware.Auth())
		photoGroup.GET("/", photoController.GetPhotos)
		photoGroup.POST("/", photoController.CreatePhoto)
		photoGroup.PUT("/:photoId", photoController.UpdatePhoto)
		photoGroup.DELETE("/:photoId", photoController.UpdatePhoto)
	}

	commentGroup := router.Group("/comments")
	{

		// commentGroup.Use(middleware.Auth())
		commentGroup.GET("/", commentController.GetComments)
		commentGroup.POST("/", commentController.PostComment)
		commentGroup.PUT("/:commentId", commentController.UpdateComment)
		commentGroup.DELETE("/:commentId", commentController.UpdateComment)
	}

	socialMediaGroup := router.Group("/socialmedias")
	{

		// socialMediaGroup.Use(middleware.Auth())
		socialMediaGroup.GET("/", socialMediaController.GetSocialMedias)
		socialMediaGroup.POST("/", socialMediaController.CreateSocialMedia)
		socialMediaGroup.PUT("/:socialMediaId", socialMediaController.UpdateSocialMedia)
		socialMediaGroup.DELETE("/:socialMediaId", socialMediaController.UpdateSocialMedia)
	}

	return router
}

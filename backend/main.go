package main

import (
	"net/http"

	"scalable-final-proj/backend/controllers"
	"scalable-final-proj/backend/middlewares"
	"scalable-final-proj/backend/models"
	"scalable-final-proj/backend/product_svc/p_controllers"

	"github.com/gin-gonic/gin"
)


func main() {

	models.ConnectDataBase()

	router := gin.Default()

	public := router.Group("/api")

	public.POST("/signup", controllers.Register)
	public.POST("/login", controllers.Login)
	protected := router.Group("/api/user")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.POST("/changepass", controllers.ChangePassword)
	protected.GET("/profile", controllers.CurrentUser)
	protected.POST("/add_product", p_controllers.AddProduct)
	protected.POST("/list_product", p_controllers.DownloadPhoto)

	// router.GET("/albums", getAlbums)
	// router.GET("/albums/:id", getAlbumByID)
	// router.POST("/albums", postAlbums)

	router.Run("localhost:8088")
}

package main

import (
	"scalable-final-proj/product_svc/p_controllers"
	"scalable-final-proj/product_svc/p_models"

	"github.com/gin-gonic/gin"
)

func main() {

	p_models.ConnectDataBase()

	router := gin.Default()
	protected := router.Group("/api/user")
	protected.POST("/add_product", p_controllers.AddProduct)
	protected.POST("/list_product", p_controllers.DownloadPhoto)

	// router.GET("/albums", getAlbums)
	// router.GET("/albums/:id", getAlbumByID)
	// router.POST("/albums", postAlbums)

	router.Run("localhost:8001")
}

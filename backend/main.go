package main

import (
	"net/http"

	"controller_svc/controllers"
	"controller_svc/middlewares"
	"controller_svc/models"
	// "scalable-final-proj/backend/product_svc/p_controllers"

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

	protected.POST("/add_product", controllers.AddProduct)
	protected.GET("/list_product", controllers.DownloadPhoto)
	
	protected.POST("/place_order", controllers.PlaceOrder)
	protected.POST("/make_payment", controllers.Pay)
	// router.GET("/albums", getAlbums)
	// router.GET("/albums/:id", getAlbumByID)
	// router.POST("/albums", postAlbums)

	router.Run("localhost:8088")
}

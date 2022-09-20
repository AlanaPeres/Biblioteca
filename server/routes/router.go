package routes

import (
	"github.com/AlanaPeres/Biblioteca/controllers"
	"github.com/gin-gonic/gin"
)

//método
func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		books := main.Group("books")
		{
			books.GET("/", controllers.ShowAllBooks)
		}
	}
	return router
}
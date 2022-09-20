package router

import (
	"github.com/gin-gonic/gin"
)
//método
func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group(api/v1)
	{
		books := main.Group("Books")
		{
			books.GET("/")
		}
	}
	return router
}
package controllers

import "github.com/gin-gonic/gin"

func ShowAllBooks(c *gin.Context) {
	c.JSON(200, gin.H{//gin.H já permite formatar a JSON.
		"value": "ok",
	})
}
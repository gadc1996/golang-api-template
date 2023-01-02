package controllers

import "github.com/gin-gonic/gin"

type LoginController struct {}

func (controller LoginController) Handle(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}
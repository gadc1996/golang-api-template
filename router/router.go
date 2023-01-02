package router

import (
	"controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", new(controllers.LoginController).Handle)
	return router
}

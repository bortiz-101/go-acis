package apihttp

import (
	"github.com/bortiz-101/go-acis/api/internal/http/handlers"
	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	router := gin.Default()

	// define all of our routes here
	router.GET("/health", handlers.Health)
	return router
}

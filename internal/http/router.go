package apihttp

import (
	"github.com/bortiz-101/go-acis/internal/http/handlers"
	"github.com/gin-gonic/gin"
)

// define all of our routes here
func CreateRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/health", handlers.Health)
	return router
}

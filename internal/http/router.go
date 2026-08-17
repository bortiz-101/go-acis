package apihttp

import (
	"github.com/bortiz-101/go-acis/internal/acis"
	"github.com/bortiz-101/go-acis/internal/http/handlers"
	"github.com/gin-gonic/gin"
)

// router now needs ACIS client to be created
func CreateRouter(acisClient *acis.Client) *gin.Engine {
	router := gin.Default()

	// gin router just marries endpoint to handler func
	router.GET("/health", handlers.Health)
	router.POST("api/stn-data", handlers.StnData(acisClient))

	return router
}

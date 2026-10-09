package handlers

import (
	"net/http"

	"github.com/bortiz-101/go-acis/internal/acis"
	"github.com/gin-gonic/gin"
)

func StnData(client *acis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload acis.StnDataRequest

		// decode the JSON sent to our API into StnDataReques struct
		// let gin handle error responses https://gin-gonic.com/en/docs/binding/#bind-vs-shouldbind
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, acis.Error{
				Message: "Invalid JSON payload",
				Code:    "INVALID_JSON",
			})
			return
		}

		// use ACIS client to send req to real service
		result, err := client.StnData(c.Request.Context(), payload)
		if err != nil {
			c.JSON(http.StatusBadGateway, acis.Error{
				Message: err.Error(),
				Code:    "ACIS_ERROR",
			})
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

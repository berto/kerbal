package views

import (
	"fmt"
	"net/http"

	"github.com/berto/kerbal/services"
	"github.com/gin-gonic/gin"
)

func apiHandler(c *gin.Context) {
	awsService := services.New(c)
	if err := awsService.AWSConnect(); err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Failed to connect to aws: %s", err.Error()))
		return
	}
	items, err := awsService.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Failed to list items: %s", err.Error()))
		return
	}
	b, err := awsService.DownloadBytes(items[0].Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Failed to get item: %s", err.Error()))
		return
	}
	c.JSON(http.StatusOK, b)
}

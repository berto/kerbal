package views

import (
	"fmt"
	"net/http"

	"github.com/berto/kerbal/controllers"
	"github.com/gin-gonic/gin"
)

func getItems(c *gin.Context) {
	items, err := controllers.GetItems(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Failed to list items: %s", err.Error()))
		return
	}
	c.JSON(http.StatusOK, items)
}

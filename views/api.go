package views

import (
	"net/http"

	"github.com/berto/kerbal/controllers"
	"github.com/berto/kerbal/responses"
	"github.com/gin-gonic/gin"
)

func getItems(c *gin.Context) {
	items, err := controllers.GetItems(c)
	if err != nil {
		responses.ServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, items)
}

package views

import (
	"net/http"

	"github.com/berto/kerbal/controllers"
	"github.com/gin-gonic/gin"
)

func createKerbal(c *gin.Context) {
	input := controllers.KerbalItems{}
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, "invalid item list: "+err.Error())
		return
	}
	controllers.CreateKerbal(input)
	c.JSON(http.StatusOK, input)
}

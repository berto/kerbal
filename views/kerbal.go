package views

import (
	"net/http"

	"github.com/berto/kerbal/controllers"
	"github.com/berto/kerbal/responses"
	"github.com/gin-gonic/gin"
)

func createKerbal(c *gin.Context) {
	input := controllers.KerbalItems{}
	if err := c.Bind(&input); err != nil {
		responses.NewClientError(c, err)
		return
	}
	if err := input.Validate(); err != nil {
		responses.NewClientError(c, err)
		return
	}
	if err := controllers.CreateKerbal(c, input); err != nil {
		responses.NewServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, input)
}

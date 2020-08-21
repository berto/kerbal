package views

import (
	"github.com/berto/kerbal/controllers"
	"github.com/berto/kerbal/responses"
	"github.com/gin-gonic/gin"
)

func createKerbal(c *gin.Context) {
	input := controllers.KerbalItems{}
	if err := c.Bind(&input); err != nil {
		responses.ClientError(c, err)
		return
	}
	if err := input.Validate(); err != nil {
		responses.ClientError(c, err)
		return
	}
	id, err := controllers.CreateKerbal(c, input)
	if err != nil {
		responses.ServerError(c, err)
		return
	}
	responses.OK(c, id)
}

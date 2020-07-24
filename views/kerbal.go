package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func kerbalHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "Welcome to Corporate Data Dashboard")
}

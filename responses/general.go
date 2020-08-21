package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ClientError sends client error
func ClientError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
	})
	c.Abort()
}

// ServerError sends server error
func ServerError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})
	c.Abort()
}

// OK sends data
func OK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
	c.Abort()
}

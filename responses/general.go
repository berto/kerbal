package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewClientError sends client error
func NewClientError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
	})
	c.Abort()
}

// NewServerError sends server error
func NewServerError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})
	c.Abort()
}

// NewOK sends data
func NewOK(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"data": err.Error(),
	})
	c.Abort()
}

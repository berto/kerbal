package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Home serves index
func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// Download serves index
func Download(c *gin.Context) {
	c.HTML(http.StatusOK, "download.html", nil)
}

// KerbalRoutes manage kerbals
func KerbalRoutes(router *gin.RouterGroup) {
	router.POST("/", createKerbal)
}

// APIRoutes public interface
func APIRoutes(router *gin.RouterGroup) {
	router.GET("/items", getItems)
}
